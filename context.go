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

import "context"

// contextKey is a private types string used as the TextProvider context value key.
type contextKey string

// lpaxContextKey context key to provider
// set context using ctx = context.WithValue(ctx,lpax.lpaxContextKey, lpax.NewProvider( etc.
const lpaxContextKey = contextKey("lpax")

// WithContext creates a new context based off the passed ctx with the textFinder bound into the new context.
func WithContext(ctx context.Context, tf TextFinder) context.Context {
	return context.WithValue(ctx, lpaxContextKey, tf)
}

// FromContext returns the TextFinder bound into the context or if none if sound returns the Default TextFinder.
func FromContext(ctx context.Context) TextFinder {
	if tf, ok := ctx.Value(lpaxContextKey).(TextFinder); ok {
		return tf
	}
	// use default
	return Default()
}
