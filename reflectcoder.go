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
	"path"
	"reflect"
	"strings"
)

// ReflectCoderString supports the Coder interface by generating a description for
// an error code from its package path.
// v must be a int value kind, otherwise v is returned as a "%v" formatted string
// level is a single variadic value optionally specifying how many levels to traverse
// to get the path name.  eg. package/path/here/place would return
// 0 - place, 1 - here etc.
func ReflectCoderString(v interface{}, level ...int) string {
	rt := reflect.TypeOf(v)

	l := 0
	if len(level) == 1 {
		l = level[0]
	}

	if rt.Kind() == reflect.Int {
		pkgPath := rt.PkgPath()

		for i := 0; i < l; i++ {
			pkgPath = strings.TrimRight(path.Dir(pkgPath), "/\\")
		}
		name := path.Base(strings.TrimRight(pkgPath, "/\\"))

		return fmt.Sprintf("%s-%05d", name, v)
	}

	return fmt.Sprintf("%v", v)
}
