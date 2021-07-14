# beego-rest-example

## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development purposes.

### Prerequisites
In order to run the project in your local machine you need to have `golang` and framework `beego` docker and `docker-compose` installed.

###

### Running
To start Database project
```
docker-compose up -d --build
```
This will start the project in Run the application by starting a local development server
```
bee run
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