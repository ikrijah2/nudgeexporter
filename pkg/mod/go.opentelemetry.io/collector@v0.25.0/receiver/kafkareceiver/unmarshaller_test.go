// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kafkareceiver

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultTracesUnMarshaller(t *testing.T) {
	expectedEncodings := []string{
		"otlp_proto",
		"jaeger_proto",
		"jaeger_json",
		"zipkin_proto",
		"zipkin_json",
		"zipkin_thrift",
	}
	marshallers := defaultTracesUnmarshallers()
	assert.Equal(t, len(expectedEncodings), len(marshallers))
	for _, e := range expectedEncodings {
		t.Run(e, func(t *testing.T) {
			m, ok := marshallers[e]
			require.True(t, ok)
			assert.NotNil(t, m)
		})
	}
}

func TestDefaultLogsUnMarshaller(t *testing.T) {
	expectedEncodings := []string{
		"otlp_proto",
	}
	marshallers := defaultLogsUnmarshallers()
	assert.Equal(t, len(expectedEncodings), len(marshallers))
	for _, e := range expectedEncodings {
		t.Run(e, func(t *testing.T) {
			m, ok := marshallers[e]
			require.True(t, ok)
			assert.NotNil(t, m)
		})
	}
}
