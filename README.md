# RBH project

This is the pre-interview project implement with domain-driven-design it's split to two context but the comunication is library integration
we can split to microservice later!

<center><img src="./bounded-context.png" alt="" width="900px" height="auto"/></center>

## Installation
    go mod download

## Run with docker
    make docker-up

## Down docker
    make docker-down

## Run with environment variable 
    (edit file .env.default and put the database url) (prepare mysql for :3306)

    make run

## Unit-test and Integration-test
    make unit-test
    
    make integration-test-up

## Down the integration-test sandbox
    make integration-test-down

## Postman collection also provided !!
    rbh.postman_collection.json

## Improvement
    - use env library such as viper to manage env, Ginkgo for more readable test.
    - improve dockerfile for more readable.
    - can change relationship between user-service, interview-appointment-service from library interaction to another comunication
      to be microservice in the future.
    - add log in service layer.