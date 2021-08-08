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
	"testing"

	"golang.org/x/text/language"
)

type (
	testTextID int
)

func (id testTextID) Single() TextID {
	return testTextID(IntTypeSingle(int(id)))
}

func (id testTextID) Plural() TextID {
	return testTextID(IntTypePlural(int(id)))
}

func (id testTextID) String() string {
	return fmt.Sprintf("%d", int(id))
}

type TestPackID int

const (
	// ExamplePackID is the unique typed id for this pack.
	ExamplePackID = TestPackID(1)

	// None is a default empty message.
	None = testTextID(iota)

	// Hello is an example message.
	Hello

	// Args has args for testing.
	Args
)

var pack = TextMap{
	Hello:  "Hello World",
	-Hello: "Hello Worlds",
	Args:   "Single %v",
	-Args:  "Plurals %v",
}

var additionalPack = TextMap{
	Hello:  "Hello Moon",
	-Hello: "Hello Moons",
	Args:   "Not Single %v",
	-Args:  "Not Plurals %v",
}

var spanishPack = TextMap{
	Hello:  "Hola Mundo",
	-Hello: "Hola Mundos",
	Args:   "único %v",
	-Args:  "plurales %v",
}

func init() {
	// register the text mappings with the share test registry.
	Default().Register(ExamplePackID, func(packID PackID, langTag Tag) TextMap {
		return pack
	}, DefaultPriority, language.English)

	Default().Register(ExamplePackID, func(packID PackID, langTag Tag) TextMap {
		return additionalPack
	}, AdditionalPacks, language.English)

	Default().Register(ExamplePackID, func(packID PackID, langTag Tag) TextMap {
		return spanishPack
	}, DefaultPriority, language.Spanish)
}

func TestByCountZero(t *testing.T) {
	id := ByCount(testTextID(1), 0)
	if id.(testTextID) != -1 {
		t.Error("Plural", id)
	}
}

func TestByCountOne(t *testing.T) {
	id := ByCount(testTextID(1), 1)
	if id.(testTextID) != 1 {
		t.Error("Single", id)
	}
}

func TestByCountTwo(t *testing.T) {
	id := ByCount(testTextID(1), 2)
	if id.(testTextID) != -1 {
		t.Error("Plural", id)
	}
}
