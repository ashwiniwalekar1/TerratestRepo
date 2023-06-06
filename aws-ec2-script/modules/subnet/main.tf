
resource "aws_subnet" "mysubnet02" {
  vpc_id            = var.vpc_id
  cidr_block        = var.subnet_cidr_block
  availability_zone = "${var.avail_zone}a"

  tags = {
    Name = "${var.env_prefix}subnet02"
  }
}

resource "aws_internet_gateway" "gw" {
  vpc_id = var.vpc_id

  tags = {
    Name = "${var.env_prefix}igw02"
  }
}

resource "aws_route_table" "route_table" {
  vpc_id = var.vpc_id
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.gw.id
  }
  tags = {
    Name : "${var.env_prefix}route02"
  }

}

resource "aws_route_table_association" "rt_subnet" {

 subnet_id = aws_subnet.mysubnet02.id

 route_table_id = aws_route_table.route_table.id

}