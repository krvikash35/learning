## pngpaste
On Mac, paste image from clipboard to folder
```
brew install pngpaste
pngpaste ./arch.png
```
 


## load testing tool
* locust
* vegeta
* pg_bench

```
echo "GET http://localhost:8081/ping" | vegeta attack -duration=5s -rate=10000 | vegeta report

echo "GET http://traps.integration.transport.golabs.io/test?duration_ms=100
x_owner_id: 123
x_owner_type: customer" | ./vegeta attack -duration=10s -rate=500 | ./vegeta report
```

echo "GET http://localhost:8080/test?duration_ms=100
x_owner_id: 123
x_owner_type: customer" | ./vegeta attack -duration=10s -rate=500 | ./vegeta report