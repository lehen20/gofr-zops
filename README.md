## GoLang Student Management using GoFr Framework 

The docker for database can be run by: 
```sh
docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=password -p 2001:3306 -d mysql:8.0.30
```
For redis: 
```sh
docker run --name gofr-redis -p 6379:6379 -d redis
```
Docker for the these were successfully run on Docker Desktop: 
![image](https://github.com/lehen20/gofr-zops/assets/98393493/6b38a5ac-6e9f-43a9-a7c0-4ac185240328)
