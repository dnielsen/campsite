# Campsite (JSON server)

## Prerequisites
- Node.js
- Yarn (`npm install -g yarn`)

## How to run
1. Go to the app directory
    ```
    cd packages/app
    ```

1. Install the dependencies
    ```
    yarn install
    ```

2. Run the JSON server
    ```
    yarn server
    ```

3. Run the client (in another terminal window)
    ```
    yarn dev
    ```

Now your app is available at `http://localhost:3000`. You can change the event, session, and speakers data by modifying
 the file `db.json`.
