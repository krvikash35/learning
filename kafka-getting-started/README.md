
Getting started with kafka


Topics

* Theory
    * what is kafka
    * what are the usecases
    * how does it work
    * concepts
* Practical
    * Setup on various platform like mac, kubernetes, docker etc single/cluster mode, with/without zookeeper
    * How to use kafka with various language like go, java etc for producing & consuming.
    * How to use kafka connect
    * How to use kafka stream api




## kafka APIs
* admin api
* producer api
* consumer api
* kafka stream api
* kafka connect api

## Commands

```
brew install kafka
zookeeper-server-start /opt/homebrew/etc/kafka/zookeeper.properties
kafka-server-start /opt/homebrew/etc/kafka/server.properties
kafka-topics --create --topic person --bootstrap-server localhost:9092
kafka-topics --list --bootstrap-server localhost:9092
kafka-console-producer --bootstrap-server localhost:9092 --topic person
kafka-console-consumer --bootstrap-server localhost:9092 --topic person
```

## brew kafka binaries
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

## zookeeper: 
* provide centralised service to provide configuration, naming, sycn in large distributed system.
* zookeeper can run in cluster mode.

## kafka & zookeeper
* kafka can't run without zookeeper
* There is KPI to remove zookeepr dependency from 2.8.0 and run it in kraft mode
* Kafka uses ZooKeeper to store its metadata about partitions and brokers, and to elect a broker to be the Kafka Controller.
* zookeeper process can run on same broker's node or seperate node




## kafka concepts
* client, server, cluster
* producer
* consumer
* broker
* topic
* offset
* partition
* ordering
* consumer group
* stream
* kafka controller
* connect, connector & converter, serialzation & deserialization
* retention


## kafka setup
* cluster vs non-cluster
* mac: brew, zip, docker
* kubernetes:
* linux:

## kafka monitoring, logging and alert
* 


## Problem Statement
create a application that take employee info from command line, publish it it to some topic, consume it from topic and save it file.

input: stdin, ask name and email
output file: employee.txt
topic: employee-log
format: json