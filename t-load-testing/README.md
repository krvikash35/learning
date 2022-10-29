## objective
* Create testing environment as possible as close to real production environment.
* Run/Simulate test scenario as possible as close to real world scenario.
* Control the incoming traffic to our system and mock the outgoing traffic.

## Test type: smoke, load, stress, spike, soak
* smoke: simple sanity test with minimal load(1 or 2 user), just to check if things are working.
* load: To check how system will behave as we increase the load.
* stress: It is kind of load testing but it pushes the boundary and tries to find the limit of system.
* spike: type of stress test, sudden increase in load
* soak: It is more of reliability testing. It helps uncover memory leaks, bugs etc. It should be run little longer. i.e if u allocate an object on every request, then this should help find the memory leak.


## locust
* locust is written in python, test are written in python, started in 2011.Using FastHttpUser can generate 3k RPM on single core, it does support concurrency and distributed request generation.It can also create stage shape like k6. It has good minimal output.
* Each `user` instance run in its own green thread then user pick up all the defined `task` randomly.
* Number of user and how the user should increase can be configured through cli. Increase no of user will result in increased throughput.
* Which tasks to pick more or how much to wait between tasks can be configured in locustfile. Using constant_throughput as wait strategies will result in increased throughput. If task lot of time then locust will wait for it before picking next task, so in such case, it won't honour the throughput config at task level. So its better to use no of user rather than task to achieve desired throughput.
* Even if we increase the no of users and api is really slow or blocking(tested with golang main thread sleep) then it might not result in more throughput. Tried to increase the connection pool still same, it seems server is not accepting the connections in this case locust seems to report the real rps instead showing high rps with high error. But `k6` was able to increase RPS without any error, so it seems some issue related config or limit on locust side.
* So there are two way to achieve throuhput either by user or task i.e `10 Users hitting api at 1RPS = 1 User hitting api at 10RPS`.
* `wait_time` attribute on user is used to define the behaviour of task throughput. Locust wait for the task to finish then it uses `wait_time` attribute to decide the next task pickup. So if your api takes more time then setting this attribute will not get desired throughput. Instead try to increse the no of user. If u want user to hit api at 5RPS then use `wait_time = constant_throughput`. other supported wait type are
    * constant
    * between
    * constant_throuhput
    * constant_spacing
* each user has its own connection pool, so if are using task rather than user for throughput then u might reach the limit of connection pool soon. i.e if 10 connection per pool then u can have only 10 connection with 1 user but 100 connection with 10 users.

```
locust --host http://localhost:5001 --users 10 --spawn-rate 10 --run-time 15s --headless -f locustfile_dir
```

```python
import time
from locust import HttpUser, task, between

class QuickstartUser(HttpUser):
    wait_time = between(1, 5)

    @task
    def hello_world(self):
        self.client.get("/hello")
        self.client.get("/world")

    @task(3)
    def view_items(self):
        for item_id in range(10):
            self.client.get(f"/item?id={item_id}", name="/item")
            time.sleep(1)

    def on_start(self):
        self.client.post("/login", json={"username":"foo", "password":"bar"})
```

```python
class LoadTest(FastHttpUser):
    
    @task
    def ping_wait(self):
        self.client.get(url="/ping_wait")
    @task
    def ping(self):
        self.client.get(url="/ping")
```

## k6
* k6 is written in go. Test are written in javascript. Started in 2017.Outputs are bit verbose and confusing. It is bit difficult to assign weightage like in locust to each api even with their executor feature which is bit complex compare to locust.
* it only shows aggregated result,  does not show the output metrics api wise.
* Scenarios are independent from each other and run in parallel, though they can be made to appear sequential by setting the startTime property of each carefully.
* locust was not able to increase the RPS beyound small no, if go backend had `time.sleep` but k6 had no such limit.
* Its not easy to configure the weighted request(like it is in locust) though its doable.
* locust support stages
* locust support conecpet scenarios and executor which are bit complex.
* RPS is different with same input(10vu and 5s) but go(blocked vs nonblocked call in handler).
    * go backend api had `time.sleep()`: 10RPS
    * go backend api had `no sleep`: 71419RPS
* On increasing Vus, RPS increasing even with blocked backend api but not that much. It means, each user is waiting for api response before calling next, same behaviour as locust. But on contrast, in locust even RPS was not increasing on increase of users. http connection is king of tied to user level. If same user want to make many request without waiting then need to see if its possible or not, or if its really required or not.
* simmilar to locust, concurrency is at user level not at api/task level.

```
k6 run --vus 100 --duration 15s script.js
```

```javascript
import http from 'k6/http';

export default function () {
    http.get('http://localhost:5001/ping_wait');
    http.get('http://localhost:5001/ping');
}
```

* if using scenario then ommit the vus and duratin from cli as these are provided in script. `k6 run  script.js`.
```javascript
import http from 'k6/http';

export const options = {
    scenarios: {
        ping: {
            executor: 'per-vu-iterations',
            exec: 'ping',
            vus: 20,
            iterations: 100,
            maxDuration: '30s',
        },
        ping_await: {
            executor: 'per-vu-iterations',
            exec: 'ping_await',
            vus: 20,
            iterations: 200,
            maxDuration: '30s',
        },
    },
};

export function ping() {
    http.get('http://localhost:5001/ping');
}

export function ping_await() {
    http.get('http://localhost:5001/ping_wait');
}
```