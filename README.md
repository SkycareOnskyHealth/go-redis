# Overview
Redis cache with golang
# Support
* Get,Set,Remove
* GetObject, SetObject

# How to use
```
    go get github.com/eneoti/go-redis
```

``` 
    // Options{
    //    Addr:     "localhost:6379",
    //    Password: "", // no password set
    //    DB:       0,  // use default DB
    // }
    client = redis.NewRedis(&options)
```
