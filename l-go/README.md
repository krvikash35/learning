```
export GOPRIVATE=source.golabs.io
```
```
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)

gvm listall
gvm list
gvm install go1.19
gvm use go1.19
```
Getting started
* package main and function main are necessary to run

```
package main

import "fmt"

func main() {
	fmt.Print("hello")
}
```

package management
* go has always missed the package management. One of the major drawbacks of the previous (pre 1.11) go get was lack of support for managing dependency versions and enabling reproducible builds.
*  go get did not supported module versioning. It always cloned dependency's currnet master branch source code in GOPATH
*  GOPATH did not allowed to have more than one copy of its dependency at a time, nor it allowed to be cloned in project directory.
* Go 1.1 brought module support. go.mod file like package.json & pipfile. machine generated go.sum like package-lock.json
*  go tool run in legacy gopath mode pre go1.1 , post that run in module-aware mode. i.e  go build/install/run automatically add dependency if it has go.mod file


```
go mod init moduleName
go mod tidy // add missing & remove unused module & updated go.mod

go get -d moduleName //download->install, use -ddeprecated to use it for build & install
go install //build->install(main|pkg $GOPATH/[bin|pkg]),i.e compile, recommended build & install
go build // build & leave the output in current directory

go list -m all //list current module & dependency
go list -m -versions github.com/getsentry/raven-go //list all available version of pkg
go get github.com/getsentry/raven-go@v0.1.2 //download specific version & udpate go.mod
```

Golang: Context is Not For Dependency Injection and a SoSolu

Context vs dependency injection
* Context and Dependency Injection a Match not Made in Heaven
* It is quite bulky to properly get the dependency from the context
* unnecessay overhead of passing all dependency to all the request
difficult to test as dep are not explicit and hard to mock
* http://squirly.ca/blog/2016/07/19/golang-context-is-not-for-dependency-injection.html


Project Structure
things to note
* no circular dependency among package
* use dependency injection for testing
* logger, config, db, redis, httprequest inbound: cli->app outbound: app->external
* external service feature should be in its own package with its own contract, model
* our app has many package, each pkg has  many component, any package or component be dependent on this ext package (i.e service.token using salesforce)
* we can layer our app and hence different package for each layer i.e
* pakcage depend on layer(handler,service, store) or function (log,config,auth,db,salesforce)
* some pkg(i.econfig) can be used by almost everyone(pkg or its component) but other pkg(i.e salesforce) might be used by just service.
* avoid circular dep i.e user <-->article i.e there is table user_article who will manage this. in this case, article should own because we are returning article domain
* project structure https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
* project layout https://github.com/golang-standards/project-layout
* clean arch https://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/
* react, golang, pg(mvc) https://github.com/satellity/satellity/tree/master/internal

```
//S01 layer wise
handler -> userHandler, articleHandler
service -> userService, articleService
store -> userStore, articleStore
cons: where to put elasticsearch, as searchService or searchStore?

//S02 module or feature wise
user -> handler, service, store
article -> handler, service, store
pros: no ambugity to put elasticsearch implement it as seprate pkg with its own 
handler, service, store

which pack this reusable function belong? renderResponse, renderErrorResponse
In S01, it belong to handler pkg as used by this pkg only
In S02, it will be used by handlers only but handler are span across different pkg 
so defining this in all pkg(user, aritcle) repeatdly does not make sense(DRY). So
defining in common seperate pkg make sense i.e may be myapp pkg as it provide small
functinality but may be in its own dedicated pkg myhttp

why logger pkg is not passed as dep?
because logger does not add any business func to test

what pkg name if colliding with existing pkg name i.e http?
myhttp, apphttp, httputil, network etc

what all in myhttp pkg?
client:httpClient, ParseHttpResponse, BuildParam, BuildHeaders
server: serve, serveResponse, serveErrorResponse, ParseRequestBody

where should error related stuff be?
error should have: code, description, trace
error should be logged at last point i.e handler as it contain trace but what about
in worker? what about if app crashed before serving handler? should we log at error
creating time? may be

lets say we need to connect to salesforce for creating ticket, whick pkg?
its own sperate pkg and pass it as dependency who need it.

worker?
worker1 takes article from pg to elasticsearch, so it depend on articleStore
and searchStore

elasticsearch should be implemented as searchStore, searchService or elasticsearch?
create seperate pkg search with its own handler,service,store. make es as
dependency to store
```

struct dep or non-dep pattern

should we use struct directly or by passing as dep ? use mixed approach, use both at its own merit. if you don't need much of testing then use directly by importing. Please not both approach need one time initialization otherwise it will be null pointer error. Dep approach make it explicit hence no null error and also easy to test.
```
//config.go
package config
type Config struct{
}
func Load() (*Config,error){
}

//main.go
package main
func main(){
    cfg:=config.Load()
    as:= NewAccountService(cfg)
}
```

```
//config.go
package config
type Config struct{
}
var AppConfig *config
func Load(){
    opt:=getOption()
    appConfig:=&opt
}

//main.go
package main
func main(){
    config.Load()
    as:= NewAccountService()
}

//account.go
package account
import "config"

func Get(){
    port:=config.AppConfig.Port   
}
```

pkg method or struct method pattern

use both pattern where required but struct approach is much interface and test(dep) friendly 
```
//user.go
pakcage user
type User struct{
    name string
}
func(u *User) CreateUser(){
    return db.create(u.name)
}

//otherpackage.go
import user
func someFunc(){
    u:=&user.User{name: "vikash"}
    u.CreateUser()
}
```

