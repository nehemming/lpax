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

func TestSprintf(t *testing.T) {
	// Test default
	tID := testTextID(24)

	s := Sprintf(tID, 1, "a")
	if s != "24:1 a" {
		t.Error("Mismatch ", s)
	}
}

func TestSprintfArgs(t *testing.T) {
	// Test default

	s := Sprintf(Args, 1)
	if s != "Single 1" {
		t.Error("Mismatch ", s)
	}
}

func TestErrorf(t *testing.T) {
	// Test default
	tID := testTextID(24)

	s := Errorf(tID, 1, "a")
	if s.Error() != "24:1 a" {
		t.Error("Mismatch ", s)
	}
}

func TestErrorfArgs(t *testing.T) {
	// Test default

	s := Errorf(Args, 1)
	if s.Error() != "Single 1" {
		t.Error("Mismatch ", s)
	}
}

func TestCtxSprintf(t *testing.T) {
	// Test default
	tID := testTextID(24)

	tm := TextMap{}

	ctx := WithContext(context.Background(), tm)

	s := CtxSprintf(ctx, tID, 1, "a")
	if s != "24:1 a" {
		t.Error("Mismatch ", s)
	}
}

func TestCtxSprintfArgs(t *testing.T) {
	// Test default

	tm := TextMap{
		Hello:  "CTX Hello World",
		-Hello: "Hello Worlds",
		Args:   "CTX Single %v",
		-Args:  "Plurals %v",
	}

	ctx := WithContext(context.Background(), tm)

	s := CtxSprintf(ctx, Args, 1)
	if s != "CTX Single 1" {
		t.Error("Mismatch ", s)
	}
}

func TestCtxErrorf(t *testing.T) {
	// Test default
	tID := testTextID(24)

	tm := TextMap{
		Hello:  "CTX Hello World",
		-Hello: "Hello Worlds",
		Args:   "CTX Single %v",
		-Args:  "Plurals %v",
	}

	ctx := WithContext(context.Background(), tm)

	s := CtxErrorf(ctx, tID, 1, "a")
	if s.Error() != "24:1 a" {
		t.Error("Mismatch ", s)
	}
}

func TestCtxErrorfArgs(t *testing.T) {
	// Test default

	tm := TextMap{
		Args: "CTX Single %v",
	}

	ctx := WithContext(context.Background(), tm)

	s := CtxErrorf(ctx, Args, 1)
	if s.Error() != "CTX Single 1" {
		t.Error("Mismatch ", s)
	}
}

func TestNewInSpanish(t *testing.T) {
	tm := Default().New("es")
	ctx := WithContext(context.Background(), tm)

	s := CtxSprintf(ctx, Hello)
	if s != "Hola Mundo" { // nolint
		t.Error("Mismatch ", s)
	}
}
