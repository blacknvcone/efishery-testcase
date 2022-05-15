# EFISHERY TESTCASE

## Table of Contents

- [Getting Started](#getting_started)
- [API Demo](#api-demo)
- [Prerequisites](#prerequisites)
- [Running](#reference)



## Getting Started <a name = "getting_started"></a>
----------------------------
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See [running](#running) for notes on how to deploy the project on a local system .

## API Demo <a name = "api-demo"></a>
----------------------------

>Auth App  : [http://103.157.27.164:3000](http://103.157.27.164:3000)
>
>Auth App Swagger  : [http://103.157.27.164:3000/swagger](http://103.157.27.164:3000/swagger)


>Fetch App  : [http://103.157.27.164:9090/](http://103.157.27.164:9090/)
>
>Fetch App Swagger  : [http://103.157.27.164:9090/swagger](http://103.157.27.164:9090/swagger)

## Prerequisites<a name = "prerequisites"></a>
----------------------------
What things you need to install the software.

1. Go Binary >= 1.17.x
2. Node >= 16.x
3. Yarn >= 1.x (optional package installation)
4. Docker (optional)

## Running <a name = "running"></a>
----------------------------
### Auth-App [Run Localy]
```
    1. Move current directory into svc-auth
    3. Install depedency module using command "npm install"
    4. Run command "node main.js"
    3. See the output log
```

### Auth-App [Container Based]
```
    1. Move current directory into svc-auth
    2. Build image using command 
       "docker build . -t svc-auth"
    3. After build success, just run image with :
       "docker run --name auth-app --publish 3000:3000 svc-auth"
```

### Fetch-App [Run Localy]
```
    1. Move current directory into svc-fetch
    3. Install depedency module using command "go mod download"
    4. Run command "go main.js"
    3. See the output log
```

### Fetch-App [Container Based]
```
    1. Move current directory into svc-fetch
    2. Build image using command 
       "docker build . -t svc-fetch"
    3. After build success, just run image with :
       "docker run --name fetch-app --publish 9090:9090 svc-fetch"
```
//TODO: Create C4 Context and API.md
