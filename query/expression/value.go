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

package expression

import (
	"fmt"
	"time"

	"github.com/square/metrics/function"
)

type Duration struct {
	Literal  string
	Duration time.Duration
}

func (expr Duration) Evaluate(context function.EvaluationContext) (function.Value, error) {
	return function.NewDurationValue(expr.Literal, expr.Duration), nil
}

func (expr Duration) Name() string {
	return expr.Literal
}
func (expr Duration) QueryString() string {
	return expr.Literal
}

type Scalar struct {
	Value float64
}

func (expr Scalar) Evaluate(context function.EvaluationContext) (function.Value, error) {
	return function.ScalarValue(expr.Value), nil
}

func (expr Scalar) Name() string {
	return fmt.Sprintf("%+v", expr.Value)
}

func (expr Scalar) QueryString() string {
	return fmt.Sprintf("%+v", expr.Value)
}

type String struct {
	Value string
}

func (expr String) Evaluate(context function.EvaluationContext) (function.Value, error) {
	return function.StringValue(expr.Value), nil
}

func (expr String) Name() string {
	return fmt.Sprintf("%q", expr.Value)
}

func (expr String) QueryString() string {
	return fmt.Sprintf("%q", expr.Value)
}
