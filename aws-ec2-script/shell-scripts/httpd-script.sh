#!/bin/bash
sudo yum -y update
sudo yum install httpd -y
echo "I Made a Terraform Module" > /var/www/html/index.html
sudo service httpd start
sudo chkconfig httpd on
# sudo sed -i 's/Listen 80/Listen 8080/g' /etc/httpd/conf/httpd.conf