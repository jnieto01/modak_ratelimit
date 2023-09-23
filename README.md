# Challenge of Rate limited for Modak
It is a limited rate microservice that can be implemented by any other service in any language through an API. The client uses a GET to know if it is available to continue. It send the flow identification,  transaction type and user identification for example (flow_id="Notifications", type="News" and user_id="123").

For store data it use KVS called Redis (https://redis.io/). It is a NoSQL database that provide fast access to data and high performance 

[![codecov](https://codecov.io/gh/jnieto01/modak_ratelimit/graph/badge.svg?token=5V28X7F757)](https://codecov.io/gh/jnieto01/modak_ratelimit)


## Signature definition

- Check status of ratelimit microservice 
```
    curl --location 'http://localhost:8080/ping'
```
    

- Get validation of ratelimit from client

```
    curl --location 'http://localhost:8080/v1/ratelimit?user_id=jd123&type=news&flow_id=news'
```

- Example of answer
In case the rate limit has not been reached

```json
{
    "isallowed": true,
    "error": {
        "id": 0,
        "messaje": ""
    }
}
```


<br/>

## Development
- Go version: 1.18
- Dependency Managment: Go modules
- Design: Clean architecture
- Database: KVS (Redis)
- CI with coverage (90%) (See ci-coverage.yml)
- Integration with codecov for analysis of coverage (https://app.codecov.io/gh/jnieto01/modak_ratelimit/tree/develop)
- Handle Secrets for CI (GitHub Actions)
- Setting Docker with environment variables




<br/>

## Table of Contents
- [Installation](#installation)
- [Documentation](#documentation)
- [Examples & Tests](#examples--tests)


<br/>

## Installation



- Install Docker (https://www.docker.com)

- Install DB (Redis) from docker 
```
 docker run -d --name modak-redis-container -p 6379:6379 redis
```

- Check if Redis is running : You can see your container "modak-redis-container" on the list
```
 docker ps
```
 
- Build the app imagen docker
```
 docker build -t modak-ratelimit .
```

- Execute app container 
```
docker run --name modak-container -p 8080:8080 modak-ratelimit
```


When you're done with the test app:

- Stop and Remove the container 
```
docker stop modak-container
docker rm modak-container
```

- Stop and Remove the Redis 
```
docker stop modak-redis-container
docker rm modak-redis-container
```


<br/>

## Documentation


<br/>

## Examples & Tests


<br/>




