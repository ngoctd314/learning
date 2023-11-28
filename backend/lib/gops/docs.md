# Gops

gops is command to list and diagnose Go processes currently running on your system.

## Manual

It is possible to use gops tool both in local and remote mode.

Local model requires that you start the target binary as the same user that runs gops binary. To use gops in a remote mode you need to know target's agent address.

In local most use process's PID as a target; in Remote mode target is a host:port combination.

**Listing all processes running locally**

To print all go processes, run gops without arguments:

```bash
➜  gops git:(master) ✗ gops
29621 29604 gopls  go1.20.4 /home/ubuntu/go/bin/gopls
31894 30377 run  * go1.20.4 /home/ubuntu/code/learning/backend/lib/gops/run
32082 30958 gops   go1.20.4 /home/ubuntu/go/bin/gops
```

## Diagnostics

For processes that start the diagnostics agent, gops can report additional information such as the current stack trace, Go version, memory stats, etc.

```go
func main() {
	if err := agent.Listen(agent.Options{
		ShutdownCleanup: true, // automatically closes on os.Interrupt
	}); err != nil {
		log.Fatal(err)
	}
	var grow [][]byte
	m := sync.Mutex{}

	go func() {
		for i := 0; i < 100000; i++ {
			time.Sleep(time.Millisecond * 100)
			fmt.Println("Goroutine")
			m.Lock()
			grow = append(grow, make([]byte, 100000))
			m.Unlock()
		}
	}()
	for i := 0; i < 100000; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println("Main goroutine")
		m.Lock()
		grow = append(grow, make([]byte, 100000))
		m.Unlock()
	}
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

To report more information about a process a process, run gops followed by a PID:


```txt
gops 31894 2s

parent PID:     30377
threads:        8
memory usage:   0.022%
cpu usage:      0.000%
username:       ubuntu
cmd+args:       ./run
elapsed time:   03:04
local/remote:   127.0.0.1:36799 <-> 0.0.0.0:0 (LISTEN)
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
