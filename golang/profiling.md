pprof

Its official tool for go to profile go application. It support below profiles. i.e we use `heap` profile to debug memory issue.
```
goroutine       — stack traces of all current goroutines
heap            — a sampling of memory allocations of live objects
allocs          — a sampling of all past memory allocations
threadcreate    — stack traces that led to the creation of new OS threads
block           — stack traces that led to blocking on synchronization primitives
mutex           — stack traces of holders of contended mutexes
```

Enable app with pprof over http
```
_ "net/http/pprof"

http.ListenAndServe(":3001", nil)
```

## Heap
Sample Types
```
inuse_space     — amount of memory allocated and not released yet
inuse_objects   — amount of objects allocated and not released yet
alloc_space     — total amount of memory allocated (regardless of released)
alloc_objects   — total amount of objects allocated (regardless of released)
```

Commands
```
curl http://localhost:3001/debug/pprof/heap > heap.out  --> download profile sample

go tool pprof heap.out                                --> analyze in terminal
go tool pprof -http=:3002 .heap.out                    --> analyze in browser
```