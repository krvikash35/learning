* asyncio feature was added in python 3.4
* python can do parallelism(utilize all cpu core) using multiple process. suitable for more cpy heavy tasks.
* python can do concurrency(utilize single core) using multiple co-routine. suitable for more io heavy tasks.
* python can do concurrency(utilize single core) using single thread but asyn. suitable for more io heavy tasks.


## process, threads and co-routines
* proces: scheduled by os. non-shared memory. Parallelism on multi core. concurrent on single core. Each proces has atleast one thread
* native and kernel threads: scheduled by os. non-shared memory. Parallelism on multi core. concurrent on single core. since memory are shared, we to take care of concurrency issue. lightweight compare to process.
* co-routine/go-routine: scheduled by programmer/language. typically (but not necessarily) within a single thread.

## parallel vs concurrent
* parallel programming is different than concurrent programming. If we are doing parallel programming or cpu intensive(i.e machine learning) work then we should utilize multi core.
* typical backend service are more io(network, db, file etc) heavy.
* most language does provide way to create threads but its better to stick with language controlled co[go]-routine for trivial io heavey we app.


## yeild, asyncio, async/await
* asyncio != async/await, asyncio is framework that implement event loop and provide a way to use async/await syntax for async programming.
* any async def function are co-routine. 
* yeild was initial control that can pause the execution of thread

```
import asyncio
# Borrowed from http://curio.readthedocs.org/en/latest/tutorial.html.
@asyncio.coroutine
def countdown(number, n):
    while n > 0:
        print('T-minus', n, '({})'.format(number))
        yield from asyncio.sleep(1)
        n -= 1

loop = asyncio.get_event_loop()
tasks = [
    asyncio.ensure_future(countdown("A", 2)),
    asyncio.ensure_future(countdown("B", 3))]
loop.run_until_complete(asyncio.wait(tasks))
loop.close()
```

```
async def py35_coro():
    await stuff()
```
## References
* https://nothingbutsnark.silvrback.com/how-the-heck-does-async-await-work-in-python-3-5