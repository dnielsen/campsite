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

## ui deployment on Amazon S3

1. Go to the ui directory
    ```
    cd packages/ui
    ```
2. Build the app
    ```
    yarn build
    ```
3. Go to `https://aws.amazon.com/console/` and sign in.
4. Click Services (in the upper top corner) and choose S3.
5. Click Create bucket
6. In Bucket name field enter your bucket name, for example `campsite-ui`. You can change the region if you want
. Click next.
7. Click Next.
8. Uncheck Block all public access and check I acknowledge that .... Click next.
9. Click Create bucket
10. Click on the created bucket's link (in our case it's `campsite-ui`).
11. Click Upload
12. Open the `build` directory we've created during the 2nd step, which is in `packages/ui/build`
13. Drag all the files from the `build` directory onto the Upload page and click Upload.  
14. Click Properties.
15. Click Static website hosting box.
16. Check Use this bucket to host a website, type index.html into the index.html placeholder and click save.
17. Click Overview, select all the files by clicking the first box, click Actions and then Make public, and click the
 Make public button.
 
Your `ui` page has been deployed! Now you can click the `index.html` link, and there's a field Object URL
 with
 the
 link to your website (in our case it's `https://campsite-ui.s3.eu-central-1.amazonaws.com/index.html`). Click it and
  see your results! Remember, we gotta have the backend running as
  well for the
  application to run properly.
