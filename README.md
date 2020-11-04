# campsite

## How to run using Docker Compose
Prerequisites: 
- Docker (installed and running in the background)

1. Run the services: 
```
docker-compose -f docker-compose.yml down && docker-compose -f docker-compose.yml up --build --force-recreate
```

Now your API (without the React UI) should be available at `http://localhost:1111/api`. You can for example try running `curl http://localhost:1111/api/events` in the terminal.

## How to run Campsite locally (without Docker Compose) (Mac).
Prerequisites: 
- Go (Golang) installed (i.e. `brew install golang`)
- Docker (installed and running in the background)
- The Campsite repo cloned into `~/go/src/`

1. Start the database: `./scripts/dev dbstart` (you might need to stop it if you started it previously: `./scripts/dev dbstop`).
2. Run all the services similarly, for `api` it is: `cd services/api` and then `go run main.go`

Now your API (without the React UI) should be available at `http://localhost:1111/api`. You can for example try running `curl http://localhost:1111/api/events` in the terminal.
Please keep in mind it's just the API. If you wanna run the interface too, please visit the `campsite-ui` repo available at `https://github.com/dnielsen/campsite-ui`.


## How to deploy on AWS EC2 (without the UI) (Mac)

1. Go to `https://aws.amazon.com/console` and sign in.
2. Once signed in, click `Services` in the upper left corner and select `EC2`.
3. Click `Launch Instance`.
4. Select `Ubuntu Server 20.04 LTS (HVM), SSD Volume Type`
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
17. Connect to the EC2: `ssh -i yourKeyName.pem ubuntu@YOUR.EC2.IP.ADDRESS`
18. Update Linux and install Docker, Docker Compose, and Nginx on your EC2:
```
sudo apt-get update
sudo apt-get -y install docker
sudo apt-get -y install nginx
sudo curl -L https://github.com/docker/compose/releases/download/1.21.0/docker-compose-`uname -s`-`uname -m` | sudo tee /usr/local/bin/docker-compose > /dev/null
sudo chmod +x /usr/local/bin/docker-compose
sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
sudo snap install docker
sudo addgroup --system docker
sudo adduser $USER docker
newgrp docker
```
19. Download your repos.
```
git clone https://github.com/dnielsen/campsite
mv campsite api
git clone https://github.com/dnielsen/campsite-ui
mv campsite-ui ui
git clone https://github.com/hypertrace/hypertrace
```
20. Add nginx config
```
cd
sudo mv -f api/nginx/sites-available/default /etc/nginx/sites-available/default
```
21. Start Docker: `sudo service docker start`.
22. Start Nginx: `sudo service nginx start`.
23. Build the ui
```
cd
cd ui
sudo apt-get install -y npm
sudo npm install
sudo npm run build
```
24. Run the api: 
```
cd
cd api
sudo docker-compose up -d
```
25. (optional) Run Hypertrace
```
cd
cd hypertrace/docker
sudo docker-compose -f docker-compose.yml up
```

Your app has just been deployed. When you go to your EC2 Dashboard there's `Public IPv4 address` and `Public IPv4 DNS`. You can use either of them to connect to your API. 

If you have run Hypertrace, it's available at `YOUR.PUBLIC.IP:2020`.


## How to add AWS RDS
1. Create the RDS
2. Replace the database config in `docker.env` with the RDS configuration.
