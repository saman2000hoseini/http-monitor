# HTTP monitor
RESTfull API to monitor HTTP endpoints

# Directory Structure
```
|-- db
|   `-- db.go
|-- handler
|   |-- handler.go
|   `-- request.go
|-- model
|   |-- url.go
|   `-- user.go
|-- router
|   |-- middleware
|   |   `-- jwt.go
|   `-- router.go
|-- store
|   |-- url.go
|   `-- user.go
|-- utils
|   |-- http_monitor.go
|   `-- jwt.go
|-- go.mod
|-- go.sum
`-- main.go
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
