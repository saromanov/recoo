# recoo

Too for quick and easy deployment

## Usage

Config file `recoo-config.yml`

```
build:
    entryfile: server.go
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
