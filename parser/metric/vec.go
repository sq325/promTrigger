// Copyright 2024 Sun Quan
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metric

import (
	"errors"

	"github.com/prometheus/client_golang/prometheus"
)

type MetricVec interface {
	Labels() []string
	GetVal(labelvalues []string) (float64, error)
}

type VecType interface {
	prometheus.GaugeVec | prometheus.CounterVec
}

// vec implements MetricVec interface
type vec[T VecType] struct {
	v      *T
	labels []string
}

func NewVec[T VecType](v *T, labels []string) MetricVec {
	return &vec[T]{
		v:      v,
		labels: labels,
	}
}

func (mv *vec[T]) Labels() []string {
	return mv.labels
}

func (mv *vec[T]) GetVal(labelvalues []string) (float64, error) {
	switch v := any(mv.v).(type) {
	case *prometheus.GaugeVec:
		g, err := v.GetMetricWithLabelValues(labelvalues...)
		if err != nil {
			return 0, err
		}
		m := NewMetric(g)
		return m.GetVal(), nil
	case *prometheus.CounterVec:
		c, err := v.GetMetricWithLabelValues(labelvalues...)
		if err != nil {
			return 0, err
		}
		m := NewMetric(c)
		return m.GetVal(), nil
	default:
		return 0, errors.New("unsupported type")
	}
}
