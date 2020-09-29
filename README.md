# campsite

## How to run Campsite on your laptop
Prerequisites: 
- Go (Golang) installed (i.e. `brew install golang`)
- Docker (installed and running in the background)

If first time:
1. If not `brew install golang` then need to clone into `~/go/src`
2. `git clone https://github.com/dnielsen/campsite/` 
3. `cd campsite` 
4. Run the script to run the database through Docker: `./scripts/dev dbstart`
5. `cd packages/event-service`
6. `go run cmd/main.go`
7. Do 5. and 6. again for `speaker-service` and `session-service`


If already cloned:
1. `cd campsite`
2. `git checkout master`
3. `git pull`

Now you should see your API is available at `http://localhost:4444`. Please keep in mind it's just the API. If you want to run the interface too, please visit the `campsite-ui` repo.

## How to run using Docker Compose
Prerequisites: 
- Docker (installed and running in the background)

1. Run the services
    ```
    docker-compose up
    ```

Now you should see your API is available at `http://localhost:4444`. Please keep in mind it's just the API. If you want to run the interface too, please visit the `campsite-ui` repo.
