# Overview

Prometheus is an open-source systems monitoring and alerting toolkit  originally built at SoundCloud. 

Prometheus collects and stores its metrics as time series data, i.e metrics information is stored with the timestamp at which it was recorded, alongside optional key-value pairs called labels.

## Features

- A multi-dimensional data model with time series data identified by metric name and key-value pairs.
- PromQL, a flexible query language to leverage this dimensionality
- no reliance on distributed storage; single server nodes are autonomous
- time series collection happens via a pull model over HTTP
- pushing time series is supported via an intermediary gateway

## What are metrics?

Metrics are numerial measurements in layperson terms. The term time series refers to the recording of changes over time. What users want to measure differs from application to application. 


