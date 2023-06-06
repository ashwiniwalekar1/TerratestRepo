
output "public_ip" {
    value = module.my_server.instance.public_ip
}

# output "sshKey" {
#   value = module.my_server.instance.key_name
# }

# output "eip" {
#   value = module.my_server.eip
# }
