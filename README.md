
# CRUD app with Go

- Written in Go
- Stores data in Gorm with PostgreSQL

# Usage
To list students:
```shell
$ curl 'localhost:8080/students'
```
Response:
```
[
  {
    "id": 1,
    "name": "Cenk",
    "age": 25,
    "CreatedAt": "0001-01-01T01:55:52+01:55",
    "UpdatedAt": "2021-07-06T05:39:55.36294+03:00",
    "DeletedAt": null
  }
]
```
To create students:
```shell
$ curl -XPOST -d '{
    "id":1,
    "name":"Cenk",
    "age": 25
}' 'localhost:8080/students'
```
Response:
```
Status:201 Created
```
To update students:
```shell
$ curl -XPUT -d '{
    "id":1,
    "name":"Deneme",
    "age": 35
}' 'localhost:8080/students/1'
```
Response:
```
Status:200 OK
```
To delete students:
```shell
$ curl -XDELETE 'localhost:8080/students/1'
```
Response:
```
Status:200 OK
```
