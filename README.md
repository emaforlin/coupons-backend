# Accounts Service

## Setup

Setup the service instace by editing `config.yaml`.

``` yaml
service:
  api: "v0.1"       # api version
  ports:  
    - web: 8014     # http port
    - rpc: 8015     # grpc port
database:
  name: "db-name"  
  user: "db-username"
  password: "changepass"
  host: "localhost:3306"
```

## Endpoints

### HTTP: /{api_vers}/accounts

* **GET** /health
* **POST** /signup
