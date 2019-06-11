# godebug-log

This is a small utility to format GODEBUG output in JSON format. Right now it
only handles the case of the `schedtrace` [GODEBUG
option](https://golang.org/pkg/runtime/#hdr-Environment_Variables) but just
raise an issue if you find it interesting and want support for more.

The rationale is that Go binaries are often run with logs to STDOUT in JSON
format, which GODEBUG would screw. This can be fixed with a pipe to `sed` or
similar, but when running in minimal containers, that's a hassle, where this
small utility, statically compiled, can run even on scratch. It also has no
external dependencies and runs quite fast as it doesn't do any regex or similar
matching.

## Sample usage

```
export GODEBUG=schedtrace=1000

./your-go-binary | godebug-log
```

This will print all your normal logs, and also the GODEBUG `schedtrace` with a
format like:

`{"severity":"debug","sched":"SCHED 600086ms: gomaxprocs=16 idleprocs=16 threads=39 spinningthreads=0 idlethreads=25 runqueue=0 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]"}`

The keys `severity`, `debug` and `sched` are all configurables via flags, these
are just the standard values, which are good for Google Cloud Logs at debug
level.
