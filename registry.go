/*
Copyright (c) 2021 The lpax Authors (Neil Hemming)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package lpax

import (
	"fmt"
	"sort"
	"sync"

	"golang.org/x/text/language"
)

// OnRegister is the a callback function signature used to locate a registered
// language pack and return a TextMap containing its contents.
// The callback will only be called for pack IDs and languages that were registered with the registry.
// OnRegister may be called each time Registry New is called.
type OnRegister = func(packID PackID, langTag Tag) TextMap

// TextRegistry maintains a list of registered resource TextMaps by language.
// The registries New function builds language specifc TextProviders from the registered resource strings
// THe registry also creates a shared runtime TextProvider using the process owners detected language.
type TextRegistry interface {
	TextFinder

	// Register maintains a collection of supported language packs by language.
	// The callback OnRegister function will be called if one of the registered language packs
	// needs to be loaded
	// The priority allows multiple overlapping TextMaps to be multiply registered for the same
	// language, giving increasing priority registrations in the order AdditionalPacks, Package and finally Override.
	// langTags is a list of languages that are being registered.  The first language has highest priority.
	Register(packID PackID, callback OnRegister, priority Priority, langTags ...Tag) TextRegistry

	// New returns a new text finder created from the registry.  Each call creates a new finder
	// options can be language Tags and additional TextMaps
	// language Tags must be supplied with the fallback language being first language in the list, if no language is provided the
	// DefaultLanguage is used.
	New(options ...interface{}) TextFinder
}

type (
	packEntry struct {
		priority  Priority
		callback  OnRegister
		supported []Tag
	}

	packEntries []packEntry

	packGroup struct {
		entries  packEntries
		isSorted bool
	}

	packEntryMap map[PackID]*packGroup

	packRegistry struct {
		registered  packEntryMap
		muInit      sync.Mutex // lock called during pack runtime shared initProvider calls
		mu          sync.Mutex // lock on internal structures
		regSequence int
		proSequence int
		textMap     TextMap
	}
)

var (
	once                sync.Once
	defaultTextRegistry TextRegistry
)

// Default returns the default text registry.
func Default() TextRegistry {
	once.Do(func() {
		defaultTextRegistry = NewRegistry()
	})
	return defaultTextRegistry
}

// NewRegistry creates a new text registry.
func NewRegistry() TextRegistry {
	return &packRegistry{
		registered: make(packEntryMap),
	}
}

// Text returns the text identified by the textID or an empty string.
func (r *packRegistry) Text(textID TextID) string {
	return r.initTextProvider().Text(textID)
}

// Find looks up the passed textID key and returns true if found.
func (r *packRegistry) Find(textID TextID) (t string, found bool) {
	return r.initTextProvider().Find(textID)
}

// Register adds a new registration resource for a pack ID and range of languages.
func (r *packRegistry) Register(packID PackID, callback OnRegister, priority Priority, langTags ...Tag) TextRegistry {
	// Check for case where nothing is registered
	n := len(langTags)
	if n == 0 {
		return r
	}

	// Check the pack id is valid
	validateTextID(packID)

	// Create the entry
	cp := make([]Tag, n)
	copy(cp, langTags)

	entry := packEntry{
		priority:  priority,
		callback:  callback,
		supported: cp,
	}

	// Write Lock
	r.mu.Lock()
	defer r.mu.Unlock()

	// Save down the packs
	group, found := r.registered[packID]
	if !found {
		group = newPackGroup(entry)
		r.registered[packID] = group
	} else {
		group.entries = append(group.entries, entry)
		group.isSorted = false
	}

	r.regSequence++

	return r
}

// New creates a new provider.
func (r *packRegistry) New(options ...interface{}) TextFinder {
	return r.newTextMap(options...)
}

func (r *packRegistry) newTextMap(options ...interface{}) TextMap {
	// resolve options
	langTag := make([]Tag, 0, len(options))
	textMaps := make([]TextMap, 0)
	for _, o := range options {
		switch v := o.(type) {
		case Tag:
			langTag = append(langTag, v)
		case string:
			langTag = append(langTag, language.MustParse(v))
		case TextMap:
			textMaps = append(textMaps, v)
		case []TextMap:
			textMaps = append(textMaps, v...)
		default:
			// developer issuer passing wrong type
			panic(fmt.Sprintf("Invalid %[1]T %[1]v", o))
		}
	}

	// default the language if none provided
	if len(langTag) == 0 {
		langTag = append(langTag, language.MustParse(DefaultLanguage))
	}

	// Gather all the text mappings
	textMap := r.getLanguageTextMap(langTag...)

	return textMap.Merge(textMaps...)
}

// newPackGroup creates a new pack group to store language pack registrations.
func newPackGroup(entries ...packEntry) *packGroup {
	return &packGroup{
		entries: append(make(packEntries, 0, len(entries)), entries...),
	}
}

func (r *packRegistry) getLanguageTextMap(langTag ...Tag) TextMap {
	// Lock
	r.mu.Lock()
	defer r.mu.Unlock()

	// List of text maps
	textMaps := make([]TextMap, 0, len(r.registered))

	for packID, group := range r.registered {
		if !group.isSorted {
			sort.Sort(group.entries)
			group.isSorted = true
		}

		entries := group.entries

		// generate a distinct set of supported keys
		distinct := make(map[Tag]bool)
		entryCount := len(entries)
		keys := make([]Tag, 0, entryCount) // guess 1 key per entryCount

		// gather the callbacks into a list at the same time
		callbacks := make([]OnRegister, 0, entryCount)

		// callbackFilter links the callback to the supported languages it was registered to provide
		callbackFilter := make([]map[Tag]bool, entryCount)
		for i, entry := range entries {
			// Add in callback
			callbacks = append(callbacks, entry.callback)

			// Create the filter
			callbackFilter[i] = make(map[Tag]bool)

			// get distinct keys
			for _, tag := range entry.supported {
				callbackFilter[i][tag] = true

				if distinct[tag] {
					continue // ignore as already in use
				}

				distinct[tag] = true
				keys = append(keys, tag)
			}
		}

		// Find best match
		m := language.NewMatcher(keys)
		matchTag, _, _ := m.Match(langTag...)

		for i, callback := range callbacks {
			// skip calling if the callback didn't register the tag
			if !callbackFilter[i][matchTag] {
				continue
			}

			tm := callback(packID, matchTag)

			if tm != nil {
				textMaps = append(textMaps, tm)
				// may be overridden by a later match, continue to search
			}
		}
	}

	return NewTextMap(textMaps...)
}

// initTextProvider is used to initialism a provider.
func (r *packRegistry) initTextProvider() TextMap {
	isCurrent := r.proSequence >= r.regSequence
	tm := r.textMap

	if tm != nil && isCurrent {
		return tm
	}

	// create the text map
	r.muInit.Lock()
	defer r.muInit.Unlock()

	tag := MustDetectLocaleLanguage(DefaultLanguage)

	r.textMap = r.newTextMap(tag)
	r.proSequence = r.regSequence

	return r.textMap
}
