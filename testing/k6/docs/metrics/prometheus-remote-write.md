# Prometheus remote write

Prometheus remote write is a protocol that makes it possible to reliably propagate data in real-time from a sender to a receiver. It has multiple compatible implementations and storage integrations.

For instance, when using the experimental-prometheus-rw output, k6 can send test-result to the remote-write endpoint and store them in Prometheus.

The output, during the k6 run execution, gets all the generated time-series data points for the k6 metrics. It then generates the equivalent Prometheus time series and sends to the Prometheus remote write endpoint.

## Metrics mapping

All k6 metric types are converted into an equivalent Prometheus metric type. The output maps the metrics into time series with Name labels. As much as possible, k6 respects the naming best practices that the Prometheus project defines:

- All time series are prefixed with the k6_ namespace.
- All time series are suffixed with the base unit of the sample value (if k6 knows what the base unit is).
- Trends and rates have the relative suffixes, to make them more discoverable.

## Trend metric conversions

This output provides two distinct mechanism to send k6 Trend metrics to Prometheus:

1. Counter and Gauge metric (default)
2. Prometheus Native histogram

k6 aggregates trend metric of test results while providing high-precision queries.

Note that k6 aggregates trend metric data before sending it to Prometheus in both options. The reasons for aggregating data are:

- Prometheus stores data in a millisecond precision (ms), but k6 metrics collect data points with higher accuracy, nanosecond(ns).
- A load test could generate vast amounts of data points. High-precision raw data could quickly become expensive and complex to scale and is unnecessary when analyzing performance trends.

**1. Counter and gauges**

By default, Prometheus supports Counter and Gauge Metric types. Therefore, this option is the default of this output and converts all k6 Trend metrics to Counter and Gauges Prometheus metrics.

You can configure how to convert all the k6 trend metrics with the K6_PROMETHEUS_RW_TREND_STATS option that  accepts a comma-separated list of stats functions: count, sum, min, max, avg, med, p(x). The default is p(99).

K6_PROMETHEUS_RW_TREND_STATS="p(90),p(95),max" transforms each trend metric into three Prometheus metrics as follow:

- k6_*_p90
- k6_*_p95
- k6_*_max

This option provides a configurable solution to  represent Trend metrics in Prometheus but has the following drawbacks:

- Convert a k6 Trend metric to several Prometheus metrics.
- It is impossible to aggregate some gauge values (especially percentiles).
- It uses a memory-expensive k6 data structure.

**2. Prometheus native histogram**

To address the limitations of the previous options, you can convert k6 trend metrics to high-fidelity histograms enabling Prometheus native histograms.

With this option, each k6 trend metric maps to its corresponding Prometheus histogram metric: k6_*. You can then query them using Prometheus histogram functions, such as histogram_quantile().
