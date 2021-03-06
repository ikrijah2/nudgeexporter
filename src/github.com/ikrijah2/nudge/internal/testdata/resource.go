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

package testdata

import (
	otlpresource "github.com/ikrijah2/nudge/internal/data/protogen/resource/v1"

	"github.com/ikrijah2/nudge/consumer/pdata"
)

func initResource1(r pdata.Resource) {
	initResourceAttributes1(r.Attributes())
}

func generateOtlpResource1() otlpresource.Resource {
	return otlpresource.Resource{
		Attributes: generateOtlpResourceAttributes1(),
	}
}

func initResource2(r pdata.Resource) {
	initResourceAttributes2(r.Attributes())
}

func generateOtlpResource2() otlpresource.Resource {
	return otlpresource.Resource{
		Attributes: generateOtlpResourceAttributes2(),
	}
}
