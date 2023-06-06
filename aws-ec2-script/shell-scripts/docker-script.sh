#!/bin/bash
sudo yum update -y && sudo yum install -y docker
sudo systemctl start docker
sudo usermod -aG docker ec2-user
# docker run -p 8080:80 nginx
# sudo yum install -y git
# git clone https://github.com/parameshBoddeda/node-basic-app.git
