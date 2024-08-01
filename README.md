# go-rest-api

Followed

> https://www.youtube.com/watch?v=1MXIGYrMk80

> https://tharindubalasooriya.medium.com/building-and-deploying-restful-api-with-go-970016d3f841


# Steps to deploy

> Local: 
    export PORT=8085 
    go run main.go

> Server:
    export PORT=8085
    go build -tags netgo -ldflags '-s -w' -o app
    ./app

