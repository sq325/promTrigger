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
	dto "github.com/prometheus/client_model/go"
)

// Metric represents a wapper for prometheus.Metric
type Metric interface {
	GetVal() float64
}

// implement Metric interface
type gauge struct{ d *dto.Metric }
type counter struct{ d *dto.Metric }

// NewMetric returns a Metric interface based on the type of prometheus.Metric
// Only Gauge and Counter are supported
func NewMetric(m prometheus.Metric) Metric {
	d := &dto.Metric{}
	m.Write(d)

	switch m.(type) {
	case prometheus.Gauge:
		return gauge{d}
	case prometheus.Counter:
		return counter{d}
	default:
		return nil
	}
}

func (g gauge) GetVal() float64 {
	return g.d.GetGauge().GetValue()
}

func (c counter) GetVal() float64 {
	return c.d.GetCounter().GetValue()
}

func GetMetricVal(pm prometheus.Metric) float64 {
	m := NewMetric(pm)
	if m == nil {
		return 0
	}
	return m.GetVal()
}
