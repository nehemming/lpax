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

func TestContext(t *testing.T) {
	tID := testTextID(23)

	tm := TextMap{
		tID:          "Hello World",
		tID.Plural(): "Hello Worlds",
	}

	ctx := WithContext(context.Background(), tm)

	if ctx == nil {
		t.Error("ctx nil")
		return
	}

	ctxMap := FromContext(ctx)

	if ctxMap == nil {
		t.Error("ctxMap nil")
		return
	}

	msg := ctxMap.Text(tID.Plural())

	if msg != "Hello Worlds" {
		t.Error("ctx message", msg)
	}
}

func TestContextNoSet(t *testing.T) {
	tID := testTextID(23)

	ctxMap := FromContext(context.Background())

	if ctxMap == nil {
		t.Error("ctxMap nil")
		return
	}
	msg := ctxMap.Text(tID.Plural())

	if msg != "" {
		t.Error("ctx message", msg)
	}
}
