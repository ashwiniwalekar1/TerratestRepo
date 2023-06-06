resource "aws_security_group" "mysecurity02" {
  name        = "${var.env_prefix}security02"
  description = "Allow TLS inbound traffic"
  vpc_id      = var.vpc_id

  ingress {
    description = "TLS from VPC"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = [var.vpc_cidr_block]
  }

  ingress {
    description = "TLS from VPC"
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "TLS from VPC"
    from_port   = 0
    to_port     = 0
    protocol    = "all"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port       = 0
    to_port         = 0
    protocol        = "-1"
    cidr_blocks     = ["0.0.0.0/0"]
    prefix_list_ids = []
  }

  tags = {
    Name = "${var.env_prefix}security02"
  }
}

data "aws_key_pair" "ssmKey" {
  key_name = "ssmKey"
}

data "aws_ami" "latest-amazon-linux-image" {
  most_recent = true
  owners = ["amazon"]

  filter {
    name = "name"
    values = ["amzn2-ami-hvm-*-x86_64-gp2"]
  }
  filter {
    name = "virtualization-type"
    values = ["hvm"]
  }
}

resource "aws_instance" "myinstance02" {
  ami           = data.aws_ami.latest-amazon-linux-image.id # ap-south-1
  instance_type = var.instance_type

  associate_public_ip_address = true
  key_name                    = data.aws_key_pair.ssmKey.key_name

  vpc_security_group_ids = [aws_security_group.mysecurity02.id]
  subnet_id              = var.subnet_id

  user_data = join("\n",[
    base64encode(file("./shell-scripts/mysql-script.sh")),
    ])

  tags = {
    Name = "TF-Instance"
  }
}

# data "aws_eip" "by_allocation_id" {
#   id = "eipalloc-0dc495d764f7cdf60"
# }

# resource "aws_eip_association" "eip_assoc" {
#   instance_id   = aws_instance.myinstance02.id
#   allocation_id = data.aws_eip.by_allocation_id.id
# }


# resource "aws_instance" "myinstance03" {
#   ami           = data.aws_ami.latest-amazon-linux-image.id # ap-south-1
#   instance_type = var.instance_type

#   associate_public_ip_address = true
#   key_name                    = data.aws_key_pair.ssmKey.key_name

#   vpc_security_group_ids = [aws_security_group.mysecurity02.id]
#   subnet_id              = var.subnet_id

#   tags = {
#     Name = "TF-Instance2"
#   }
# }


