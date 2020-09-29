# campsite

## How to run using Docker Compose
Prerequisites: 
- Docker (installed and running in the background)

1. Run the services
    ```
    docker-compose up
    ```

Now you should see your API is available at `http://localhost:4444`. Please keep in mind it's just the API. If you want to run the interface too, please visit the `campsite-ui` repo.

## How to run Campsite locally (without Docker Compose).
Prerequisites: 
- Go (Golang) installed (i.e. `brew install golang`)
- Docker (installed and running in the background)

If first time:
1. Clone the repo: `git clone https://github.com/dnielsen/campsite/` (if you haven't installed Go using `brew`, you might need to clone it into `~/go/src`).
2. Go to the cloned directory: `cd campsite`. 
4. Run the script that's gonna run the database using Docker: `./scripts/dev dbstart` (it might take about 15 seconds to run fully).
5. Go to the event service directory: `cd packages/event-service`
6. Run the event service: `go run cmd/main.go`
7. Open another terminal tab/window and go to session service directory: `cd packages/session-service`.
8. Run the session service: `go run cmd/main.go`
9. Open another terminal tab/window and go to speaker service directory: `cd packages/speaker-service`.
8. Run the speaker service: `go run cmd/main.go`

(if already cloned:
```
cd campsite
git checkout master
git pull
```
)

Now you should see your API is available at `http://localhost:4444`. Please keep in mind it's just the API. If you want to run the interface too, please visit the `campsite-ui` repo.


## How to deploy on AWS (EC2)

1. Go to `https://aws.amazon.com/console` and sign in.
2. Once signed in, click `Services` in the upper left corner and select `EC2`.
3. Click `Launch Instance`.
4. Select `Amazon Linux 2`
5. Select `t2.micro` (free tier eligible).
6. Click `Next: Configure Instance Details`.
7. Click `Next: Add Storage`.
8. Click `Next: Add Tags`.
9. Click `Next: Configure Security Group`
10. Click `Add Rule` and set the `Type` to `All traffic` and `Source` to `Anywhere`.
11. Click `Review and Launch`.
12. Click `Launch`.
13. Select `Create a new key` and click `Download Key Pair`.
14. Click on the name of your instance (inside the green box, something like `i-0b39c78…`).
15. Go to the directory where you've saved the `.pem` file (key pair), and say `chmod 400 yourKeyName.pem`.
16. Copy the project into the EC2 machine: `scp -r -i yourKeyName.pem ~/path/to/your/project/root/directory ec2-user@YOUR.EC2.IP.ADDRESS:~/`
17. Connect to the EC2: `ssh -i yourKeyName.pem ec2-user@YOUR.EC2.IP.ADDRESS`
18. Update Linux and install Docker and Docker Compose on your EC2:
```
sudo yum update
sudo yum install docker
sudo curl -L https://github.com/docker/compose/releases/download/1.21.0/docker-compose-`uname -s`-`uname -m` | sudo tee /usr/local/bin/docker-compose > /dev/null
sudo chmod +x /usr/local/bin/docker-compose
sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
```
19. Start Docker: `sudo service docker start`.
20. Start your services: `sudo docker-compose up`

Your app has now been deployed. When you go to your EC2 Dashboard there's `Public IPv4 address` and `Public IPv4 DNS`. You can use either of them to connect to your API. 