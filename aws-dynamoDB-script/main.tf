provider "aws" {
  region = "ap-south-1"
}

resource "aws_dynamodb_table" "example" {
  name         = "my_table"
  hash_key     = "userId"
  range_key    = "department"
  billing_mode = "PAY_PER_REQUEST"

  server_side_encryption {
    enabled = true
  }
  point_in_time_recovery {
    enabled = true
  }

  attribute {
    name = "userId"
    type = "S"
  }
  attribute {
    name = "department"
    type = "S"
  }

  ttl {
    enabled        = true
    attribute_name = "expires"
  }

  tags = {
    Environment = "production"
  }
}

resource "aws_dynamodb_table_item" "item1" {
  table_name = aws_dynamodb_table.example.name
  hash_key   = aws_dynamodb_table.example.hash_key
  range_key  = aws_dynamodb_table.example.range_key

 
  item = jsonencode({
    "userId": {"S": "1"},
    "department": {"S": "TCB"}
  })
}
resource "aws_dynamodb_table_item" "item2" {
  table_name = aws_dynamodb_table.example.name
  hash_key   = aws_dynamodb_table.example.hash_key
  range_key  = aws_dynamodb_table.example.range_key

 
  item = jsonencode({
    "userId": {"S": "2"},
    "department": {"S": "NBA"}
  })
}
