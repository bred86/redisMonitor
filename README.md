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
