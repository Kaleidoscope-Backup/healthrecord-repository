# Karte Health Record Repository

This project provides models and server components implemented in Go and using MongoDB as storage. This uses the Karte Libraries that exist in the utilities folder.

## Structure

    go-graphql-starer
    │   README.md
    │   Config.toml         --- the configuration file for setting server parameter
    │   Dockerfile
    │   Gopkg.lock          --- generated file from dependency management tool, dep
    │   Gopkg.toml          --- generated file from dependency management tool, dep
    │   graphiql.html       --- the html for graphiql which is used for testing query & mutation
    └───context             --- application context like db configuration
    └───handler             --- the handler used for chaining http request like authentication, logging etc.
    └───loader              --- implementation of dataloader for caching and batching the graphql query
    └───model               --- the folder putting struct file
    └───resolver            --- implementation of graphql resolvers
    └───schema              --- implementation of graphql resolvers
    │   │   schema.go       --- used for generate go code from static graphql files inside 'type' folder
    │   │   schema.graphql  --- graphql root schema
    │   └───type            --- folder for storing graphql schema in *.graphql format
    │       └───...         --- graphql schema files in *.graphql format
    └───service             --- services for users, authorization etc.
    └───util                --- utilities
    └───cmd                 --- contains main package

## Requirement

1. Mongo database
2. Golang
3. Karte [Mongo-Lib](https://gitlab.com/karte/mongo-lib) located in Karte Utilities directory

#### NOTE: BEFORE CONTRIBUTING TO THIS, MAKE SURE TO GO TO CONTRIBUTION SECTION

### Getting started

1. Create the file structure in your $GOPATH
    ```bash
    mkdir $GOPATH/src/gitlab.com/karte
    ```

2. Navigate to that new directory
    ```bash
    cd $GOPATH/src/gitlab.com/karte
    ```

3. Clone the project
    ```bash
    git clone https://gitlab.com/karte/healthrecord-repository
    ```

### Usage(Without docker)

1. Install go dependencies
    ```bash
    go get ./...
    ```

2. Install go-bindata (tool used to convert our graphql schema into binary)
    ```bash
    go get -u github.com/kevinburke/go-bindata/...
    ```

3. Run the following command at root directory to generate Go code from .graphql file
    ```bash
    go-bindata -ignore=\.go -pkg=schema -o=schema/bindata.go schema/...
    ```

    OR

    ```bash
    go generate ./schema
    ```
    There would be bindata.go generated under `schema` folder

4. Start the server (Ensure your Mongo database is live and its setting in Config.toml is correct)
    ```bash
    go run ./cmd/health_record_repository/main.go
    ```

### Usage(With docker)

Coming Soon

### Usage(With Minikube)

#### Mongo

docker pull mongo:latest
kubectl run mongo --image=mongo:latest --port=27017
kubectl expose deployment mongo
kubectl get services

#### healthrecord-repo

docker build -t healthrecord-repository:v3 .

kubectl run healthrecord-repository --image=healthrecord-repository:v3 --port=3000

kubectl expose deployment healthrecord-repository --type=LoadBalancer

minikube service healthrecord-repository

kubectl delete deployments healthrecord-repository
kubectl delete service healthrecord-repository

#### Tools

minikube dashboard

kubectl create -f healthrecord-repository-deployment.yaml,healthrecord-repository-service.yaml

### API Documentation

The API docs are served statically from the default endpoint '/'.

For example, to view the docs locally, navigate to: http://localhost:5000

#### Generate New Documentation

```bash
graphdoc -e http://localhost:5000/query -x "Authorization: Bearer <OAUTH_TOKEN>" -o ./doc/schema --force
```

##### NOTE: See below on generating a new OAUTH_TOKEN from Auth0

### Interact with the API locally

1. Download https://github.com/skevy/graphiql-app

2. Set 'GraphQL Endpoint' to "http://localhost:5000/query"

3. Generate new OAUTH_TOKEN from Auth0
    ```curl
        curl --request POST \
        --url https://karte-dev.auth0.com/oauth/token \
        --header 'content-type: application/json' \
        --data '{"client_id":"<CLIENT_ID_FROM_AUTH0>","client_secret":"CLIENT_SECRET_FROM_AUTH0","audience":"https://healthrecord-repository.karte.io","grant_type":"client_credentials"}'
    ```

4. Add Header by Clicking "Edit HTTP Headers" and including the header: 'Authorization: Bearer <OAUTH_TOKEN>'

#### Query

Coming Soon

#### Mutation

Coming soon

### Contribute

To set up your enviroment further to contribute to this repository, you must complete a few steps:

#### Setup Git Hooks (needed only to be done once per project)

1. Make scripts executable
    ```bash
    chmod +x scripts/run-tests.bash scripts/pre-commit.bash scripts/install-hooks.bash
    ```

2. Install Git Hook
    ```bash
    ./scripts/install-hooks.bash
    ```

### Test

We use [Ginkgo](http://onsi.github.io/ginkgo/).  Make sure it is installed in your gopath

##### NOTE: If you are running the Functional Tests, then make sure a mongodb is running

- Run Tests for entire project (Including Functional and Integration)
    ```bash
    go test ./...
    ```

- Run Unit Tests for sub package (pretty print)
    ```bash
    cd ./service/service_test
    go test
    ```

## Reference

### Originally from Go Graphql Starter

[![GitHub license](https://img.shields.io/github/license/OscarYuen/go-graphql-starter.svg)](https://gitlab.com/karte/healthrecord-repository/blob/master/LICENSE)

This project aims to use [graph-gophers/graphql-go](https://github.com/graph-gophers/graphql-go) to build a starter web application. This project has already been used as backend application in production.

In case you need to get called from another frontend side, CORS may needed to be enabled in this application as this project mainly focuses on backend logic at this stage.

This project would be continuously under development for enhancement. Pull request is welcome.

-[graph-gophers/graphql-go](https://github.com/graph-gophers/graphql-go)

-[tonyghita/graphql-go-example](https://github.com/tonyghita/graphql-go-example)

### License - KARTE

Copyright 2018 KARTE

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
