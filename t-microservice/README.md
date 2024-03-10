# Microservices

iaas vs paas vs saas
* iass: gcp, azure, aws
* paas: heroku, redhat openshift, google app engine
* saas: Jira, google doc, microsoft office online

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
* provider(ias)
    * gcp
    * azure
    * aws
* iac
    * terraform
    * chef: knife, cookbook, recipe


config
* zookeeper

distribution
* containerised i.e docker

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

data
* postgres
* kafka
* redis
* rabbitmq
* bigquery
* metabase

notes
* proxy can be used as api gatway or reverse proxy


tools: used in ms network
* Ngnix
* router

tools: used by dev
* ls
* scp
* ssh


prefix
* t-lma (logging monitoring alert) i.e lma-pd, lma-barrito, lma-newrelic
* t-depl: ci/cd, k8, vm, gitops, devops
* t-iaas: iaas-gcp, iaas-aws, iaas-azure, iaas-tools
* t-iac: iac-terraform, iac-chef
* t-infra: vpn, router, proxy, apiGateway
* sf(software) i.e sf-design-pattern, sf-design-principle
* linux: ssh, scp
* net-



** Hystrix **
* enabled
* DefaultTimeout - 1000 - is how long to wait for command to complete, in milliseconds
* requestVolumeThreshold 
    * This property sets the minimum number of requests in a rolling window that will trip the circuit
    * default - 20
* errorThresholdPercentage
    * This property sets the error percentage at or above which the circuit should trip open and start short-circuiting requests to fallback logic
    * default - 50%
* sleepWindowInMilliseconds
    * This property sets the amount of time, after tripping the circuit, to reject requests before allowing attempts again to determine if the circuit should again be closed
    * default - 5000
* rollingWindow
    * default - 10s


* DefaultTimeout = 1000 - DefaultTimeout is how long to wait for command to complete, in milliseconds
* DefaultMaxConcurrent = 10 - DefaultMaxConcurrent is how many commands of the same type can run at the same time
* DefaultVolumeThreshold = 20 -DefaultVolumeThreshold is the minimum number of requests needed before a circuit can be tripped due to health
* DefaultSleepWindow = 5000 -DefaultSleepWindow is how long, in milliseconds, to wait after a circuit opens before testing for recovery
* DefaultErrorPercentThreshold = 50 -DefaultErrorPercentThreshold causes circuits to open once the rolling measure of errors exceeds this percent of requests




https://github.com/afex/hystrix-go/blob/master/hystrix/metric_collector/default_metric_collector.go

```
metrics
    case "success":
		r.Successes = 1
	case "failure":
		r.Failures = 1
		r.Errors = 1
	case "rejected":
		r.Rejects = 1
		r.Errors = 1
	case "short-circuit":
		r.ShortCircuits = 1
		r.Errors = 1
	case "timeout":
		r.Timeouts = 1
		r.Errors = 1
	case "context_canceled":
		r.ContextCanceled = 1
	case "context_deadline_exceeded":
		r.ContextDeadlineExceeded = 1
```

```
    eventType := "failure"
	if err == ErrCircuitOpen {
		eventType = "short-circuit"
	} else if err == ErrMaxConcurrency {
		eventType = "rejected"
	} else if err == ErrTimeout {
		eventType = "timeout"
	} else if err == context.Canceled {
		eventType = "context_canceled"
	} else if err == context.DeadlineExceeded {
		eventType = "context_deadline_exceeded"
	}
    c.reportEvent(eventType)
```

```
SUCCESS	 - execution complete with no errors
FAILURE -	execution threw an Exception
TIMEOUT	- execution started, but did not complete in the allowed time
BAD_REQUEST -	execution threw a HystrixBadRequestException
SHORT_CIRCUITED	- circuit breaker OPEN, execution not attempted
THREAD_POOL_REJECTED - thread pool at capacity, execution not attempted
SEMAPHORE_REJECTED -	semaphore at capacity, execution not attempted
```