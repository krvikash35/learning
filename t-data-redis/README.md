
```
brew install redis
brew service info redis
brew service start redis

redis-cli
redis-cli -h localhost -p 6379
redis-cli -c  -> cluster mode


keys "*"
keys "customer*"
RANDOMKEY

get customer_1
set customer_1 abc
```