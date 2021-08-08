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
	"reflect"
)

func validateTextID(textID interface{}) {
	// run time check that the textID is actually one of the permitted types
	// if a invalid type has been used the program will panic.
	r := reflect.ValueOf(textID)

	switch r.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return
	case reflect.String:
		return
	case reflect.Struct:
		return
	}

	panic(fmt.Sprintf("invalid type:%T", textID))
}

// TextMap maps TextID keys to strings.
type TextMap map[TextID]string

// NewTextMap creates a new text map by merging zero or more exiting maps.
func NewTextMap(texts ...TextMap) TextMap {
	tm := make(TextMap)
	return tm.Merge(texts...)
}

// Merge combines tha passed TextMap's with the receiver TextMap.
// Items are merged left to right.  Duplicate keys are overridden subsequent
// maps.
func (tm TextMap) Merge(texts ...TextMap) TextMap {
	for _, a := range texts {
		for k, v := range a {
			// check the key is a valid type, will panic if not
			validateTextID(k)

			tm[k] = v
		}
	}
	return tm
}

// Text returns the text identified by the textID or an empty string.
func (tm TextMap) Text(textID TextID) string {
	return tm[textID]
}

// Find looks up the passed textID key and returns true if found.
func (tm TextMap) Find(textID TextID) (t string, found bool) {
	f, ok := tm[textID]
	return f, ok
}
