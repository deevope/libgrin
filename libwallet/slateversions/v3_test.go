// Copyright 2020 BlockCypher
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http//www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package slateversions

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalSlateV3(t *testing.T) {
	slateV3JSON, _ := ioutil.ReadFile("test_data/v3.slate")
	var slateV3 SlateV3
	err := json.Unmarshal(slateV3JSON, &slateV3)
	assert.Nil(t, err)
}

func TestMarshalSlateV3(t *testing.T) {
	slateV3JSON, _ := ioutil.ReadFile("test_data/v3_raw.slate")
	var slateV3 SlateV3
	err := json.Unmarshal(slateV3JSON, &slateV3)
	assert.Nil(t, err)

	serializedSlateV3, err := json.Marshal(slateV3)
	assert.Equal(t, slateV3JSON, serializedSlateV3)
}
