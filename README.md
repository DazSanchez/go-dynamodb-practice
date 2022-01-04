# Go DynamoDB Practice

Practice project to connect a REST API on Go made with Gin to AWS DynamoDB database.

## Getting started

### Install Go

To run this project you need to install Go and set your Go workspace first. Get Go binaries from the [downloads page](https://go.dev/dl/). 

### Setup AWS Credentials

The project uses environment variables to get AWS keys to connect with AWS DynamoDB, you should follow [these guides](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html) on how to setup the AWS CLI and configure your account.

### Create the Users DynamoDB table

The project assumes you have a `Users` table on your DynamoDB and it has an `ID` partition key. You must create it in order for the queries to work.

### Clone the repo

Clone the repo in your local machine by executing the following command in a terminal:

```shell
$ git clone https://github.com/DazSanchez/go-dynamodb-practice.git

$ cd go-dynamodb-practice
```

### Install dependencies

You must install project's dependencies in order to be able to run it.

```shell
$ go get .
```

### Run the server

Once you have installed the dependendies you can run the server with the following command:

```shell
$ go run .
```

This will run the Gin server on port `8000`.

## Endpoints

Currently there's three endpoints for interacting with DynamoDB:

### Get all users

```shell
$ curl http://localhost:8000/users -H "Accept: application/json" 
```

### Get user by UUID

```shell
$ curl http://localhost:8000/users/<user-id> -H "Accept: application/json" 
```

### Create a user

```shell
$ curl -X POST http://localhost:8000/users \
   -H "Content-Type: application/json" \
   -H "Accept: application/json" \
   -d '{ "name":"my_name", "firstSurname":"my_firstSurname", "email": "my_email" }'
```