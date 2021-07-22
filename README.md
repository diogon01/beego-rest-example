# beego-rest-example

## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development purposes.

## Prerequisites
In order to run the project in your local machine you need to have `golang` and framework `beego`

## Installing bee tool

Install bee tool with the following command:

	go get github.com/beego/bee/v2

Update the bee tool with the following command:

	go get -u github.com/beego/bee/v2

`bee` is installed into `GOPATH/bin` by default. You need to add `GOPATH/bin` to your PATH, otherwise the `bee` command won't work.


## Database PostgreSQL with docker (optional) 
necessary docker and `docker-compose` installed.
To start Database project
```
docker-compose up -d --build
```

## Running

Before run set database credencials in file `conf/app.conf`
```
sqlconn = postgres://dev:dev@127.0.0.1:5432/gotest
```

This will start the project in Run the application by starting a local development server
```
bee run
```
## Doc for all routes
```
http://localhost:8080/swagger/
```

## Import postman routes (optional)

Get postman routes in the link:


## Populate Database:

Populating the database
```
curl --location --request GET 'http://localhost:8080/v1/deposit/seed'
```
List all deposits
```
curl --location --request GET 'http://localhost:8080/v1/deposit/?sortby=amount&order=desc&limit=1000' \
--header 'Content-Type: application/json' \
--data-raw '{
    "sortby": "amount"
}'
```

Get deposit by id
```
curl --location --request GET 'http://localhost:8080/v1/deposit/1001'
```

Update status deposit from external API:
```
curl --location --request GET 'http://localhost:8080/v1/deposit/status/10'
```


Create deposit
```
curl --location --request POST 'http://localhost:8080/v1/deposit/' \
--header 'Content-Type: application/json' \
--data-raw '{
        "id": 1001,
        "email": "jhylands70@plala.or.jp",
        "txid": "1LD4njKX8PGXCrTiojM1q5NH1gEVux29CW",
        "currency": "PRGX",
        "amount": 986.27,
        "status": "waiting",
        "createdAt": "2021-02-06 14:42:52",
        "updatedAt": "2021-07-06T19:01:44.933218Z"
}'
```

Upate deposit
```
curl --location --request PUT 'http://localhost:8080/v1/deposit/1001' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "teste@gmail.com",
    "txid": "1LD4njKX8PGXCrTiojM1q5NH1gEVux29CW",
    "currency": "PRGX",
    "amount": 986.27,
    "status": "waiting",
    "createdAt": "2021-02-06 14:42:52",
    "updatedAt": "2021-07-06T19:01:44.933218Z"
}'
```