// Copyright 2024 Humanitec
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

package util

import (
	"reflect"
	"testing"
)

func TestPrepareEnvVariables(t *testing.T) {
	tests := []struct {
		name string
		arr  []string
		want []string
	}{
		{"PrepareEnvVariables with one letter",
			[]string{
				"echo hello $A ${B} $C world",
				"echo hello $A-${B}-$C world",
				"echo hello $A${B}$C world",
			},
			[]string{
				"echo hello $$A $${B} $$C world",
				"echo hello $$A-$${B}-$$C world",
				"echo hello $$A$${B}$$C world",
			},
		},
		{"PrepareEnvVariables with several letters",
			[]string{
				"echo hello $AAA ${BBB} $CCC world",
				"echo hello $AAA-${BBB}-$CCC world",
				"echo hello $AAA${BBB}$CCC world",
			},
			[]string{
				"echo hello $$AAA $${BBB} $$CCC world",
				"echo hello $$AAA-$${BBB}-$$CCC world",
				"echo hello $$AAA$${BBB}$$CCC world",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrepareEnvVariables(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrepareEnvVariables()\nget  = %v\nwant = %v\n", got, tt.want)
			}
		})
	}
}
