# Random stuff service
This micro service is a golang docker container with 1 GET API that returns a random name with a random joke to the user. Other APIs in the future will also return random information to the user.

## Running the tests
- go test -v auth/authenticator_test.go
- go test -v middleware/ratelimiter/rate_limiter_test.go
- go test -v middleware/recovery/recovery_test.go
- go test -v rest/handlers/handlers_test.go

## Running docker locally
- bash ./scripts/redeploy-docker.sh
- curl http://localhost:5000/random-name-with-joke

## Upcoming features in order to go to production
- Set up auth0 tenant 
    - call access token validator in API handler to gurantee user has right permissions
- CI/CD Pipeline
    - Need unit/integration test and docker deployment jobs
- GET API query params
    - limitTo, firstName, and lastName should be optional query params
- Custom CORS Middleware