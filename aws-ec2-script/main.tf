provider "aws" {
  region = var.avail_zone
}
terraform {
  backend "s3" {
    bucket = "parmi"
    key    = "terraform.tfstate"
    region = "ap-south-1"
  }
}


resource "aws_vpc" "myvpc02" {
  cidr_block = var.vpc_cidr_block

  tags = {
    Name = "${var.env_prefix}vpc02"
  }
}

module "my_subnet_module" {
  source            = "./modules/subnet"
  subnet_cidr_block = var.subnet_cidr_block
  avail_zone        = var.avail_zone
  env_prefix        = var.env_prefix
  vpc_id            = aws_vpc.myvpc02.id
}

module "my_server" {
  source        = "./modules/webserver"
  avail_zone    = var.avail_zone
  env_prefix    = var.env_prefix
  vpc_id        = aws_vpc.myvpc02.id
  instance_type = var.instance_type
  my_ip         = var.my_ip
  subnet_id     = module.my_subnet_module.subnet.id
  vpc_cidr_block = var.vpc_cidr_block
}