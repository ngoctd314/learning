# Gops

gops is command to list and diagnose Go processes currently running on your system.

## Diagnostics

For processes that start the diagnostics agent, gops can report additional information such as the current stack trace, Go version, memory stats, etc.

```go
func main() {
    if err := agent.Listen(agent.Options{}); err != nil {
        log.Fatal(err)
    }
    time.Sleep(time.Hour)
}
```

```bash
$ gops

983   980    uplink-soecks  go1.9   /usr/local/bin/uplink-soecks
52697 52695  gops           go1.10  /Users/jbd/bin/gops
4132  4130   foops        * go1.9   /Users/jbd/bin/foops
51130 51128  gocode         go1.9.2 /Users/jbd/bin/gocode
```

The output displays:

- PID
- PPID
- Name of the program
- Go version used to build the program
- Location of the associated program

```bash
gops <pid> [duration]
```


```txt
gops 3425652 2s

parent PID:     1767834
threads:        19
memory usage:   0.108%
cpu usage:      0.052%
cpu usage (2s): 0.000%
username:       ubuntu
cmd+args:       go run .
elapsed time:   13:50
```

```bash
gops stack (<pid>|<addr>)
```

```bash
gops memstats (<pid>|<addr>)
```

```bash
gops gc (<pid>|<addr>)
```
If you want to force run garbage collection on the target program, run gc. It will block until the GC is completed.
