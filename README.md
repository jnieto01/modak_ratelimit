# Challenge of Rate limited for Modak

It is a rate limited  microservice that can be implemented by any other service in any language through an API. The client uses a GET to know if it is available to continue. It send the flow identification,  transaction type and user identification for example (flow_id="Notifications", type="News" and user_id="123"). For internationalization purposes, optional language support is included for message adjustment.

For store data it use KVS called Redis (https://redis.io/). It is a NoSQL database that provide fast access to data and high performance 


<br/>

## Table of Contents
- [Installation](#installation)
- [Development](#development)
- [Signatures](#Signatures)
- [Examples](#examples)
- [Test](#tests)
- [Suggestion](#suggestion)
- [GitFlow](#gitflow)

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



<br/>


## Development
- Go version: 1.18
- Dependency Managment: Go modules
- Design: Clean architecture
- Database: KVS (Redis)
- CI with coverage. (See ci-coverage.yml) ( 30% just for example)
- Unit Test + Integration test. Load test and Regression test were not developed for the challenge. (see  [Suggestion](#suggestion) )
- All unit tests related to kvs (redis) were commented to avoid problems with CI. However, each test was satisfactorily tested locally.
- Integration with codecov for analysis of coverage (https://app.codecov.io/gh/jnieto01/modak_ratelimit/tree/develop)
- Handle Secrets for CI (GitHub Actions)
- Setting Docker with environment variables
- Support api version for microservice 



<br/>


## Signatures

- Check status of ratelimit microservice 
```
    curl --location 'http://localhost:8080/ping'
```
    

- Get validation of ratelimit from client using language (optional)

```
    curl --location 'http://localhost:8080/v1/ratelimit?user_id=jd123&type=news&flow_id=news&lang=en'
```


- Get validation of ratelimit from client without language (use English by default)

```
    curl --location 'http://localhost:8080/v1/ratelimit?user_id=jd123&type=news&flow_id=news'
```

- All configurations are in one file including rate limit rules (see config/config.go). The times indicated in the json are minutes and are adjusted to what was requested in the challenge. These rules are fully configurable

```json
 "RateLimitRules": [
        {
            "flowid": "notifications",
            "settings": [
                {
                    "key": "status",
                    "maxrequests": 2,
                    "timeinterval": 1
                },
                {
                    "key": "news",
                    "maxrequests": 1,
                    "timeinterval": 1440
                },
                {
                    "key": "marketing",
                    "maxrequests": 3,
                    "timeinterval": 60
                }
            ]
        },
        {
            "flowid": "another_example",
            "settings": [
                {
                    "key": "tes1",
                    "maxrequests": 1,
                    "timeinterval": 1
                },
                {
                    "key": "test2",
                    "maxrequests": 10,
                    "timeinterval": 1
                },
                {
                    "key": "test3",
                    "maxrequests": 20,
                    "timeinterval": 1
                }
            ]
        }
    ]
```

<br/>

## Examples

- Request to RateLimit microservice from Client. Sending the language parameter is optional. However, you can try changing it between the options (en - es) to see the setting of the response messages for the purpose of internationalization of the app 

```
curl --location 'http://localhost:8080/v1/ratelimit?user_id=test123&type=news&flow_id=notifications&lang=en'
```

Answer to Client from Rate Limit microservice
- Successful
```json
{
    "isallowed": true,
    "error": {
        "id": 0,
        "messaje": ""
    }
}
```

- Suspended
```json
{
    "isallowed": false,
    "error": {
        "id": 20,
        "messaje": "Temporarily suspended service for: 15 (Minutes) 55 (Seconds)"
    }
}
```

<br/>

## Test

- CI with GitHub Actions ( https://github.com/jnieto01/modak_ratelimit/actions )
- Metrics and analytics ( https://app.codecov.io/gh/jnieto01/modak_ratelimit )


- Manual verification of test coverage. Remember to have the docker redis container running to have access to the DB
```
go test -cover ./...
```

<br/>

## Suggestion

This implementation has been developed with a focus on the challenge. However, for a real purpose you should take the following considerations:

- Assign password to the KVS database in redis
- Tracking system for issues like Kibana (https://www.elastic.co/es/kibana)
- Tracking system for metric of product like Datadog (https://www.datadoghq.com)
- Handle token JWT 
- Load testing for concurrency
- Regression test
- End to end testing
- Storage all sensitive data safely such as secrets
- Pentest

<br/>

## GitFlow

Branches used for this challenge

- Main Branches

main: The main branch represents the latest stable and deployable version of the application. Every commit to main must be a production-ready version.

develop: The develop branch is the main development branch. New features and changes are merged into this branch while they are being developed and tested.

- Feature Branches
feature/feature-name: To develop a new feature, create a branch from develop. For example, feature/login to implement the login feature.

Work on feature branch, commit and test.

Once the feature is complete, create a Pull Request to merge the feature branch into develop.

- Release Branches
release/vX.X.X: When a new release is being prepared, creates a release branch from develop. For example, release/v1.0.0.

Perform final testing and bug fixes on the release branch.

Once the release is ready for deployment, merge the release branch into main and develop.

- Use of Tags
To mark specific versions in master, use Git tags. For example, v1.0.0.

