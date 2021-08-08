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
	"fmt"
	"strings"
)

// Sprintf is identical to fmt.Sprintf except the format string is taken from the default text finder.
// If the string is not found the string version of the id is printed along with a
// space separated %v version of each arg.
func Sprintf(id TextID, args ...interface{}) string {
	tm := Default()

	f, ok := tm.Find(id)

	if !ok {
		f = strings.TrimRight("%s:"+strings.Repeat("%v ", len(args)), " ")

		args = append([]interface{}{id}, args...)
	}

	return fmt.Sprintf(f, args...)
}

// CtxSprintf is identical to fmt.Sprintf except the format string is taken from the text finder
// linked to the passed context. If the context has no finder the default finder is used.
// If the string is not found the string version of the id is printed along with a
// space separated %v version of each arg.
func CtxSprintf(ctx context.Context, id TextID, args ...interface{}) string {
	tm := FromContext(ctx)

	f, ok := tm.Find(id)

	if !ok {
		f = strings.TrimRight("%s:"+strings.Repeat("%v ", len(args)), " ")

		args = append([]interface{}{id}, args...)
	}

	return fmt.Sprintf(f, args...)
}

// Errorf is identical to fmt.Errorf except the format string is taken from the default text finder.
// If the string is not found the string version of the id is printed along with a
// space separated %v version of each arg.
func Errorf(id TextID, args ...interface{}) error {
	tm := Default()

	f, ok := tm.Find(id)

	if !ok {
		f = strings.TrimRight("%s:"+strings.Repeat("%v ", len(args)), " ")

		args = append([]interface{}{id}, args...)
	}

	return fmt.Errorf(f, args...)
}

// CtxErrorf is identical to fmt.Errorf except the format string is taken from the text finder
// linked to the passed context. If the context has no finder the default finder is used.
// If the string is not found the string version of the id is printed along with a
// space separated %v version of each arg.
func CtxErrorf(ctx context.Context, id TextID, args ...interface{}) error {
	tm := FromContext(ctx)

	f, ok := tm.Find(id)

	if !ok {
		f = strings.TrimRight("%s:"+strings.Repeat("%v ", len(args)), " ")

		args = append([]interface{}{id}, args...)
	}

	return fmt.Errorf(f, args...)
}
