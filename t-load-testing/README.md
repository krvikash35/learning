## locust
* Each `user` instance run in its own green thread then user pick up all the defined `task` randomly.
* Number of user and how the user should increase can be configured through cli. Increase no of user will result in increased throughput.
* Which tasks to pick more or how much to wait between tasks can be configured in locustfile. Using constant_throughput as wait strategies will result in increased throughput. If task lot of time then locust will wait for it before picking next task, so in such case, it won't honour the throughput config at task level. So its better to use no of user rather than task to achieve desired throughput.
* Even if we increase the no of users and api is really slow or blocking(tested with golang main thread sleep) then it might not result in more throughput. Tried to increase the connection pool still same, it seems server is not accepting the connections in this case locust seems to report the real rps instead showing high rps with high error.
* So there are two way to achieve throuhput either by user or task i.e `10 Users hitting api at 1RPS = 1 User hitting api at 10RPS`.
* `wait_time` attribute on user is used to define the behaviour of task throughput. Locust wait for the task to finish then it uses `wait_time` attribute to decide the next task pickup. So if your api takes more time then setting this attribute will not get desired throughput. Instead try to increse the no of user. If u want user to hit api at 5RPS then use `wait_time = constant_throughput`. other supported wait type are
    * constant
    * between
    * constant_throuhput
    * constant_spacing
* each user has its own connection pool, so if are using task rather than user for throughput then u might reach the limit of connection pool soon. i.e if 10 connection per pool then u can have only 10 connection with 1 user but 100 connection with 10 users.


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
    # wait_time = constant_throughput(10)
    pool_manager = PoolManager(maxsize=160, block=True)
    
    @task
    def ping_wait(self):
        self.client.get(url="/ping_wait")
```