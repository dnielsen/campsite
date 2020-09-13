# campsite

## How to run
    ```
    docker-compose up -d
    ```
   
Now you can use see your working app at `http://localhost:3000`
 
 ## pgadmin
If you want, you can go to `localhost:8080`, there's pgadmin installed so you can easily edit database data there. By default, the pgadmin credentials are:
 - username:`root@root.com`
 - password: `root`. 
 
 Database server:
 - hostname: `host.docker.internal`
 - port: `5432`
 - username: `postgres`
 - password: `postgres`
 
You can configure those values by modifying `docker.env`
