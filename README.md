# Fiber REST API + Socket 
Fiber REST API + Socket 

# Test cases 
```
$ go test -v .

```
# Run 
```
$ ./test.sh
```

# Docker build 
```
$ docker-compose -f docker/docker-compose.yaml  up --build
```


# Docker compiled build for app 
```
$ docker pull anishdhanka/fiber-app
```

# Docker run from build  app 
```
$ docker run -d -p 3000:3000 -e REDIS_HOST='redis' -e  REDIS_PASSWORD=''  -e  REDIS_PORT='6379' -e  POSTGRES_DB='fiber_api' -e  POSTGRES_PASSWORD='test' -e  POSTGRES_USER='postgres_user' --name fiber-api --restart=always  -td  anishdhanka/fiber-app
```