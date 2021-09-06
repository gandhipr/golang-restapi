## Golang-RESTful API Server

### Overview

The goal of this project is to implement a golang based RESTapi framework with in-memory datastore.
The following is an in-depth walkthrough of this project.
This is a demo API, so the "business" intent of it is to support basic CRUD (Create, Read, Update, Delete) operations for an application metadata datastore.

### Requirements

* Build a Golang RESTful API server for application metadata.
* An endpoint to persist application metadata (In memory is fine). The API must support YAML as a valid payload format.
* An endpoint to search application metadata and retrieve a list that matches the query parameters.

### Installation and Run
```bash
# Download project
   git clone https://github.com/gandhipr/golang-restapi.git

# Build, test and run
   make build
   make test
   make run ---> API Endpoint : http://localhost:8080

# Format go source code
   make fmt

# Genarte mock files using mockgen for mocked interfaces
   make generate-mock
```

### Structure
```
├── apiserver         
    ├── apis                    // Handlers for CRUD operations
        ├── deleteapi
        ├── getapi
        ├── postapi
        └── putapi
    ├── datastore               // In-memory datastore and corresponding tests
        ├── store
        ├── store_mock
        └── store_test
    ├── messages                // Error,success messages categorized based on modules - api, datastore, validators.
    ├── samples                 // Sample input yaml files used for testing
    ├── test
        └── apiserver_test      // Ent-to-end apiserver test
    ├── utils                   // Common methods and structs used in the project
    ├── validators              // Validators defined for input struct
    └── apiserver.go            // Application server
```
### Dependencies
* [gin](https://github.com/gin-gonic/gin) : implementing http framework and handlers
* [gomock](https://github.com/golang/mock) and [mockgen](https://github.com/golang/mock) : in-memory datastore used particularly for testing
* [validator](https://github.com/go-playground/validator) : validation framework for input struct
* [testing](https://pkg.go.dev/testing) and [testify](https://github.com/stretchr/testify) : support for automated testing
* [httptest](https://pkg.go.dev/net/http/httptest) : support for http utilities while testing

### CRUD operations and features
#### Scope
* Given below is an example for a yaml that can be used. This project make following criteria to be satisfied:<br />
    - All fields are compulsory (non-empty). <br />
    - Have implemented restrictions on naming conventions for title and version.<br />
    - Should provide valid email-id. <br />
```
## metdata.yaml
title: App1
version: 1.0.1
maintainers:
  - name: Firstname Lastname
    email: apptwo@hotmail.com
company: Upbound Inc.
website: https://upbound.io
source: https://github.com/upbound/repo
license: Apache-2.0
description: |
  ### blob of markdown
  More markdown
```
* A metadata is identified by title and version. It is possible to have different versions of metadata for same title.
* For feasibility and testing purpose, ```yaml file``` and ```binary data format for yaml``` is supported. It is important to specify correct location of the yaml file.

#### Handlers
API Endpoint : [http://localhost:8080](http://localhost:8080)

*  ```POST apiserver/metadata/``` , ```POST apiserver/metadata/:filepath```:<br />
Store metadata in the store. Returns error if metadata with same title and version is already present. It is important to specify ```--data-binary``` when latter request is used.

* ```PUT apiserver/metadata/``` , ```PUT apiserver/metadata/:filepath``` :<br />
Currently, the update call supports only one functionality of updating existing metadata for given title + version combination. Returns error if no such metadata is present in the datastore.  It is important to specify ```--data-binary``` when latter request is used.

* ```GET apiserver/metadata/``` : lists on the metadata in the store.<br />

* ```GET apiserver/metadata/:title``` : lists all the versions of the metadata for a given title.<br />
* ```GET apiserver/metadata/:title/:version``` : returns a specific metadata for a given title, version.<br />

* ```DELETE apiserver/metadata/:title``` : delete all the metadata version for a given title<br />
* ```DELETE apiserver/metadata/:title/:version``` : delete a specific metadata for a given title, version.<br />

### Testing
This project includes tests for 3 major sections:
1. Datastore - unit tests using mock framework <br />
2. Validators - unit tests <br />
3. End-to-end test for the apiserver - tests for http handlers <br />

### Future scope
* Add support for more fine-grained queries/updates for CRUD operations.
* Support Authentication with user for securing the APIs.
* Write the tests for all APIs.
* Make use of a persistent storage.
* Building a deployment process (automated deployment).


