# Benchmarking MySQL

Benchmarking is an essential skill for MySQL novices and power users alike. A bench-mark, simply put, is a workload designed to stress your system. The usual goal is to learn about the system's behavior, but there are other worthwhile reasons for running benchmarks, such as reproducing a desired system state or burning in new hardware.

Benchmarking is uniquely convenient and effective for studying what happens when you give systems work to do. A bench-mark can help you observe the system's behavior under load, determine the system's capacity, learn which changes are important, or see how your application performs with different data. Benchmarking lets you create fictional circumtances, beyond the real conditions you can observe. 

- Validate your assumptions about the system, and see whether your assumptions are realistic.
- Reproduce a bad behavior you're trying to eliminate in the system.
- Measure how your application currently performs. If you don't know fast it currently runs, you can't be sure any changes you make are helpful. You can also use historical benchmark results to diagnose problems you didn't foresee.
- Simulate a higher load than your production systems handle, to identify the scalability bottleneck that you'll encounter first with growth.
- Plan for growth. Benchmarks can help you estimate how much hardware, network, capacity, and other resources you'll need for your projected future load. This can help reduce risk during upgrades or major application changes.
- Test your application's ability to tolerate a changing environment.

The problem with benchmarking is that it isn't real The workload you use to stress the system is usually very simple in comparison with real-life workloads. There's a reason for that: real-life workloads are nondeterministic, varying, and too complex to understand readily. If you benchmarked your system with real workloads, it would be harder to draw accurate conclusions from the benchmarks.

In what ways is a benchmark's workload unrealistic? There are many artificial dimensions to a benchmark - the data size, the distribution of data and queries - but perhaps the most important is that a benchmark usually runs as fast as it possibly can

## Summary

Everyone who uses MySQL has reasons to learn the basics of benchmarking it. Bench-marking is not just a practical activity for solving business problems, it's also highly educational. Learning how to frame a problem in such a way that a benchmark can help provide an answer is analogous to working from word problems to setting up equations in a match course.

If you haven't done so yet, we recommend at least getting familiar with sysbench. Learn how to use its oltp and fileio benchmarks, if nothing else. The oltp benchmark is very handy for quickly comparing different systems. Filesystem and disk benchmarks. 
