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

func (set ScalarSet) ToSeriesList(timerange api.Timerange) (api.SeriesList, *ConversionFailure) {
	list := api.SeriesList{
		Series: make([]api.Timeseries, len(set)),
	}
	for i := range list.Series {
		list.Series[i] = api.Timeseries{
			TagSet: set[i].TagSet,
			Values: make([]float64, timerange.Slots()),
		}
		for j := range list.Series[i].Values {
			list.Series[i].Values[j] = set[i].Value
		}
	}
	return list, nil
}
func (set ScalarSet) ToString() (string, *ConversionFailure) {
	return "", &ConversionFailure{
		From: "scalar set",
		To:   "string",
	}
}
func (set ScalarSet) ToScalar() (float64, *ConversionFailure) {
	if len(set) == 1 && set[0].TagSet.Equals(api.TagSet{}) {
		return set[0].Value, nil
	}
	return 0, &ConversionFailure{
		From: "scalar set",
		To:   "scalar",
	}
}
func (set ScalarSet) ToScalarSet() (ScalarSet, *ConversionFailure) {
	return set, nil
}
func (set ScalarSet) ToDuration() (time.Duration, *ConversionFailure) {
	return 0, &ConversionFailure{
		From: "scalar set",
		To:   "duration",
	}
}
