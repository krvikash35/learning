# Microservices

infra
* project
* cluster
* region
* nodes(vm or kubernets based)
* vpc(virtual private cloud)
* vpn(Virtual Private Network)
* vpn tunnel
* vpn gatway
* gatway tunnel
* API Gateway
* load balancer
* dns
* router
* nat router
* firewall


deployment
* vm
* kubernetes: helm
* provider
    * gcp
    * azure
    * aws
* iac
    * terraform
    * chef: knife, cookbook, recipe


config
* zookeeper


proxy
* ngnix
* haproxy
* envoy

api gateway
* kong(based on ngnix)

load balancer
* glb
* l4 vs l7


data format
* json
* protobuf

server
* http
* grpc

client
* Hystrix: circuit breaker, retry
* 


pattern
* server worker
* connection pool
* db migration up/down deployment
* pattern based on language used
* 

software design principal
* language neutral
* DRY(dont repeat yourself)
* KISS (keep it stupid simple)
* tunnel vision: this should not only focus on completing task but on other side effect also.
* should not re-invent the wheel
* maintainable: accommodate changes
* proper error handling,
* proper logging(info, debug, error)
* proper metrics
* Design is not coding and coding is not design
* partitoning: simple, easytounderstand, testable, maintainable
* abstraction: functional abstratction and data abstraction
* modular: 



software design pattern
* creational
    * Abstract factory
    * Builder
    * Dependency Injection
    * Singleton
    * Object pool
    * prototype
* structural
    * Adaptor, wrapper
    * Decorator
    * Module
* Bhavioural
    * observer or publish/subscribe
* concurrency
    * lock
    * join
    * thread pool
    * monitor
    * event loop
    * kernel io wait interupt


logging
* barito
* gcp logger


monitoring
* TICK
* grafana
* prometheous
* new relic
* influx db
* telegraf

alerting
* pager duty
* slack
* email

notes
* proxy can be used as api gatway or reverse proxy