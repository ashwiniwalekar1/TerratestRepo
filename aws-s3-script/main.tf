provider "aws" {
  region = "ap-south-1"
}

resource "aws_s3_bucket" "parmi-s3" {
  bucket = "parmi-s3"

  tags = {
    Name = "parmi-s3"
    Environment = "dev"
  }
}

# resource "aws_s3_bucket_acl" "myacl" {
#   bucket = aws_s3_bucket.parmi-s3.id
#   acl = "public-read"
# }


# locals {
#   folder_path = "D:/html"
# }

resource "aws_s3_object" "example" {
  # for_each = fileset(local.folder_path, "**/*")

  bucket = aws_s3_bucket.parmi-s3.id
  key = "index.html"
  source = "index.html"
  # key    = each.value
  # source = "${local.folder_path}/${each.value}"
  content_type = "text/html"
}


# resource "aws_s3_bucket_policy" "allow_access_from_another_account" {
#   bucket = aws_s3_bucket.parmi-s3.id
#   policy = data.aws_iam_policy_document.allow_access_from_another_account.json
# }

# data "aws_iam_policy_document" "allow_access_from_another_account" {
#   statement {
#     effect = "Allow"

#     actions = [
#       "s3:*",
#     ]

#     resources = [
#       "${aws_s3_bucket.parmi-s3.arn}/*"
#     ]

#     principals {
#       type        = "AWS"
#       identifiers = ["*"]
#     }
#   }
# }

resource "aws_s3_bucket_website_configuration" "website" {
  bucket = aws_s3_bucket.parmi-s3.id
  index_document {
    suffix = "index.html"
  }
  error_document {
    key = "error.html"
  }
}

output "website" {
  value = aws_s3_bucket_website_configuration.website.website_endpoint
}