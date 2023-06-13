output "dns_name" {
  value = aws_lb.my_lb.dns_name
}

output "task_definition" {
  value = aws_ecs_task_definition.my_task_definition.arn
}