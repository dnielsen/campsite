# campsite

## How to run using Docker Compose
Prerequisites: 
- Docker (installed and running in the background)

1. Run the services
    ```
    docker-compose down && docker-compose up --build
    ```

Now you should see the UI at `http://localhost:3000` and the Event Service API should be available at `http://localhost:4444/events`

## How to transform the monolith into microservice architecture

1. Make 3 copies of `server` directory and call them `event-service`, `speaker-service`, `session-service`.
2. Event service
    1. `internal/service/speaker.go`
        -delete the existing fuctions    
        -instead of calling the database here, it should call the speaker service via HTTP. You can do that using the HTTP client provided by the Go's standard library.
    2. `internal/service/session.go` analogical to the `speaker.go`
    3. Add Speaker and Session Service environment variables to the Config
3. Speaker service
    1. Remove all the handlers besides the speaker handlers and remove the `speakers` part from the handler routes.
    2. Remove all the business logic from `internal/service/event.go` and `internal/service/session.go`. Leave only the `SpeakerInput`, `Speaker`, `EventInput`, `Event` structs.
    3. Adjust the config
    4. (optionally) Remove `RUN mkdir /images` and `COPY --from=builder /images /images` from Dockerfile
4. Session service: analogical to the speaker service
5. Adjust `docker-compose.yml` and `docker.env` appropriately
