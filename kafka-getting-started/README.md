
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
* schema registry(Avro, Json, Protobuf)


### Schema registry
Schema and APIs are contract b/w services/teams. Consumer need to know the format(`json`, `avro`, `protobuf` etc) of topic data in order to deserialise it using proper decoder. This is where schema registry help. Now updating schema so that no consumer breaks is different problem that is solved by CI/SOP etc. We should prevent producer from producing incompatible changes.

P1 --> json{"name": "v", age: 12}
P1 --> proto{"v":12}
P1 --> json{"name": "v", age: "12"}

* Producer suddenly started sending in proto format then consumer will break
* Producer changed the data type of age from number to string then consumer will also break

* Some data format has more advantage i.e protobuf is binary format, so more efficient wrt payload size. Its also more strict wrt data types. Json are schema less.
* schema can be send with each event but it will waste network resource, instead keep schema in registry and send the schema ID with each event.
* Schema(some) registry only helps with contract format. Even with schema registry it is easy for producer to break the consumer as mentioned in above examples.
* Avro supports different types of compatibility, such as forward compatible or backward compatible, and data architects can specify the compatibility rules that will be used when validating schemas for each subject
* schema registry are client server based model. `stencil` `confluent-schema-registry` etc are example of it. 
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


## clients
Kafka officially provide sdk for java only. For other language, it is community based. Many of the client are wrapper around `librdkafka` which is c++ lib that implement kafka API.
* java
    * official client: provide lib `kafka-clients`(contain consumer, producer & admin api) , `kafka-streams` provide low level api.
    * 
```
<dependency>
	<groupId>org.apache.kafka</groupId>
	<artifactId>kafka-clients</artifactId>
	<version>3.3.1</version>
</dependency>
```
* go
    * `confluent-kafka-go`: librdkafka wrapper, no schema-registry support
    * `kafka-go`: native.
    * `sarma`: native, no stream support.
* python
    * `confluent-kafka-python`: librdkafka wrapper.