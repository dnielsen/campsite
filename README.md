# Campsite (JSON server)

## Prerequisites
- Node.js
- Yarn (`npm install -g yarn`)

## How to run
1. Go to the server directory
    ```
    cd packages/server
    ```

2. Install the dependencies
    ```
    yarn install
    ```

3. Run the JSON server
    ```
    yarn dev
    ```

4. Go to the client directory (packages/app)
    ```
    cd packages/app
    ```

5. Install the dependencies
    ```
    yarn install
    ```

6. Run the client (in another terminal window)
    ```
    yarn dev
    ```

Now your app is available at `http://localhost:3000`. You can change the event, session, and speakers data by modifying
 the file `db.json` in the server directory.
