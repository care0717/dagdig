# workflow-tool
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

## Test
### server
```shell
cd server
go test ./...
```
