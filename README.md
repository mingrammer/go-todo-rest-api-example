# Go Todo REST API Example
A RESTful API example for simple todo application with Go

It is a just simple tutorial or example for making simple RESTful API with Go using **gorilla/mux** (A nice mux library) and **gorm** (An ORM for Go)

## Installation & Run
```bash
# Download this project
go get github.com/mingrammer/go-todo-rest-api-example

# Build and Run
cd go-todo-rest-api-example
go build
./go-todo-rest-api-example
```

## API

#### /projects
* `GET` : Get all projects
* `POST` : Create a new project

#### /projects/:titld
* `GET` : Get a project
* `PUT` : Update a project
* `DELETE` : Delete a project

#### /projects/:titld/archive
* `PUT` : Archive a project
* `DELETE` : Restore a project 

#### /projects/:titld/tasks
* `GET` : Get all tasks of a project
* `POST` : Create a new project

#### /projects/:titld/tasks/:id
* `GET` : Get a project of a project
* `PUT` : Update a project of a project
* `DELETE` : Delete a project of a project

#### /projects/:titld/tasks/:id/complete
* `PUT` : Complete a task of a project
* `DELETE` : Undo a task of a project

## Todo

- [x] Support basic REST APIs.
- [ ] Support Authentication with user for securing the APIs.
- [ ] Make convenient wrappers for creating API handlers.
- [ ] Write the tests for all APIs.
