# Insert sample data into Redis

``` bash
docker container exec -it simple-rest-api-by-gingonic-gin_redis_1 redis-cli

HMSET event:id_go ID id_go Title "Go" Description "https://golang.org/ (Redis)"
HMSET event:id_gin ID id_gin Title "Gin" Description "https://github.com/gin-gonic/gin (Redis)"
```
