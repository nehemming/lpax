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
	"context"
	"testing"
)

func TestNewRegistry(t *testing.T) {
	r := NewRegistry()

	_, found := r.Find(testTextID(10))

	if found {
		t.Error("Find found something")
	}

	s := Sprintf(testTextID(10))
	if s != "10:" {
		t.Error("s", s)
	}

	// Default new
	ctx := WithContext(context.Background(), r.New())
	s = CtxSprintf(ctx, testTextID(10))
	if s != "10:" {
		t.Error("S is", s)
	}
}

func TestNewTextMapAddMap(t *testing.T) {
	tm := TextMap{
		Hello:  "CTX Hello World",
		-Hello: "Hello Worlds",
		Args:   "CTX Single %v",
		-Args:  "Plurals %v",
	}

	r := NewRegistry().(*packRegistry)

	m := r.newTextMap(tm)

	if len(m) != 4 {
		t.Error("len m", len(m))
	}
}

func TestNewTextMapAddMaps(t *testing.T) {
	tm := []TextMap{{
		Hello:  "CTX Hello World",
		-Hello: "Hello Worlds",
	}, {
		Args:  "CTX Single %v",
		-Args: "Plurals %v",
	}}

	r := NewRegistry().(*packRegistry)

	m := r.newTextMap(tm)

	if len(m) != 4 {
		t.Error("len m", len(m))
	}
}

func TestNewTextMapStringTag(t *testing.T) {
	tm := []TextMap{{
		Hello:  "CTX Hello World",
		-Hello: "Hello Worlds",
	}, {
		Args:  "CTX Single %v",
		-Args: "Plurals %v",
	}}

	r := NewRegistry().(*packRegistry)

	m := r.newTextMap("en", tm)

	if len(m) != 4 {
		t.Error("len m", len(m))
	}
}

func TestNewTextMapBadArgPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("no panic")
		}
	}()

	r := NewRegistry().(*packRegistry)

	r.newTextMap(10.7)
}
