# Signals

When you start the server processing with Server.Run(Handler), it will block and wait for incoming signals.

There are two types of signals you can send to a running program to gracefully shutdown the process.

- TSTP: this signal tells Server to stop processing new tasks.
- TERM or INT: This signal tells Server to terminate (i.e. shutdown).

It's recommended you first send TSTP signal to stop processing new tasks and wait for all in-progress tasks to finish before sending TERM signal to terminate the program.

Use kill command to send signals.
 
```sh
kill -TSTP <pid> # stop processing new tasks
kill -TERM <pid> # shutdown the server
```

Note: If you send TERM or INT signal without sending TSTP signal, the Server will start a timer for 8 seconds to allow for all workers to finish (To customize this timeout duration, use ShutdownTime config). If there are workers that didn't finish within that time frame, the task will be transitioned back to pending state and will be processed once the program restarts.
