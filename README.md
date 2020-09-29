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