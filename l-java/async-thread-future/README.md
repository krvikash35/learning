* In java, everything is blocking. So if u want to do many task parallely then create as many thread required.
* Main thread should not be blocked.
* If t1  has not join main thread, then main will be exited but app will still be waiting for child thread to complete. JVM continue to run thread until exit method has been called or all non-demon threads have died either by returning or throwing error.



## Thread example
* `join` throws `InterruptedException`, on catching this we should cancle thread to avoid doing unnecessary work.
```java
class MyThread extends Thread{    
    public void run(){  
        for(int i=1;i<5;i++){   
            Thread.sleep(500)
            System.out.println(i)
        }
    }
}

public class App{
    public static void main(String args[]){
        MyThread t1 = new MyThread();
        MyThread t2 = new MyThread();

        t1.start();
        t1.join(); //main will wait for this thread
        Thread.sleep(500); //wait on main thread    
        t2.start();    //main thread will not wait for t2 thread

        
    }
}
```

## ExecutorService
* we can use executor also for waiting for multiple async task. It provide nice api to run tasks in async mode. 
* It also provide a pool of threads adn api to assign task to it.
* Shutdown executor when not required i.e even if app is reached end and threads not doing anything, it will cause JVM to keep running.
* In Java7, developer can use `Fork/Join`

```java
ExecutorService es = Executors.newCachedThreadPool();
for(int i=0;i<5;i++)
    es.execute(new Runnable() { /*  your task */ });
es.shutdown();
boolean finished = es.awaitTermination(1, TimeUnit.MINUTES);
```

```java
Runnable runnableTask = () -> {
    try {
        TimeUnit.MILLISECONDS.sleep(300);
    } catch (InterruptedException e) {
        e.printStackTrace();
    }
};

executorService.execute(runnableTask); //return void, doesn't wait for any result
Future<String> future = executorService.submit(callableTask); // return future for Callable or a Runnable task
String result = executorService.invokeAny(callableTasks); // return the result of any one successfull task
List<Future<String>> futures = executorService.invokeAll(callableTasks); // return list of future when all task are done

executorService.shutdown();
try {
    if (!executorService.awaitTermination(800, TimeUnit.MILLISECONDS)) {
        executorService.shutdownNow();
    } 
} catch (InterruptedException e) {
    executorService.shutdownNow();
}
```

## Future
* get method block the current thread until future is completed
* You can think of this as `await` api in javascript.
* Kotlin, Scala has much better syntactial support for async/await
```java
Future<String> future = executorService.submit(callableTask);
String result = null;
try {
    result = future.get(200, TimeUnit.MILLISECONDS);
} catch (InterruptedException | ExecutionException e) {
    e.printStackTrace();
}

boolean canceled = future.cancel(true);
boolean isCancelled = future.isCancelled();
```

## CompletableFuture
* CompletableFuture class implements the Future interface.
* you can think of it as `async` and `promise`  function in javascript. It gives a way to write async function.
```java
public Future<String> calculateAsync() throws InterruptedException {
    CompletableFuture<String> completableFuture = new CompletableFuture<>();

    Executors.newCachedThreadPool().submit(() -> {
        Thread.sleep(500);
        completableFuture.complete("Hello");
        return null;
    });

    return completableFuture;
}


Future<String> completableFuture = CompletableFuture.completedFuture("Hello");
String result = completableFuture.get();
assertEquals("Hello", result);

CompletableFuture<String> future = CompletableFuture.supplyAsync(() -> "Hello");
CompletableFuture<String> future = completableFuture.thenApply(s -> s + " World");
CompletableFuture<Void> future = completableFuture.thenAccept(s -> System.out.println("Computation returned: " + s));
CompletableFuture<Void> future = completableFuture.thenRun(() -> System.out.println("Computation finished."));
CompletableFuture<Void> combinedFuture = CompletableFuture.allOf(future1, future2, future3);

String combined = Stream.of(future1, future2, future3)
  .map(CompletableFuture::join)
  .collect(Collectors.joining(" "));


CompletableFuture<String> completableFuture = CompletableFuture.supplyAsync(() -> {}).handle((s, t) -> s != null ? s : "Hello Stranger!"); // error handling
```


## Callback
```java
public class Test {
    public static void main(String[] args) throws  Exception {
        new Test().doWork(new Callback() { // implementing class            
            @Override
            public void call() {
                System.out.println("callback called");
            }
        });
    }

    public void doWork(Callback callback) {
        System.out.println("doing work");
        callback.call();
    }

    public interface Callback {
        void call();
    }
}
```

## Fork and Join
```java
public class FactorialSquareCalculator extends RecursiveTask<Integer> {
    private Integer n;
    public FactorialSquareCalculator(Integer n) {
        this.n = n;
    }

    @Override
    protected Integer compute() {
        if (n <= 1) {
            return n;
        }
        FactorialSquareCalculator calculator = new FactorialSquareCalculator(n - 1);
        calculator.fork();
        return n * n + calculator.join();
    }
}

public class App{
    public static void main(String args[]){
        ForkJoinPool forkJoinPool = new ForkJoinPool();
        FactorialSquareCalculator calculator = new FactorialSquareCalculato(10);
        forkJoinPool.execute(calculator);
    }
}
```