provider "aws" {
  region = "ap-south-1"
}

# Create an ECS cluster
resource "aws_ecs_cluster" "my_cluster" {
  name = "my-ecs-cluster"
}

# Create a task definition
resource "aws_ecs_task_definition" "my_task_definition" {
  family                = "terratest"
  network_mode          = "awsvpc"
  cpu                   = 256
  memory                = 512
  requires_compatibilities = ["FARGATE"]

  container_definitions = <<-JSON
    [
      {
        "image": "parameshboddeda/node-basic-app:latest",
        "name": "node-basic-app",
        "networkMode": "awsvpc",
        "portMappings": [
          {
            "containerPort": 8888,
            "hostPort": 8888,
            "protocol": "tcp"
          }
        ]
      }
    ]
  JSON
}

data "aws_vpc" "default" {
  default = true
}

data "aws_subnets" "all" {
  filter {
    name   = "vpc-id"
    values = [data.aws_vpc.default.id]
  }
}

data "aws_security_groups" "test" {
  filter {
    name   = "vpc-id"
    values = [data.aws_vpc.default.id]
  }
}

# Create a load balancer
resource "aws_lb" "my_lb" {
  name               = "my-load-balancer"
  load_balancer_type = "application"
  subnets            = data.aws_subnets.all.ids

  security_groups = [data.aws_security_groups.test.ids[0]]
}

# Create a target group
resource "aws_lb_target_group" "my_target_group" {
  name     = "my-target-group"
  port     = 8888
  protocol = "HTTP"
  vpc_id   = data.aws_vpc.default.id
  target_type = "ip"  # or "ip_port"
}

# Create a listener
resource "aws_lb_listener" "my_listener" {
  load_balancer_arn = aws_lb.my_lb.arn
  port              = 80
  protocol          = "HTTP"

  default_action {
    target_group_arn = aws_lb_target_group.my_target_group.arn
    type             = "forward"
  }
}

# Create an ECS service
resource "aws_ecs_service" "my_service" {
  name              = "my-service"
  cluster           = aws_ecs_cluster.my_cluster.arn
  task_definition   = aws_ecs_task_definition.my_task_definition.arn
  desired_count     = 1
  launch_type       = "FARGATE"

  network_configuration {
    security_groups = data.aws_security_groups.test.ids
    subnets         = data.aws_subnets.all.ids

    assign_public_ip = true
  }

  load_balancer {
    target_group_arn  = aws_lb_target_group.my_target_group.arn
    container_name    = "node-basic-app"
    container_port    = 8888
  }
}
