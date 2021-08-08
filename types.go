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

	"golang.org/x/text/language"
)

type (
	// TextID is a unique identifier for a text message.
	// Source/package family should use a unique type to identify its messages
	// valid types must be of an integer (intX or UIntX) type, a string or a struct
	// The provider will panic if any other type is used
	// If the identifier supports exterr.ErrorCoder or fmt.Stringer these interfaces will be used
	// when rasing errors using Errorf or Sprintf.
	TextID interface {
		fmt.Stringer
		// Single is the id for the singular version of the text.
		Single() TextID
		// Plural is the id for the plural version of the text.
		Plural() TextID
	}

	// PackID uses Go's strong typing system to create a unique identifier for a collection of string resources
	// The pack ID is used to link different language version of the same pack together.
	PackID = interface{}

	// Tag alias language.Tag.
	Tag = language.Tag

	// TextFinder looks up a text ID and returns the string associated with it or an empty string, found.
	TextFinder interface {
		// Text looks up a text ID and returns the string associated with it or an empty string.
		Text(textID TextID) string

		// Find looks up the passed textID key and returns true if found
		Find(textID TextID) (t string, found bool)
	}
)

// ByCount returns the plural version of a count if count 1= 1.
func ByCount(id TextID, count int) TextID {
	if count != 1 {
		return id.Plural()
	}
	return id.Single()
}
