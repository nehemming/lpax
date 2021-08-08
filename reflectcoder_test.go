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

type carID int

func TestReflectCoderString(t *testing.T) {
	s := ReflectCoderString(carID(1))

	if s != "lpax-00001" {
		t.Error(s)
	}

	s = ReflectCoderString("car")

	if s != "car" {
		t.Error(s)
	}

	s = ReflectCoderString(carID(1), 0)

	if s != "lpax-00001" {
		t.Error(s)
	}
}
