packer {
  required_plugins {
    amazon = {
      version = ">= 0.0.2"
      source  = "github.com/hashicorp/amazon"
    }
  }
}

data "amazon-ami" "latest-amazon-linux-image" {
  filters = {
    name                = "amzn2-ami-hvm-*-x86_64-gp2"
    root-device-type    = "ebs"
    virtualization-type = "hvm"
  }
  most_recent = true
  owners      = ["amazon"]
  region      = "ap-south-1"
}


source "amazon-ebs" "ec2-user" {
  ami_name      = "my_ami"
  instance_type = "t2.micro"
  region        = "ap-south-1"
  encrypt_boot  = false
  source_ami    = data.amazon-ami.latest-amazon-linux-image.id
  ssh_username  = "ec2-user"
  ssh_timeout   = "30m"
}

build {
  name    = "my-packer"
  sources = ["source.amazon-ebs.ec2-user"]

  # provisioner "shell" {
  #   inline = [
  #     "sudo yum update",
  #     "sudo yum -y install nginx"
  #   ]
  # }

}
