# Benchmarking Strategies

There are two primary benchmarking strategies: you can benchmark the application as a whole, or isolate MySQL. We call these two strategies full-stack and single component benchmarking, respectively. There are several reasons to measure the application as a whole instead of just MySQL:

- You're testing the entire application, including the web server, the application code, the network, and the database. This is useful because you don't care about MySQL's performance in particular; you care about the whole application.
- MySQL is not always the application bottleneck, and a full-stack benchmark can reveal this.
- Only by testing the full application can you see how each part's cache behaves.
- Benchmarks are good only to the extent the reflect your actual application's behavior, which is hard to do when you're testing only part of it.

On the other hand, application benchmarks can be hard to create and even harder to set up correctly. If you design the benchmark badly, you can end up making bad decisions, because the results don't reflect reality.

Sometimes, however, you don't really want to know about the entire application. You might just need a MySQL benchmark, at least initially. Such as benchmark is useful if:

- You want to compare different schemas or queries.
- You want to benchmark a specific problem you see in the application.
- You want to avoid a long benchmark in favor of a shorter one that gives you a faster "cycle time" for making and measuring changes.

It's also useful to benchmark MySQL when you can repeat your application's queries against a real dataset. The data itself and the dataset's size both need to be realistic. If possible, use a snapshot of actual production data.

Unfortunately, setting up a realistic benchmark can be complicated and time-consuming, and if you can get a copy of the production dataset, count yourself lucky. It might be impossible - for example, you might be developing a new application that has few users and little data. If you want to know how it'll perform when it grows very large, you'll have no option but to simulate the larger application's data and workload.

## What to Measure

It's best to identify your goals before you start benchmarking - indeed. Before you even design you benchmarks. Your goals will determine the tools and techniques you'll use to get accurate, meaningful results. Try to frame your goals as a questions, such as "Is this CPU better than that one?" or "Do the new indexes work better than the current ones?".

You sometimes need different approaches to measure different things. For example, latency and throughput might require different benchmarks.

**Throughput**

Throughput is defined as the number of transactions per unit of time. This is one of the all-time classics for benchmarking database applications. Standardized benchmarks such as TPC-C are widely quoted, and many database vendors work very hard to do well on them. These benchmarks measure only transactions processing (OLTP) throughput and are most suitable for inter-active multiuser applications. The usual unit of measurement is transactions per second, although it is sometimes transactions per minute.

**Response time or latency**

This measures the total time a task requires. Depending on your application, you might need to measure time in micro-or milliseconds, seconds, or minutes. From this you can derive aggregate response times, such as average, maximum, minimum, and percentiles. Maximum response time is rarely a useful metric, because the longer the benchmark runs, the longer the maximum response time is likely to be. It's also not all repeatable, because it's likely to vary widely between runs. For this reason, it's common to use percentile response times instead. For example, if the 95th percentile response time si 5 milliseconds, you know that the task finishes in 5 milliseconds or less 95% of the time.

**Concurrency**

Concurrency is an important but frequently misused and misunderstood metric. For example, it's popular to say how many users are browsing a website at the same time, usually measured by how many sessions there are. However, HTTP is stateless and most users are simply reading what's displayed in their browsers, so this doesn't translate into concurrency on the web server. Likewise, concurrency on the web server doesn't necessarily translate to the database server; the only thing it directly relates to is how much data your session storage mechanism mus tbe able to handle. A more accurate measurement of concurrency on the web server is how many simultaneous requests are running at any given time.

You can measure concurrency at different places in the application, too. The higher concurrency on the web server might cause higher concurrency at the database level, but the language and toolset will influence this. Be sure that you don't confuse open connections to the database server with concurrency. A well-designed application might have hundreds of connections open to the MySQL server, but only a fraction of these should be running queries at the same time. Thus, a website with 50000 users that a time might require only 10 or 15 simultaneously running queries on the MySQL server!

In other words, what you should really care about benchmarking is the working concurrency, or the number of threads or connections doing work simultaneously running queries on the MySQL server!

**Scalability**


