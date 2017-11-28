# About it
It's a work in progress. I'm using it to learn GO.
What does it do?
It monitors Redis (it's not a Beat).
It pushs to Redis these fields:

* key name (and size)
* used memory
* total memory allocated to Redis (0 if it's not set)
* Redis IP
* hostname
* date (to input into Elasticsearch)
* app name (in case you need it)
* team (in case you need it)
* cluster tye (in case you need it)


# Compile it

```
cd src/github.com/bred86/redisMonitor
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
```


# Build it

```
docker build -t redismonitr .
```


# Run it

```
docker run -it --rm redismonitr
```
