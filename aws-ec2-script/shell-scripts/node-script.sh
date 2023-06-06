#!/bin/bash
sudo yum update -y
sudo yum install -y git
curl -sL https://rpm.nodesource.com/setup_14.x | sudo bash -
sudo yum install -y nodejs
sudo npm install -g create-react-app --save
create-react-app my-app
cd my-app
npm start
