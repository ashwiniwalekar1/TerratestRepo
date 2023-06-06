#!/bin/bash
wget -P /home/ec2-user https://dl.google.com/go/go1.20.4.linux-amd64.tar.gz
cd /home/ec2-user
sudo tar -C /usr/local/ -xzf go1.20.4.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> /home/ec2-user/.profile
source /home/ec2-user/.profile