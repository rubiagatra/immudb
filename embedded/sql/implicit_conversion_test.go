/*
Copyright 2022 Codenotary Inc. All rights reserved.

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

package sql

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApplyImplicitConversion(t *testing.T) {
	for _, d := range []struct {
		val          interface{}
		requiredType SQLValueType
		expected     interface{}
	}{
		{1, IntegerType, int64(1)},
		{1, Float64Type, float64(1)},
		{1.0, Float64Type, float64(1)},
		{"1", IntegerType, int64(1)},
		{"4.2", Float64Type, float64(4.2)},
	} {
		t.Run(fmt.Sprintf("%+v", d), func(t *testing.T) {
			convVal, err := mayApplyImplicitConversion(d.val, d.requiredType)
			require.NoError(t, err)
			require.EqualValues(t, d.expected, convVal)
		})
	}
}
