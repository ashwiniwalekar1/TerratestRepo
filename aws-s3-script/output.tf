output "bucket_name" {
  value = aws_s3_bucket.parmi-s3.bucket
}
output "region" {
  value = aws_s3_bucket.parmi-s3.region
}