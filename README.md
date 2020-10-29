# campsite

## How to run using Docker Compose
Prerequisites: 
- Docker (installed and running in the background)

1. Run the services: 
```
docker-compose -f docker-compose.yml down && docker-compose -f docker-compose.yml up --build --force-recreate
```

Now your API (without the React UI) should be available at `http://localhost:1111`. 

## How to run Campsite locally (without Docker Compose) (Mac).
Prerequisites: 
- Go (Golang) installed (i.e. `brew install golang`)
- Docker (installed and running in the background)
- Possibly the Campsite repo cloned into `~/go/src/`

1. Start the database: `./scripts/dev dbstart` (you might need to stop it if you started it previously: `./scripts/dev dbstop`).
2. Run all the services similarly, for `api` it is: `cd services/api` and then `go run main.go`

Now your API should be available at `http://localhost:1111`. Please keep in mind it's just the API. If you wanna run the interface too, please visit the `campsite-ui` repo available at `https://github.com/dnielsen/campsite-ui`.


## How to deploy on AWS EC2 (without the UI) (Mac)

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
13. Select `Create a new key`, give it a name, click `Download Key Pair`, and then `Launch Instance`.
14. Click on the name of your instance (inside the green box, like `i-0b39c78…`).
15. Go to the directory where you've saved the `.pem` file (key pair), and run `chmod 400 yourKeyName.pem`.
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

Your app has just been deployed. When you go to your EC2 Dashboard there's `Public IPv4 address` and `Public IPv4 DNS`. You can use either of them to connect to your API. 