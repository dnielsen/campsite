# campsite

## How to run Campsite on your laptop
Prerequisites: 
- Golang installed (ie brew install golang)
- Docker Desktop running

If first time:
1. // if not 'brew install golang' then need to clone into ~go/src
2. git clone https://github.com/dnielsen/campsite/ 
2. cd campsite 
5. ./scripts/dev dbstart // Run script from project root folder
6. cd packages/event-service
7. go run cmd/main.go
8. do the same for speaker-service, session-service and ui


If already cloned
 cd campsite
git checkout master
4. git pull

Now you should see your API is available at `http://localhost:4444`. Please keep in mind it's just the API. If you want to run the interface too, please visit the `campsite-ui` repo.

## How to run using Docker Compose
Prerequisites: 
- Docker (installed and running in the background)

1. Run the services
    ```
    docker-compose up
    ```

Now you should see your API is available at `http://localhost:4444`. Please keep in mind it's just the API. If you want to run the interface too, please visit the `campsite-ui` repo.
