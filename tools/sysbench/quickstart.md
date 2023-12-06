# sysbench

sysbench is a scriptable multi-threaded benchmark tool based on LuaJIT. It is most frequently used for database benchmarks, but can also be used to create arbitrarily complex workloads that do not involve a database server.

sysbench comes with the following bundled benchmarks:

- oltp_*.lua: a collection of OLTP-like database benchmarks
- fileio: a filesystem-level benchmark
- cpu: a simple CPU benchmark
- memory: a memory access benchmark 
- threads: a thread-based scheduler benchmark
- mutex: a POSIX mutex benchmark

The sysbench tool can run a variety of "tests" (benchmarks). It was designed to test not only database performance, but also how we a system is likely to perform as a database server. It fact, Peter and Vadim originally designed it to run benchmarks specifically relevant to MySQL performance, even though they aren't actually all MySQL benchmarks. We'll start with some tests

## Usage

```bash
sysbench [options] ... [testname] [command]
```

- testname is an optional name of built-in test (e.g. fileio, memory, cpu, etc.), or a name of one of the bundled Lua scripts (e.g. oltp_read_only), or a path to a custom Lua script.
- command is an optional argument that will be passed by sysbench to the built-in test or script specified on the
