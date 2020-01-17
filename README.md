# HTTP monitor
RESTful API to monitor HTTP endpoints with Go
using gorm (An ORM for Go)
and echo (Go web framework)

# Directory Structure
```
.
├── db
│   └── db.go
├── Dockerfile
├── go.mod
├── go.sum
├── handler
│   ├── handler.go
│   ├── request.go
│   ├── routes.go
│   ├── url.go
│   ├── user.go
│   └── user_test.go
├── main.go
├── model
│   ├── url.go
│   └── user.go
├── monitor
│   └── http_monitor.go
├── README.md
├── router
│   ├── middleware
│   │   └── jwt.go
│   └── router.go
├── store
│   ├── url.go
│   └── user.go
└── utils
    └── jwt.go
```

# Models

### User

|  ID  | CreatedAt | UpdatedAt | DeletedAt |  Username  |  Password  | URLs |
|:----:|:---------:|:---------:|:---------:|:----------:|:----------:|:----:|
| uint |  datetime |  datetime |  datetime | string     | string     |  []  |

### URL

|  ID  | CreatedAt | UpdatedAt | DeletedAt |Address|Threshold|SuccessCall|FailedCall| Alert |
|:----:|:---------:|:---------:|:---------:|:-----:|:-------:|:---------:|:--------:|:-----:|
| uint |  datetime |  datetime |  datetime | string|    uint |    uint   |    uint  |Message|

### Message

|  ID  | CreatedAt | UpdatedAt | DeletedAt |  Message   | FailedCall |  RefID  |
|:----:|:---------:|:---------:|:---------:|:----------:|:----------:|:-------:|
| uint |  datetime |  datetime |  datetime |   string   |    uint    |   uint  |

## Installation & Run

```
# Download this project
$ go get "github.com/saman2000hoseini/http-monitor"
# Build and Run
$ cd http-monitor
$ go build
$ ./http-monitor
```

#### Also you can use Dockerfile

```
# Download this project
$ go get "github.com/saman2000hoseini/http-monitor"
$ cd http-monitor
$ docker build -t mydocker .
$ docker run -d -p 8080:8080 mydocker
```

## API

#### /user/register
* `POST` : Register new account

#### Example: (password should be at least 8 characters)
```
$ curl -X POST http://localhost:8080/api/user/register -H 'Content-Type: application/json' -d '{"user":{"username":"user","password":"pass"}}'
```

#### /user/login
* `POST` : Login into your account

#### Example:
```
$ curl -X POST http://localhost:8080/api/user/login -H 'Content-Type: application/json' -d '{"user":{"username":"user","password":"pass"}}'
```

#### /user/update
* `PUT` : Update your account info

#### Example:
```
$ curl -X PUT 'http://127.0.0.1:8080/api/user/update?username=newuser&password=newpass' -H 'Authorization: Bearer <token>'
```

#### /user/url

* `GET` : Get specific url monitored data

#### Example:
```
curl -X GET 'http://127.0.0.1:8080/api/user/url?id=<urlID>' -H 'Authorization: Bearer <token>'
```

* `POST` : Add url to monitor

#### Example:
```
curl -X POST 'http://127.0.0.1:8080/api/user/url?url=<url>&threshold=<threshold>' -H 'Authorization: Bearer <token>' 
```

* `PUT` :  Update url

#### Example:
```
curl -X PUT 'http://127.0.0.1:8080/api/user/url?id=<urlID>&url=<url>&threshold=<threshold>' -H 'Authorization: Bearer <token>'
```

#### /user/url/all
* `GET` : Get all urls monitored data

#### Example:
```
curl -X GET http://127.0.0.1:8080/api/user/url/all -H 'Authorization: Bearer <token>'
```

#### /user/url/:id
* `GET` : Get specific url's daily statistics

#### Example:
```
$ curl -X GET http://127.0.0.1:8080/api/user/url/<urlID> -H 'Authorization: Bearer <token>'
```

