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

// Priority type to specify pack registration priority
// Priority is used by Text Registration types to prioritize their
// stored items.
// Items registered by an implementing package should register with Package priority
// If registering additional languages can use AdditionalPacks priority
// There packs will not take priority over the package implementors text map for a
// language defined in both.  However AdditionalPacks priority allows extra languages
// to be added.
// Override priority can (and should only) be used by main application registrations
// where there is a need to replace a package registrations.
type Priority int

const (

	// AdditionalPacks are added after the original packs.
	AdditionalPacks = Priority(iota)

	// Package standard priority when a language included
	// its own messages, these are treated as the default.
	Package

	// Override use when registration wants preference over
	// the default implementation.
	// This should only be done by applications (package main's).
	Override

	// DefaultPriority default type priority.
	DefaultPriority = Package
)
