# aMovie Restful API
## Description
WatchaMovie Project is a database for movies with details that you can access and has several other features, A Restful API Project that made for Final Project of Backend Engineer, the owner participates in a Studi Independen, Fullstack Engineer by Alterra Academy, one of the programs created by the Kampus Merdeka and starts in August 2021.
## Technology Used
Framework & Libraries:
* [Echo Labstack](echo.labstack.com): A Golang Web Framework
* [GORM](https://gorm.io/): A library for implementing ORM
* [Viper](https://github.com/spf13/viper): Used for environment configurations
* [vektra/mockery](https://github.com/vektra/mockery): Used for creating mocks for unit testing
* [stretchr/testify](https://github.com/stretchr/testify): Used for unit testing

DBMS & Deployment:
* MySQL: The main database used are relational database, this is where the API stores the entity
* MongoDB: Used for storing log messages from echo logger
* Docker: Open source containerization platform
* Amazon Web Services (AWS) EC2 Continuous Integration and Continuous Deployment (CI/CD) using GitHub Actions is implemented to automate the deployment process.
* MongoDB Atlas: Cloud database for MongoDB, while automating database administration tasks such as database configuration, infrastructure provisioning, patches, scaling events, backups, and more, freeing up developers to focus on what matters to them most.

## Open API
In this project,  I use OMDbAPI to get movies and their details, to find out more about this open api, go to this page http://www.omdbapi.com/

## Clean Architecture
This project implements Clean Architecture. The four layers on the project are:
  * Domain Layer
  * Repository Layer
  * Usecase Layer
  * Controller Layer

## Documentation
[Swagger](https://app.swaggerhub.com/apis/cendiastian/WatchaMovie/1.0.0)
[ERD](https://drive.google.com/file/d/1SOA0obwfwpkLZ4vKznyTmeJzG5S_oRSw/view?usp=sharing)
