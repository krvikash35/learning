
kafka

Theory
* what is kafka
* what are the usecases
* how does it work
* concepts
* 

Practical
* how to setup
* how to use

* event streaming
* producer and consumer
* partitioned topic
* consumer group
* clients


kafka apis
* admin api
* producer api
* consumer api
* kafka stream api
* kafka connect api

```
brew install kafka
zookeeper-server-start /opt/homebrew/etc/kafka/zookeeper.properties
kafka-server-start /opt/homebrew/etc/kafka/server.properties
kafka-topics --create --topic person --bootstrap-server localhost:9092
kafka-topics --list --bootstrap-server localhost:9092
kafka-console-producer --bootstrap-server localhost:9092 --topic person
kafka-console-consumer --bootstrap-server localhost:9092 --topic person
```

brew kafka bins
* connect-distributed
* connect-standalone
* kafka-cluster
* kafka-configs
* kafka-console-consumer
* kafka-console-producer
* kafka-server-start
* kafka-server-stop
* kafka-storage
* kafka-topics
* zookeeper-server-start
* zookeeper-server-stop

zookeeper: 
* provide centralised service to provide configuration, naming, sycn in large distributed system.
* zookeeper can run in cluster mode.

kafka & zookeeper
* kafka can't run without zookeeper
* There is KPI to remove zookeepr dependency from 2.8.0 and run it in kraft mode
* Kafka uses ZooKeeper to store its metadata about partitions and brokers, and to elect a broker to be the Kafka Controller.
* zookeeper process can run on same broker's node or seperate node




kafka concepts
* producer
* consumer
* broker
* topic
* offset
* partition
* ordering
* consumer group
* stream
* cluster
* kafka controller
* connect
* connector and converter 
* serialzation and deserialization
* retention

kafka setup
* cluster vs non-cluster
* mac: brew, zip, docker
* kubernetes:
* linux:

kafka monitoring, logging and alert
* 