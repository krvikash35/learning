
**Basics**
```
brew install redis
brew service info redis
brew service start redis

redis-cli
redis-cli -h localhost -p 6379
redis-cli -c  -> cluster mode

command.txt - SET MY_KEY "copy pasted the value"
redis-cli < commands.txt


keys "*"
keys "customer*"
RANDOMKEY

get customer_1
set customer_1 abc
del customer_1

info
info server

config get cluster-node-timeout
config set cluster-node-timeout 5000
```
**ACL**

since v6, redis has added ACL. before this, there was "default" user, where app did not had to mandatory pass username("", "default" or no username are all equivalant), they can only pass password for default user if it is set. password for default user was set using command `config set requirepass mypassword`

auth command was also extended to support username
* new form: `auth <username> <password>`
* old form: `auth  <password>`


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


**Cluster**
* `hash tags` -  way to ensure that multiple keys are allocated in the same hash slot for multi-key operations
* N number of master wiht their respective X replicas(optional)
* keys are evenly distributed across hash slot using `crc16` algo
* support sharding(on remove existing node) and rebalancing(on add new node)
```
redis-cli --cluster create 172.20.0.31:6373 172.20.0.32:6374 172.20.0.33:6375 172.20.0.34:6376 172.20.0.35:6377 172.20.0.36:6378 --cluster-replicas 1 --cluster-yes
redis-cli --cluster add-node <new_node_ip:port> <any_cluster_node_ip:port>
redis-cli --cluster rebalance <new_node_ip:port> --cluster-use-empty-masters
redis-cli --cluster check <new_node_ip:port>
redis-cli --cluster reshard 192.168.11.131:6379 --cluster-from <node_id_remove> --cluster-to  <node_id_keep> --cluster-slots 3276 --cluster-yes
redis-cli --cluster del-node <node_ip>:<node:port>  <node_id>
cluster nodes
```

**Sentinel**
