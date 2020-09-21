# campsite

## How to run using Docker Compose
Prerequisites: 
- Docker (installed and running in the background)
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
3. Install dependencies
    ```
    yarn install
    ```
4. Run the ui server
    ```
    yarn start
    ```

Now you should see your app at `http://localhost:3000`
 
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