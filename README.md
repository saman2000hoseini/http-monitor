# HTTP monitor
RESTfull API to monitor HTTP endpoints

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

|  ID  | CreatedAt | UpdatedAt | DeletedAt |  Message   |  RefID  |
|:----:|:---------:|:---------:|:---------:|:----------:|:-------:|
| uint |  datetime |  datetime |  datetime |   string   |   uid   |
