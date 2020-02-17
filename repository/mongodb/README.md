# Insert sample data into MongoDB

``` bash
docker container exec -it simple-rest-api-by-gingonic-gin_mongodb_1 mongo

use demo;
db.event.insert({id:"id_go", title:"Go", description:"https://golang.org/ (MongoDB)" });
db.event.insert({id:"id_gin", title:"Gin", description:"https://github.com/gin-gonic/gin (MongoDB)" });
```
