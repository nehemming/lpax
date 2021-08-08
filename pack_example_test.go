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

package lpax_test

import (
	"fmt"

	"github.com/nehemming/lpax"
	"golang.org/x/text/language"
)

type (
	// PackID is the local package id.
	// Each package should implement its own PackID type.
	// They are typically derived from the `int` type.
	PackID int

	// TextID is the id type for all messages defined by this package.
	// the TextID type must implement the lpax.TextID interface, but is typically
	// based on a `int` type.  As the package defines its own TextID it is
	// unique from other package types, and thus so are the ids.
	TextID int
)

// Single returns the id of for a single version of a message.
// Typical implementations use the lpax helper functions to implement the id system.
// Positive id's are used for single messages while -ve id's areused for plural versions.
func (id TextID) Single() lpax.TextID {
	return TextID(lpax.IntTypeSingle(int(id)))
}

// Plural returns the plural version of the id.
func (id TextID) Plural() lpax.TextID {
	return TextID(lpax.IntTypePlural(int(id)))
}

// String implements stringer function.
// ReflectCoderString reflects the package name used
// by the parent package along with the +ve integer id of the key.
func (id TextID) String() string {
	return lpax.ReflectCoderString(id.Single(), 1)
}

const (
	// ExamplePackID is the unique typed id for this pack.
	ExamplePackID = PackID(1)

	// None is a default empty message (id 0).
	None = TextID(iota)

	// Hello is an example message.
	Hello
)

var pack = lpax.TextMap{
	Hello:  "Hello World",
	-Hello: "Hello Worlds",
}

// register is the text registry callback function used to return the text mappings.
func register(packID lpax.PackID, langTag lpax.Tag) lpax.TextMap {
	return pack
}

func init() {
	// register the text mappings with the share test registry.
	lpax.Default().Register(ExamplePackID, register, lpax.DefaultPriority, language.English)
}

func Example() {
	fmt.Println(lpax.Sprintf(Hello))
	fmt.Println(lpax.Sprintf(Hello.Plural()))

	// Output:
	// Hello World
	// Hello Worlds
}
