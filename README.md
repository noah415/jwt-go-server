# JWT Go Server

This project was inspired by a previous project I started and is meant to be a quick way to start a Go server project while using JWT authentication. Also, the architecture used in this project was strongly inspired by my good friend [mir-mirsodikov](https://github.com/mir-mirsodikov) and this [github repo](https://github.com/golang-standards/project-layout) where Go server layout standards are specified.

This server has some baseline functionality to help get started creating new routes for your server using [gin](https://github.com/gin-gonic/gin) and the [golang-jwt](https://github.com/golang-jwt/jwt) Go packages.

**TODO:** This README is not complete but I hope it is somewhat helpful.

## Getting Started

The easiest way to start this server is to run the `run` make command in the project directory in a termainal.

**NOTE:** all make command examples are ran using gin's debug mode. You can change this to release mode in the .env file when deploying into a production environment.

```bash
$ make run
go run ./cmd/main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

.env file successfully loaded
[GIN-debug] GET    /                         --> github.com/noah415/jwt-go-server/internal/controller.GetHome (5 handlers)
[GIN-debug] POST   /register                 --> github.com/noah415/jwt-go-server/internal/controller.PostRegister (5 handlers)
[GIN-debug] POST   /login                    --> github.com/noah415/jwt-go-server/internal/controller.PostLogin (5 handlers)
[GIN-debug] GET    /auth                     --> github.com/noah415/jwt-go-server/internal/middleware.AuthorizeHandler.func1 (5 handlers)
[GIN-debug] Environment variable PORT="5050"
[GIN-debug] Listening and serving HTTP on :5050
```

If you want to build a Go executable to run, call the following commands in the project directory in a terminal.

```bash
$ make build 
go build -o bin/main cmd/main.go
$ make start
./bin/main
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

.env file successfully loaded
[GIN-debug] GET    /                         --> github.com/noah415/jwt-go-server/internal/controller.GetHome (5 handlers)
[GIN-debug] POST   /register                 --> github.com/noah415/jwt-go-server/internal/controller.PostRegister (5 handlers)
[GIN-debug] POST   /login                    --> github.com/noah415/jwt-go-server/internal/controller.PostLogin (5 handlers)
[GIN-debug] GET    /auth                     --> github.com/noah415/jwt-go-server/internal/middleware.AuthorizeHandler.func1 (5 handlers)
[GIN-debug] Environment variable PORT="5050"
[GIN-debug] Listening and serving HTTP on :5050
```

## Project Layout

On the root layer of the project, there are 3 main folders. For more details on the root level of this project, visit this [link](https://github.com/golang-standards/project-layout).

- cmd
- internal
- bin

### cmd

The cmd folder holds the main.go, which is just the entry point of the whole server. 

### internal

The internal folder holds all of the Go packages used for the rest of the codebase for the server eg. middleware, router creation, and business logic.

#### router

The router folder is where all routers and router groups are created as well as where the middleware is used. Also note that a global variable, router, is created in this package using the gin.Default() function.

All files other than router.go are used to define init route functions named using the standard "Init\<Routename\>Routes". The purpose of this function is to call all controller functions relating to a specific route, and we use this to decrease the amount of routes being declared in the router.go file. It just makes the router.go file look more clean :).

##### router.go

**InitRouter**
This file is the main layer of the router folder. It is where all router init functions are called and middleware is used. The first most important function is the InitRouter() function. This is called from the main.go file in the command folder. The InitRouter() function is responsible for initializing all routes and middleware for the server and calling the router.Run() function (booting up the gin server).

**getRoutes**
The second most important function in this file is the getRoutes() function. This function is called in the InitRouter() function to tell gin what routes exist on the server. This function creates all of the top level `route groups`. More details on route groups can be found in the [gin](https://github.com/gin-gonic/gin) documentation.

### bin

The bin folder simply holds all of the binary/executable files generated from running the `go build` command. This folder is also in the .gitignore file. To create it you must use the `make build` command.

### makefile

The makefile is a very primative makefile that is only used for running Go's build and run commands. Below are the commands and their functions.

#### run

`run` simply calls the `go run` command on the main.go file. This is a quick way to start the server without having to build the binary first.

```bash
make run
```

#### start

`start` attempts to execute the binary file existing in the bin folder. Of course to use this command, you need to have built the binary by using either the makefile `build` command or using the `go build` command manually.

```bash
make start
```

#### build

`build` as stated above, the build command is used to create the binary executable and store it in the bin folder.

```bash
make build
```

#### clean

`clean` is used to remove the bin folder and all of its contents.

```bash
make clean
```

`NOTE: by no means is this README exhaustive of all details pertaining the code in this repo or is this architecture the best by any standard as I am very new to creating Go servers and implementing JWT authentication/authorization.`

## Sources to check out

- https://github.com/mir-mirsodikov
- https://github.com/golang-standards/project-layout
- https://github.com/gin-gonic/gin
- https://github.com/golang-jwt/jwt
- https://dzone.com/articles/secret-rotation-for-jwt-tokens-1