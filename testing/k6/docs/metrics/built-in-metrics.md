# Metrics

Metrics measure how a system performs under test conditions. By default, k6 automatically collects built-in metrics. Besides built-ins, you can also make custom metrics.

Metrics fall into four broad types:

- **Counters** sum values.
- **Gauges** track the smallest, largest, and latest values.
- **Rates** track how frequently a non-zero value occurs.
- **Trends** calculates statistics for multiple values (like mean, mode or percentile).

To make a test fail a certain criteria, you can write a Threshold based on the metric criteria (the specifics of the expression depend on the metric type). To filter metrics, you can use Tags and groups. You can also export metrics in various summary and granular formats, as documented in Results output.

## What metrics to look at?

Each metric provides a different perspective on performance. So the best metric for your analysis depends on your goals. 

Howerver, if you're unsure about the metrics to focus on, you can start with the metrics that measures the requests, errors, and duration.

- http_reqs, to measure requests
- http_request_failed, to measure error rate
- http_req_duration, to measure duration

## Metric name restrictions

Metric must not start with a number and a metric name can be 1 to 128 symbols of:

1. any Unicode Letters
2. any Unicode Number
3. _ (an underscore)

## Built-in metrics

Every k6 test emits built-in and Custom metrics. Each supported protocol also has its specific metrics.

**Standard built-in metrics**

**HTTP-specific built-in metrics**

## Create custom metrics

Besides the built-in metrics, you can create custom metrics. For example, you can compute a metric for your business logic, or use the Response.timings object to create a metric for a specific set of endpoints.

Each metric type has a constructor to create a custom metric. The constructor creates a metric object of the declared type. Each type has an add method to take metric measurements.

### Counter

Counter is an object for representing a custom cumulative counter metric. It is one of the four custom metric types. It is one of the four custom metric.

|Parameter|Type|Description|
|-|-|-|
|name|string|The name of the custom metric|

|Method|Description|
|-|-|
|Counter.add(value, [tags])|Add a value to the counter metric|

**Counter usage in Threshold**

When Counter is used in a threshold expression, the variable must be called count or rate (lower rate). For example:

- count >= 200 // value of the counter must be larger or equal to 200.
- count < 10 // less than 10.

```js
import {Counter} from 'k6/metrics';

const myCounter = new Counter('my_counter');

export default function() {
    myCounter.add(1);
    myCounter.add(2, {tag1: 'myValue', tag2: 'myValue2'});
}
```

```js
import http from 'k6/http';
import {Counter} from 'k6/metrics'; 

const CounterErrors = new Counter('Errors') ;

export const options = {thresholds: {Errors: ['count<100']}};

export default function() {
    const resp = http.get('https://test-api.k6.io/public/crocodiles/1/');
    const contentOK = res.json('name') === 'Bert';
    CounterErrors.add(!contentOK);
}
```

```js
import {Counter} from 'k6/metrics';
import {sleep} from 'k6';
import http from 'k6/http';

const allErrors = new Counter('error_counter');

export const options = {
    vus: 1,
    duration: '1m',
    thresholds: {
        'error_counter': [
            'count<10', // 10 of fewer total errors are tolerated
        ],
        'error_counter{errorType:authError}': [
            'count<=2', // max 2 authentication errors are tolerated
        ],
    },
};

export default function() {
    const auth_resp = http.post('https://test-api.k6.io/auth/token/login/', {
        username: 'test-user',
        password: 'supersecure',
    });

    if (auth_resp.status >= 400) {
        allErrors.add(1, {errorType: 'authError'}); // tagged value creates submetric 
    }

    const other_resp = http.get('https://test-api.k6.io/public/crocodiles/1/');
    if (other_resp.status >= 400) {
        allErrors.add(1);
    }

    sleep(1);
}
```

## Gauge

Gauge is an object for representing a custom metric holding only the latest value added. It is one of the four custom metrics.

|Parameter|Type|Description|
|-|-|-|
|name|string|The name of the custom metric|
|isTime|boolean|A boolean indicating whether the values added to the metric are time values or just uptyped values.|

|Method|Description|
|-|-|
|Gauge.add(value, [tags])|Add a value to the gauge metric. Only the latest added will be kept|

## Rate

Rate is an object for representing a custom metric keeping track of the percentage of added values that are non-zero. It is one of the four custom metrics.

|Parameter|Type|Description|
|-|-|-|
|name|string|The name of the custom metric.|

|Method|Description|
|-|-|
|Rate.add(value, [tags])|Add a value to the rate metric.|

**Rate usage in Threshold**

When Rate is used in a threshold expression, the variable must be called rate (lower case).

- rate < 0.1 // less than 10%
- rate >= 0.9 // more or equal to 90%

The value of the rate variable ranges between 0.00 and 1.00.

```js
import {Rate} from 'k6/metrics';

const myRate = new Rate('my_rate');

export default function() {
    myRate.add(true);
    myRate.add(false);
    myRate.add(1);
    myRate.add(0, {tag1: 'value1', tag2: 'value2'});
}
```

## Trend
