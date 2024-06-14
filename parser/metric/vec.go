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
