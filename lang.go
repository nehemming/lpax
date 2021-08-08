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
	jib "github.com/cloudfoundry-attic/jibber_jabber"
	"golang.org/x/text/language"
)

// DefaultLanguage is the default fallback language.
const DefaultLanguage = "en"

// DetectLocaleLanguage returns the language of the current process or an error if it cannot be found.
func DetectLocaleLanguage() (language.Tag, error) {
	var langCode string
	var err error

	// use jibber jabber to find language from os
	langCode, err = jib.DetectLanguage()

	// Check for failure
	if err != nil {
		return language.Tag{}, err
	}

	// Try and parse the result to a language tag
	return language.Parse(langCode)
}

// MustDetectLocaleLanguage attempts to get the current the users language.
// If this is not discoverable the fallback language string will be tried.
// If this also fails or is not provided the function panics.
func MustDetectLocaleLanguage(fallback string) language.Tag {
	tag, err := DetectLocaleLanguage()
	if err != nil {
		if fallback == "" {
			panic(err)
		}

		tag, err = language.Parse(fallback)
		if err != nil {
			panic(err)
		}
	}

	return tag
}
