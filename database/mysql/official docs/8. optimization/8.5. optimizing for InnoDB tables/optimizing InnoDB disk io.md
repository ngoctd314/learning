# Optimizing InnoDB Disk I/O

If you follow best practices for database design and tuning techniques for SQL operations, but your database is still slow due to heavy disk I/O activity, consider these disk I/O optimizations. If the Unix top tool or the Windows Task Manager shows that the CPU usage percentage with your workload is less than 70%, your workload is probably disk-bound.

- Increase buffer pool size

- Adjust the flush method

- Use a noop or deadline I/O scheduler with antive AIO on Linux

InnoDB uses the asynchronous I/O subsystem (native AIO) on Linux to perform read-ahead and write requests for data file pages.

- Use direct I/O on Solaris 10 for x86_64 architecture

- Use raw storage for data and log files with Solaris 2.6 or later

- Use additional storage devices

- Consider non-rotational storage

- Increase I/O capacity to avoid backlogs

- Lower I/O capacity is flushing does not fall behind

- Store system tablespace files on Fusion-io devices

- Disable logging of compressed pages
