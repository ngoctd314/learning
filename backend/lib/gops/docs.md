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

In order to print the current stack trace from a target program, run 

```txt
gops stack 50863
goroutine 35 [running]:
runtime/pprof.writeGoroutineStacks({0x560fe8, 0xc00018e008})
        /usr/local/go/src/runtime/pprof/pprof.go:703 +0x70
runtime/pprof.writeGoroutine({0x560fe8?, 0xc00018e008?}, 0xc000052c00?)
        /usr/local/go/src/runtime/pprof/pprof.go:692 +0x2b
runtime/pprof.(*Profile).WriteTo(0x535010?, {0x560fe8?, 0xc00018e008?}, 0x0?)
        /usr/local/go/src/runtime/pprof/pprof.go:329 +0x14b
```

```bash
gops memstats (<pid>|<addr>)
```

To print the current memory stats, run the following command:

```txt
gops memstats 50863
alloc: 2.24MB (2347384 bytes): alloc represents all bytes allocated by the application.
total-alloc: 2.24MB (2347384 bytes): represets the total bytes allocated since the program started.
sys: 13.00MB (13628680 bytes)
lookups: 0
mallocs: 436
frees: 8
heap-alloc: 2.24MB (2347384 bytes)
heap-sys: 3.59MB (3768320 bytes)
heap-idle: 896.00KB (917504 bytes)
heap-in-use: 2.72MB (2850816 bytes)
heap-released: 832.00KB (851968 bytes)
heap-objects: 428
stack-in-use: 416.00KB (425984 bytes)
stack-sys: 416.00KB (425984 bytes)
stack-mspan-inuse: 43.44KB (44480 bytes)
stack-mspan-sys: 47.81KB (48960 bytes)
stack-mcache-inuse: 14.06KB (14400 bytes)
stack-mcache-sys: 15.23KB (15600 bytes)
other-sys: 1020.07KB (1044547 bytes)
gc-sys: 6.56MB (6881536 bytes)
next-gc: when heap-alloc >= 4.00MB (4194304 bytes)
last-gc: -
gc-pause-total: 0s
gc-pause: 0
gc-pause-end: 0
num-gc: 0
num-forced-gc: 0
gc-cpu-fraction: 0
enable-gc: true
debug-gc: false
```

```bash
gops gc (<pid>|<addr>)
```
If you want to force run garbage collection on the target program, run gc. It will block until the GC is completed.

```bash
gops setgc (<pid>|<addr>)
```

Sets the garbage collection target to certain percentage. The following command sets it to 10%:

```bash
gops setgc (<pid>|<addr>) 10
```

The following command turns off the gc:

```bash
gops setgc (<pid>|<addr>) off
```

**gops version (<pid>|<addr>)**

gops version 52004
go1.20.4

**gops stats (<pid>|<addr>)**

To print the runtime statistics such as number of goroutines and GOMAXPROCS

## Profiling

**pprof**

gops supports CPU and heap pprof profiles. 
