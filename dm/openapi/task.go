// Copyright 2021 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package openapi

import (
	"encoding/json"

	"github.com/pingcap/ticdc/dm/pkg/terror"
)

var defaultMetaSchema = "dm_meta"

// Adjust adjusts task and set default value.
func (t *Task) Adjust() error {
	if t.MetaSchema == nil {
		t.MetaSchema = &defaultMetaSchema
	}
	// check some not implemented features
	if t.OnDuplicate != TaskOnDuplicateError {
		return terror.ErrOpenAPICommonError.Generate("`on_duplicate` only supports `error` for now.")
	}
	return nil
}

// FromJSON unmarshal json to task.
func (t *Task) FromJSON(data []byte) error {
	return json.Unmarshal(data, t)
}

// ToJSON marshal json to task.
func (t *Task) ToJSON() ([]byte, error) {
	return json.Marshal(t)
}
