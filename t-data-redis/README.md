
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

```
config get requirepass
config set requirepass mypassword
config set requirepass ""      -> disable password
auth mypassword
redis-cli --user default --pass mypassword
redis-cli --pass mypassword 
```

```
> ACL LIST
1) "user default on nopass ~* &* +@all"

all key (~*)
all Pub/Sub channel(&*)
all command (+@all).
```

```
acl list
acl users
acl whoami
acl getuser default
acl deluser bigboxuser

acl setuser bigboxuser on >bigboxpass ~* +@all &*
acl setuser bigboxuser2 on >somecomplexpass ~product:* +@read +@write
acl setuser bigboxuser resetkeys
acl setuser bigboxuser -@all --> remove access to all command
acl setuser bigboxuser >anotherpass
```