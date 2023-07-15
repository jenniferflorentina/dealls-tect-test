# Dealls Tech Test

This project is RESTful API service for sign up & login as basic functionalities of the Dating App

## Getting Started

These instructions will give you a copy of the project up and running on
your local machine for development and testing purposes.

### Prerequisites

- Go 1.17+
- Docker
- (Optional) [Ginkgo v2 binary](https://onsi.github.io/ginkgo/#installing-ginkgo)

> Note:
>   Optional means that it is not required

### Structures

    ├── internal                # Source code files
    ├── migrations              # Migration file for database 
    ├── tools                   # Tools and utilities
    ├── main.go                 # main file for running the application
    └── README.md

#### Internal folder

    ├── domain                  # Model, entity, or base DTO that being used in many features
    ├── features                # features or list functional endpoint of the application
    └── server                  # Server setup like database, middleware, logging, etc

To avoid cyclic dependency and make tracing error easier one endpoint will be one folder inside `features` folder, every feature will have 
    - handler file : controller handler
    - use case file : service and logical code for data
    - repository file : persistence or query
    - testing file: testing that is being used is integration testing using gnomock library that uses Docker to create temporary containers for application dependencies, setup their initial state and clean them up in the end.

### Installing

A step by step series of examples that tell you how to get a development
environment running

Start with installing all the dependencies
    
``` sh
go install .
```

And then you can run the project using

``` sh
go run main.go
```

## Running the tests

You can run tests on this project by using go's default test runner

``` sh
go test ./...
```

### Create File & Run Migration

- reference: https://github.com/golang-migrate/migrate
- Install go migrate binary by using the following command: `brew install golang-migrate`
- Just to be safe, load the env variables in your terminal session: `source ./default_env.sh`.
- To generate migration file with auto sequence number, use this command: `make create-migration file=file_name`.
  Example below:

   ```bash
   make create-migration file=create_table_like`
   ```

- To migrate the database, use the command below:

   ```bash
   ./tools/migrate.sh
   ```

### Notes
- Login used JWT Token for stateless login so no need to save access token in database
- Password saved in database is hashed using `bicrypt` library

### Future Development
- Refresh Token