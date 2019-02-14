# header-printer-go
Simple code to return headers as response body, useful for debugging


## Steps to run this container on an EC2 instance
```bash
sudo chmod 400 your.pem
ssh -i /path/to/[your].pem ec2-user@[public-DNS]
sudo yum update -y
sudo yum install -y docker
sudo service docker start
sudo usermod -a -G docker ec2-user
docker info
docker run -d -p 80:80 matheustp/golang-header-printer
docker ps
```
