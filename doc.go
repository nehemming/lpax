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

// Package lpax is a strong keyed string repository implementation supporting
// internationalisation and separation of text definition from usage.
//
// Applications and library packages typically embed string and error messages
// within the code.  Placing strings inline within the code can make it harder to
// internationalise the code or to extract a list of all error messages generated
// by an application.  `lpax` provides method of registering strings that can be
// accessed via decentralised keys.  The key system uses Go's strong typing to
// allow each package to create its own key system without risking overlapping
// with other packages keys.  multiple languages can be supported for key, with
// fallback logic to us the default `english` language.
//
// The example below shows how to implement a typical language package.
//
package lpax
