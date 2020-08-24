文档 https://frezc.github.io/2019/08/03/prometheus-metrics/
https://yuerblog.cc/2019/01/03/prometheus-client-usage-and-principle/

- counter
- gauge
- summary
- Histogram


gauge
gauge可增可减，可以任意设置，就代表了某个指标当前的值而已。

比如可以设置当前的CPU温度，内存使用量等等，它们都是上下浮动的，不是只增不减的。

定义和counter基本一样：

Gauge是一个用来记录实时值的指标，常用于表示CPU使用率、内存使用率、线程数等指标。

比如prometheus暴露的go协程数指标

# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 40
用例
gauge类型指标最常见的就是用来标识服务是否存活的up指标了，这个指标在大多的exporter上都会有，属于一个可以建通用警报规则的指标。

大多数gauge指标用法差不多，我就拿go_goroutines来举例。查询go_goroutines我们可以看到一些自动生成的labels，主要关注instance，这个代表了我们的prometheus实例。假设我们有prometheus实例并且想查询所有实例的平均协程数。



请求次数
rate(grpc_request_total{type="sv"}[5m])
表示过去五分钟每秒收到的请求数
请求花费时间

histogram_quantile(0.99, sum by (le) (rate(request_cost_second_bucket[10m])))

histogram_quantile(0.9, sum by (le) (rate(request_cost_second_bucket[10m])))


histogram_quantile(0.5, sum by (le) (rate(request_cost_second_bucket[10m])))

