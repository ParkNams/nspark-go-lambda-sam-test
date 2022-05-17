# Go Lambda local 테스트 example with sam

## build and test

```shell
go build main.go

sam build

aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws

sam local invoke TestFunction
```
