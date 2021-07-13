# beego-rest-example

## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development purposes.

### Prerequisites
In order to run the project in your local machine you need to have `golang` and framework `beego` docker and `docker-compose` installed.

### Installing bee tool

Install bee tool with the following command:

	go get github.com/beego/bee/v2

Update the bee tool with the following command:

	go get -u github.com/beego/bee/v2

`bee` is installed into `GOPATH/bin` by default. You need to add `GOPATH/bin` to your PATH, otherwise the `bee` command won't work.


### Running
To start the project run
```
docker-compose up -d --build
```
This will start the project in detached mode. If you want to see the logs, run
```
docker-compose logs -f --tail=100 <service-name>
```

To see which containers are running, run
```
docker ps
```
or 
```
docker-compose ps
```

To stops and removes containers and networks created by `up`, run 
```
docker-compose down
