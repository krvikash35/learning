* Go 1.5 is set to make the default value of GOMAXPROCS the same as the number of CPUs on your machine
* Programs that perform parallel computation should benefit from an increase in GOMAXPROCS. However, be aware that concurrency is not parallelism.
* However, if program is meant only for concurrency but not for parallelism might not benefit with multi-core usage since other core can be used by os to schedule other process/thread and just use single thread for main program. 
* Sometime it might even become slower if the program is more of heavy io than being true intrinsic parallel becausing sending data b/w threads involve costly context switching.
* Go's goroutine scheduler is not as good as it needs to be. In future, it should recognize such cases and optimize its use of OS threads. For now, GOMAXPROCS should be set on a per-application basis


## Go scheduler

* `logical CPU(P)`: It is equal to `noOfCore * perCoreHardwareThread`. i.e 8 core intel core i9 has 2 hardware thread per core. So total 16 logical cpus. You can get this by calling `runtime.NumCPU()`. 
* `Kernel/os Thread(M)`: These are small peice of code or function that can run concurrently and managed(scheduled) by kernel space code(osScheduler) for effiecient use of cpu.
    * The GOMAXPROCS variable limits the number of operating system threads that can execute user-level Go code(goroutine) simultaneously. Blocked threads does not count against this GOMAXPROCS limit.
    * Starting Go1.5 it default to no of logical cpu available, prior to this it was set to  1.
    * Each M are assigned a set of goroutine to run using `local run queue(LRN)`.
* `Go Routine(G)`: These are small peice of code or function that can run concurrently and managed(scheduled) by user space code(goScheduler) for efficient use of os thread. These are kind of green thread or co-routine that are managed by user space code than kernel space code. Highly lightweigt and efficient.
    * Each goroutine are

* Every P(logicalCPU) is assigned M(OSThread) & every M is assigned G(GoRoutine).
* Global Run Queue(GRQ) -> It is for Goroutines that have not been assigned to a P yet.
* Local Run Queue(LRQ) -> Each P is given a LRQ that manages the Goroutines.

### Blocking Call (sync system call)
* i.e file i/o
* scheduler create new system thread (`M2`) if current goroutine are blocked then rest other goroutine are run on this new thread. when blocked goroutine is ready then it is put back into LRQ.
### Non Blocking Call (aysnc system call)
* like network i/o using plaform dependent lib like epoll(linux), kqueue(bsd) etc.
* schedule does not need to create new system thread, puts such goroutine on kernel queue(network poller), continue running others from LRQ on `M1`.
### Work Stealing
* goroutine(G) are distributed among multiple threads(M). It make sure that each threads are doing equal work. If one of the threads have finished all the assigned goroutines then scheduler will steal goroutine from other thread and run on this thread.

### Preemptive cooperative
* os thread are preemptively scheduled i.e os will give some time to thread to use cpu but it will preempt once time is finished irrespective of knowing what thread was doing. 
* go routine are mostly cooperative, they will run to completion unless blocked. Once goroutine finish, go will execute other goroutine.
* go does non-cooperative preemption if one of the goroutine takes more time i.e infinite loop `for {}` starting go1.14. Infinite loop — preemption (~10ms time slice) — soft limit
## Go GC(Garbage Collector)

### Explicit/Manual memory management
* In language like c/c++, programmer has to allocate(`calloc`/`malloc`) and deallocate(`free`) memory.
* So there is chance of programmer mistake to forget freeing memory causing `memory leak`. Or free unintended meory causing `crash`.
* Memory are allocated to `stack` using alloc which are auto cleared when function are poped of stack.
* Memory are allocatd to `heap` using malloc and need to be manually freed using free.

### Automatic memory management
* In language like go, java, javascript etc are memory are managed(allocation/deallocation) by lang runtime.
* Compile already knows where to allocate the memory using the language syntax primitive. Some variable have local scope while other are global.
* If compiler detect local variable then it allocate to `stack`, which are automatically freed.
* If compiler detect non-local variable then it allocate to `heap` which has to be later freed by runtime called `garbage collector`.
* In order to know which memory is required or not in heap, it uses graph data structure through reference system.
* Go’s garbage collector is a non-generational concurrent, tri-color mark and sweep garbage collector.
* GC run at fixed regular interval for some duration. This time, program are stoped to run called `stop the world` and put a `write barrier` then worker perform the mark and sweep garbage collection. This result in reducing heap size by freeing memory.
## Go compiler
* go lang has been defined by spec, so there can be many implemnantation compiler. currently two compiler `gc` and `gccgo`
* Gc is the original compiler, and the go tool uses it by default.
* Gccgo is a different implementation with a different focus.
* gccgo is slower to compile but produces more efficient cpu code. henc gccgo compiled code can run faster.
* gc compiler support only popular processor x86 and ARM. While later support SPARC, MIPS etc.


## go  concurrency premitive
* goroutine
* channel(buffered and non-buffered)
* select over channel operation
* range over channel operation
* Atomic counter
* Mutex
* Timeout
* Ticker
* Worker Pool(using goroutine and channels)
* WaitGroups



### channels
* By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.
* Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty. 
* A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
* Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

```go
import "sync/atomic"
var ops uint64
atomic.AddUint64(&ops, 1)
```

```go
ticker := time.NewTicker(1 * time.Second)
for {
    select {
    case t := <-ticker.C:
        fmt.Println("Ticking..")
    case <-time.After(5 * time.Second):
        ticker.Stop()
    }
}
```

```go
ticker := time.NewTicker(1 * time.Second)
for _ = range ticker.C {
    fmt.Println("Ticking..")
}
```

```go
var m sync.Mutex
var counter int
func inc(){
    m.Lock()
    defer m.Unlock()
    counter ++
}

func main(){
    var wg sync.WaitGroup
    
    wg.add(3)
    go inc()
    go inc()
    go inc()

    wg.Wait()
}
```

## References
* https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html
* https://medium.com/@ankur_anand/illustrated-tales-of-go-runtime-scheduler-74809ef6d19b