// Copyright 2015 - 2016 Square Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package function

import (
	"time"

	"github.com/square/metrics/api"
)

type TaggedScalar struct {
	TagSet api.TagSet
	Value  float64
}

type ScalarSet []TaggedScalar

// ToSeriesList is a conversion function.
func (set ScalarSet) ToSeriesList(timerange api.Timerange, description string) (api.SeriesList, error) {
	return api.SeriesList{}, ConversionError{"scalar set", "SeriesList", description}
}

// ToString is a conversion function.
func (set ScalarSet) ToString(description string) (string, error) {
	return "", ConversionError{"scalar set", "string", description}
}

// ToScalar is a conversion function.
func (set ScalarSet) ToScalar(description string) (float64, error) {
	return 0, ConversionError{"scalar set", "scalar", description}
}

// ToScalarSet is a conversion function.
func (set ScalarSet) ToScalarSet(description string) (ScalarSet, error) {
	return set, nil
}

// ToDuration is a conversion function.
func (set ScalarSet) ToDuration(description string) (time.Duration, error) {
	return 0, ConversionError{"scalar set", "duration", description}
}
