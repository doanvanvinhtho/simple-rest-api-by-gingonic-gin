# Insert sample data into MongoDB

``` bash
docker container exec -it simple-rest-api-by-gingonic-gin_mongodb_1 mongo

use demo;
db.event.insert({ID:"id_go", Title:"Go", Description:"https://golang.org/ (MongoDB)" });
db.event.insert({ID:"id_gin", Title:"Gin", Description:"https://github.com/gin-gonic/gin (MongoDB)" });
```
