# recoo

Quick and easy deployment for Go and Python

## Usage

Config file `recoo-config.yml`

```sh

build:
    entry: server.go
    ports: 
        - "8086:8086"
release:
    registry: 
        login: motorcode
        url: local
deploy:
    services:
        - 
            image: redis
```



`recoo run` start of the building pipeline

`recoo rm` removing of pipeline
