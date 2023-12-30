# Microservice Auth Example With Go

Basic token based authorization example by using nginx + Go + Redis

---

### Components:

- `nginx-web`: Behaves like an api gateway. For each request, checks if user's token is valid via `go-auth` service
- `go-auth`: Has 2 endpoints: `/` and `/login`. `/` checks if user's token is valid by looking to Redis. `/login` endpoint check if username and password is in the database and returns an auth token to user.
- `nginx-load-balancer`: If user is a valid user, the user can access access services throught load balancer
- `go-service-1`: A hello world service
- `go-service-2`: A hello world service
  
![resim](https://github.com/anilsenay/go-nginx-redis-auth-example/assets/1047345/15a5fd66-c45b-497e-9d91-146d05befe0a)
