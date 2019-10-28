# Resource Helper

[![N|Solid](https://cldup.com/dTxpPi9lDf.thumb.png)](https://github.com/nguyencatpham/request-handler)

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://github.com/nguyencatpham/request-handler)

Resource helper support for parser data from resource service.

## Getting Started!

  ```
    Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
     }
    client = redis.NewRedis(&options)
  ```

### Installation

Request handler requires [Go](https://golang.org/) v1.13+ to run.

Install the package.

```sh
$ go get -u github.com/onskycloud/go-redis
```

#### Kubernetes + Google Cloud

See [KUBERNETES.md](https://github.com/joemccann/dillinger/blob/master/KUBERNETES.md)


### Todos

 - Write MORE Tests
 - Add Night Mode

License
----

MIT