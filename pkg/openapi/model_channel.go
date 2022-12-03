// Copyright 2022 Cisco Systems, Inc. and its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Flame REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Channel - Defines how different roles are connected.
type Channel struct {
	Name string `json:"name"`

	Description string `json:"description,omitempty"`

	Pair []string `json:"pair"`

	GroupBy ChannelGroupBy `json:"groupBy,omitempty"`

	FuncTags map[string][]string `json:"funcTags,omitempty"`

	IsUnidirectional bool `json:"isUnidirectional,omitempty"`

	Backend CommBackend `json:"backend,omitempty"`
}

// AssertChannelRequired checks if the required fields are not zero-ed
func AssertChannelRequired(obj Channel) error {
	elements := map[string]interface{}{
		"name": obj.Name,
		"pair": obj.Pair,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertChannelGroupByRequired(obj.GroupBy); err != nil {
		return err
	}
	return nil
}

// AssertRecurseChannelRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Channel (e.g. [][]Channel), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseChannelRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aChannel, ok := obj.(Channel)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertChannelRequired(aChannel)
	})
}
