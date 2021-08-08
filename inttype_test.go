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

import "testing"

func TestIntTypeSingle(t *testing.T) {
	id := IntTypeSingle(1)
	if id != 1 {
		t.Error("Single", id)
	}
}

func TestIntTypeSingleSingle(t *testing.T) {
	id := IntTypeSingle(IntTypeSingle(1))
	if id != 1 {
		t.Error("SingleSingle", id)
	}
}

func TestIntTypePluralSingle(t *testing.T) {
	id := IntTypeSingle(IntTypePlural(1))
	if id != 1 {
		t.Error("PluralSingle", id)
	}
}

func TestIntTypeSinglePlural(t *testing.T) {
	id := IntTypePlural(IntTypeSingle(1))
	if id != -1 {
		t.Error("SinglePlural", id)
	}
}

func TestIntTypePlural(t *testing.T) {
	id := IntTypePlural(1)
	if id != -1 {
		t.Error("Plural", id)
	}
}

func TestIntTypePluralPlural(t *testing.T) {
	id := IntTypePlural(IntTypePlural(1))
	if id != -1 {
		t.Error("Plural", id)
	}
}
