# campsite

## How to run
1. Run the backend and wait until it finishes:
    ```
    docker-compose up
    ```
2. Run the frontend:
    0. Make sure you're in the project's root directory
    1. Build the image:
        ```
        docker build -t campsite-frontend ./packages/app
        ```
    2. Run the image 
        ```
        docker run -d -p 3000:3000 campsite-frontend ./packages/app
        ```
   
Now you can use see your working app at `localhost:3000`
