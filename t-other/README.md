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
```