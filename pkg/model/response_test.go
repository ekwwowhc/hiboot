// Copyright 2018 John Deng (hi.devops.io@gmail.com).
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResponse(t *testing.T) {
	var response Response

	response = new(BaseResponse)

	response.SetCode(200)
	assert.Equal(t, 200, response.GetCode())

	response.SetMessage("Hello world")
	assert.Equal(t, "Hello world", response.GetMessage())

	response.SetData("example")
	assert.Equal(t, "example", response.GetData())
}
