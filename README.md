# campsite

## How to run using Docker Compose
Prerequisites: 
- Docker (installed and running in the background)
- Docker Compose (on Mac/Windows it's included as part of Docker Desktop)
- Node.js (v12 or later)
- Yarn (`npm install -g yarn`)

1. Run the services
    ```
    docker-compose up
    ```
2. Go to the `ui` directory (in another terminal window/tab)
    ```
    cd packages/ui
    ```
3. Install the dependencies
    ```
    yarn install
    ```
4. Run the ui server
    ```
    yarn start
    ```

Now you should see your app at `http://localhost:3000`
 
 ## Sample pages
 Events page: `http://localhost:3000`
 Event by id page: `http://localhost:3000/events`
 Session by id page: `http://localhost:3000/sessions/be13940b-c7ba-4f97-bdab-b4a47b11ffed`
 Speaker by id page: `http://localhost:3000/speakers/9c08fbf8-160b-4a86-9981-aeddf4e3798e`
 
 ## pgadmin
If you want, you can go to `localhost:8080`, there's pgadmin installed, so you could easily edit database data there. By default, the pgadmin credentials are:
 - username:`root@root.com`
 - password: `root`. 
 
 Database server:
 - hostname: `host.docker.internal`
 - port: `5432`
 - username: `postgres`
 - password: `postgres`
 
You can configure those values by modifying `docker.env`