```
//user.go
package user
func CreateUser(u *User){
    return db.create(&u.name)
}

//otherpackge.go
import user
func someFunc(){
    user.CreateUser(user.User{name: "vikash"})
}
```

```
//pattern1
package myhttp
var client *http.Client = &http.Client{Timeout: defaultTiemout}
# no need to init, just import and use. You can't pass as dependency

//pattern2
type config struct {
}
var Config *config
func Init() {
	v := viper.New()
	v.Unmarshal(Config)
}
#init in main or init before use, only one instance of config is available, 
you may pass as dependency.


//pattern3
package user
type store struct{
  db *gorm.DB
}
func NewStore() {
	return &store{}
}
func(s *store) CreateUser(){
}

#create new in main or  before use, many instance is possible, you can pass 
as dependency. can't call method on pkg, u have to call through struct object only

```

server request & response

Two way to read request body 2nd one is preferred
1. json.Unmarshal([]bye)  read entire request in one go
2. json.NewDecoder(io.Reader).Decode(interface{}) preferred as give early validation error

Simillary two way to write response body, 2nd one is preffered
1. json.Unmarshal([]bye, interface{})  read entire request in one go
2. json.NewEncoder(io.Writer).Encode(interface{}) preferred as give early validation error


errors

* error as value and error as type are different. 
* error.is(err, somerr) checks for value equality, will work only for error value 
* error.As(err, &someerr) check for type of error

```
//error as value
var AuthError = errors.New("auth error")

//error as type
type AuthError struct{

}
func(a *AuthError) Error() string{

}
```

enums
```
type Direction int

const (
    North Direction = iota
    East
    South
    West
)

type Status int
const (
	created Status = iota + 1
	progres
)
func (s Status) String() string {
	switch s {
	case created:
		return "created"
	case progres:
		return "progress"
	default:
		return ""
	}
}
s1 := created
				s2 := progres
				fmt.Printf("s1 %s", s1)
				fmt.Printf("s1 %v", s1)
				fmt.Printf("s2 %s", s2)
```

should we expose or no expose data atrribute on struct along with method? go convention does not require to  put abstraction on data attribute and if you use pointer receiver then anyway u can manipulate data attribute see here. see logrus example also

```
type oauth struct{
    token string
}

func(o *oauth) GetProfile(){
    return googleProfile(o.token)
}
```
json marshal/unmarshal encoding/decoding

* Package encoding/json unamrshals interface{} type values into a map[string]interface{}


```
//struct-interface to struct
type Customer struct {
    Name string `json:"name"`
}

type UniversalDTO struct {
    Data interface{} `json:"data"`
    // more fields with important meta-data about the message...
}

//approach 1
receivedCustomer := &Customer{}
receivedDTO := UniversalDTO{Data: receivedCustomer}
json.Unmarshal(byteData, &receivedDTO)
    
//approach 2 if data is not available
receivedDTO := &UniversalDTO{}
json.Unmarshal(byteData, receivedDTO)
m := receivedDTO.Data.(map[string]interface{})
customer := Customer{}
if name, ok := m["name"].(string); ok {
    customer.Name = name
}
```

time

```
Epoch, also known as Unix timestamps, is the number of seconds (not milliseconds!) 
that have elapsed since January 1, 1970 at 00:00:00 GMT (1970-01-01 00:00:00 GMT).
time.Now().Add(5 * time.Minute).Unix()

unix millisecond
time.Now().Add(5 * time.Minute).UnixMilli()

always use iso 8601 format: 2019-11-14T00:55:31.820Z

new Date().toISOString()
"2019-11-14T00:55:31.820Z"

new Date().toUTCString()
"Thu, 14 Nov 2019 00:55:16 GMT"

UTC(Coordinated Universal Time) & GMT(Greenwich Mean Time) are time-keeping, 
not format. For common business apps, there is no significant difference, 
literally less than a second’s difference


2019-11-14T00:55:31.820Z
2019-11-14T01:55:31.820+01:00 [Africa/Tunis]
2019-11-13T19:55:31.820-05:00[America/Montreal]
2019-11-14T06:25:31.820+05:30 [Asia/Kolkata]

epoch vs iso
-----------
epoch
+ Compact
+ Easy to do arithmetic actions without any libraries, i.e. var tomorrow=now()+60*60*24
 Not human-readable
- Cannot represent dates before 1 January 1970
- Cannot represent dates after 19 January 2038 (if using Int32)
- Timezone and offset are "external" info, there is ambiguity if the value is UTC or any other offset.
- Officially the spec supports only seconds.
- When someone changes the value to milliseconds for better resolution, there is an ambiguity if the value is seconds or milliseconds.
- Older than ISO 8601 format
- Represents seconds since 1970 (as opposed to instant in time)
- Precision of seconds

iso
+ Human readable
+ Represents instant in time, as opposed to seconds since 1970
+ Newer then Unix time format
+ Specifies representation of date, time, date-time, duration and interval!
+ Supports an offset representation
+ Precision of nanoseconds
- Less compact
- For any arithmetic actions, reach library is required (like java.time.OffsetDatetime)

```

Links

* https://gobyexample.com/
* https://golang.org/doc/effective_go
* https://github.com/golang/go/wiki/CodeReviewComments
* error handling https://middlemost.com/failure-is-your-domain/
* project structure https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
* clean arch https://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/
* Language design https://talks.golang.org/2012/splash.article
* Testing https://www.youtube.com/watch?v=yszygk1cpEc