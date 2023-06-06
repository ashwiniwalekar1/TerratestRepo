output "instance" {
  value = aws_instance.myinstance02
}

# output "eip" {
#   value  = data.aws_eip.by_allocation_id.public_ip
# }