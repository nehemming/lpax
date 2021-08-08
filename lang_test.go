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
	"os"
	"testing"
)

func TestGetLocaleLanguage(t *testing.T) {
	// os test
	locale := os.Getenv("LANG")
	defer func() {
		if locale != "" {
			os.Setenv("LANG", locale)
		}
	}()

	os.Setenv("LANG", "en_US")

	l, e := DetectLocaleLanguage()

	if e != nil {
		t.Error(e)
	}

	if l.String() != "en" {
		t.Error("wrong lang (" + l.String() + ")")
	}
}

func TestMustGetLocaleLanguage(t *testing.T) {
	// os test
	locale := os.Getenv("LANG")

	defer func() {
		if r := recover(); r != nil {
			t.Error("panic")
		}
	}()

	defer func() {
		if locale != "" {
			os.Setenv("LANG", locale)
		}
	}()
	os.Setenv("LANG", "en_US")

	l := MustDetectLocaleLanguage("es")

	if l.String() != "en" {
		t.Error("wrong lang (" + l.String() + ")")
	}
}

func TestMustGetLocaleLanguageFallsback(t *testing.T) {
	// os test
	locale := os.Getenv("LANG")

	defer func() {
		if r := recover(); r != nil {
			t.Error("panic")
		}
	}()

	defer func() {
		if locale != "" {
			os.Setenv("LANG", locale)
		}
	}()
	os.Setenv("LANG", "")

	l := MustDetectLocaleLanguage("es")

	if l.String() != "es" {
		t.Error("wrong lang (" + l.String() + ")")
	}
}
