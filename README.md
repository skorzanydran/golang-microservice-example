# golang-microservice-example

## Project Structure
```
<service_name>
│   README.md
│   .gitignore    
│
└───cmd
│   │
│   └───<service_name>
│       │   main.go
│   
└───internal
│   │
│   └───database
│   │   │   database.go
│   │
│   └───env
│   │   │   config-local.yml
│   │   │   config-prod.yml
│   │   │   config-test.yml
│   │   │   env.go
│   │
│   └───server
│   │   │   router.go
│   │   │   server.go
│   │
│   └───<entity_name>
│   │   │   controllers.go
│   │   │   models.go
│   │   │   serializers.go
│   │   │   services.go
│   │   │   validators.go
│   │
│   └───<entity2_name>
│   │   │   controllers.go
│   │   │   models.go
│   │   │   serializers.go
│   │   │   services.go
│   │   │   validators.go
│   │
│   │   ...
│
└───pkg
    │
    └───api
```
