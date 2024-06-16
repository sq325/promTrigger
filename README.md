# Go Prometheus Metrics Library

这是一个用于处理Prometheus指标的Go库。它提供了一种简单的方式来创建和操作Gauge和Counter类型的指标。

## 主要特性

- 支持Gauge和Counter类型的指标
- 提供了一个简单的接口来获取指标值
- 支持带有标签的指标

## 使用方法

首先，你需要导入`metric`包：

```go
import "metric"
然后，你可以使用NewMetric函数来创建一个新的指标：

m := metric.NewMetric(myPrometheusMetric)
你可以使用GetVal方法来获取指标的值：

val := m.GetVal()
如果你的指标有标签，你可以使用NewVec函数来创建一个新的带有标签的指标，并使用GetVal方法来获取指标的值：

mv := metric.NewVec(myPrometheusMetricVec, myLabels)
val, err := mv.GetVal(myLabelValues)
许可证
本库在Apache License 2.0下发布。详情请参阅LICENSE文件。
