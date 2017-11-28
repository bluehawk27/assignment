# Go-Redis Proxy

## About
---

This is a simple http proxy Redis written in Go and implemnets an in-memory lru cache layer

## Architecture
---
The application is split up into packages where there is a clear seperation of concerns ie. transport/business logic(package service)/redis. 

* __assignment/main.go__ - is the entry point to the application and fires up an http server with declared routes

* __assignment/httpapi/..__ - The transport layer lives here and only has knowledge of the business logic(service) layer.  The thinking here is that you can just as easily have a grpc or other transport protocol of your choice hit the Service layer

* __assignment/service/..__ - The service layer is where the business logic should be written for the application.  It has no knowledge of the transport layer.  It does communicate with the LRU in memory(Cache) and Redis(store) layers.

* __assignment/store/..__ - The store layer generates the client for the Redis store and implements all of the getters and setters in Redis.

* __assignment/config/..__ - Reads in all the configurables from the config.yml file

* __assignement/cache/..__  - Implements the getters and setters for the in-memory lru cache library

* __assignment/vendor/__ - this project uses glide to manage dependencies. They are all stored in this directory

## Quick Start
___
In your $GOPATH/src/github.com/bluehawk27

This will build and run the all unit tests in the project.

    $ git clone https://github.com/bluehawk27/assignment.git
    $ cd assignment
    $ make deps
    $ make test
    $ make start


You should now be able to use an http Client like Postman and Add/Get to the backing Redis instance via:

#### POST:

`http://127.0.0.1:8082/add/{KEY}`

The Body should have a body. Any type of information that you would like to store.

#### GET:
`http://127.0.0.1:8082/get/{KEY}`


## TODO
---
* Expose Store interface to Mock unit tests dependant on Redis
* Complete Dockerization
* Hash (md5, sha1) Keys so key size remains smaller. Currently Keys can be of any length.


## Time Allocation
---

* Architecture - 1 hr
* Store - 30 min
* Config - 30 min
* CCache - 15 min
* Service - 1 hr
* HTTP API - 20 min
* Docker - 1 Day
* Redis Config - 1 hr

Time complexity of CCache O(logn)
