# dagdig
PFN 2019 インターンシップ課題を解くためのレポジトリ

## How to use
### server
```shell
cd server
go run main.go
```
```shell
curl "localhost:8080/api/jobs"
curl "localhost:8080/api/jobs?created=00:00:05"
```
### worker
serverを起動後
```shell
cd worker
go run cmd/main.go
```

## Test
### server
```shell
cd server
go test ./...
```
### worker
```shell
cd worker
go test ./...
```
