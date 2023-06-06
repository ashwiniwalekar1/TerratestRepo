#!/bin/bash
echo 'Hello instance1' > /home/ec2-user/file.txt
# wget https://parmi.s3.ap-south-1.amazonaws.com/ssmKey.pem -P /home/ec2-user/
# sudo chmod 644 /home/ec2-user/ssmKey.pem
# scp -o StrictHostKeyChecking=no -i /home/ec2-user/ssmKey.pem /home/ec2-user/file.txt ec2-user@${aws_instance.myinstance03.public_ip}:~
