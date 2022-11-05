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

## python modules
* `threading` provide higher-level threading interfaces on top of the lower level _thread module.
    * In CPython, due to the Global Interpreter Lock, only one thread can execute Python code at once.
    *  If you want `cpu bound tasks`(use multi-core), better to use multiprocessing or concurrent.futures.ProcessPoolExecutor
    * this threading module is still an appropriate model if you want to run `multiple I/O-bound tasks` simultaneously.
* `concurrent.futures.ThreadPoolExecutor` offers a higher level interface to push tasks to a background thread without blocking execution of the calling thread, while still being able to retrieve their results when needed.
* `queue` provides a thread-safe interface for exchanging data between running threads.
* `asyncio` offers an alternative approach to achieving task level concurrency without requiring the use of multiple operating system threads.

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

## async timeout
```python
import threading
def callback(arg):
    print(f"callback arg: {arg}")
    return


def main():
    timer = threading.Timer(5.0, lambda: callback(arg))
    timer.start()

```

## with
* `with` statement automatically manage the resource like `files`, `locks` etc. If these resource are not released then will casue memory leak. Other approach is `try/catch`
```python
with lock: #no need to call lock.acquire() or lock.release()
    count += 1

with open('file_path', 'w') as file: # no need to call file.close()
    file.write('hello world !')

file = open('file_path', 'w')
try:
    file.write('hello world')
finally:
    file.close()
```

## threading
```python
import threading

class Thread(threading.Thread):
    def __init__(self, t, *args):
        threading.Thread.__init__(self, target=t, args=args)
        self.start()
count = 0
lock = threading.Lock()

def increment():
    global count 
    lock.acquire()
    try:
        count += 1    
    finally:
        lock.release()
   
def bye():
    while True:
        increment()
        
def hello_there():
    while True:
        increment()

def main():    
    hello = Thread(hello_there)
    goodbye = Thread(bye)
    
    while True:
        print count

if __name__ == '__main__':
    main()

```

## References
* https://nothingbutsnark.silvrback.com/how-the-heck-does-async-await-work-in-python-3-5