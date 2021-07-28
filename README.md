
# CRUD app with Go

- Written in Go
- Stores data in Gorm with PostgreSQL

# Quickstart
```shell
git clone https://github.com/cenkayla/crud.git
cd crud
sudo docker-compose up -d
```


# Usage
To list students:
```shell
curl 'localhost:8080/students'
```
Response:
```
[
  {
    "id": 1,
    "name": "Cenk",
    "age": 25
  }
]
```
To create students:
```shell
curl -XPOST -d '{
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
curl -XPUT -d '{
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
curl -XDELETE 'localhost:8080/students/1'
```
Response:
```
Status:200 OK
```
