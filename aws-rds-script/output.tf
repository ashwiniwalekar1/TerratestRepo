output "db_instance_id" {
  value = aws_db_instance.myrds.id
}
output "endpoint" {
  value = aws_db_instance.myrds.endpoint
}