# Challenge of Rate limited for Modak
It is a limited rate microservice that can be implemented by any other service in any language through an API. The client uses a GET to know if it is available to continue. It send the flow identification,  transaction type and user identification for example (flow_id="Notifications", type="News" and user_id="123")

## Signature definition

- For Checking if the microservice is running
    curl --location 'http://localhost:8080/ping'

- For get ratelimit from client
    curl --location 'http://localhost:8080/ping?flow_id=notifications&user_id=jd123&type=news'

## Implementation considerations
- Setting CI with coverage (90%) on main and develop repos
- Handle Secrets for CI (GitHub Actions)
- Setting Docker with environment variables
- Implement scaffolding with clean code architecture



<br/>

## Table of Contents
- [Installation](#installation)
- [Documentation](#documentation)
- [Examples & Tests](#examples--tests)


<br/>

## Installation
 
docker build -t modak .

<br/>

## Documentation


<br/>

## Examples & Tests


<br/>




