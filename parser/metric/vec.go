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
	"github.com/prometheus/client_golang/prometheus"
)

// MetricVec represents a wapper for prometheus.MetricVec
type MetricVec interface {
	Labels() []string
	GetVal(labelvalues []string) (float64, error)
}

type VecType interface {
	// Prometheus GaugeVec and CounterVec implement this interface
	GetMetricWithLabelValues(lvs ...string) (prometheus.Metric, error)
}

// vec implements MetricVec interface
type vec struct {
	v      VecType
	labels []string
}

func NewVec[T VecType](v VecType, labels []string) MetricVec {
	return &vec{
		v:      v,
		labels: labels,
	}
}

func (mv *vec) Labels() []string {
	return mv.labels
}

func (mv *vec) GetVal(labelvalues []string) (float64, error) {
	metric, err := mv.v.GetMetricWithLabelValues(labelvalues...)
	if err != nil {
		return 0, err
	}
	return GetMetricVal(metric), nil
}
