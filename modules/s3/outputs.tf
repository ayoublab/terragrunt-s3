output "bucket_name" {
  description = "S3 bucketname"
  value       = aws_s3_bucket.this.id
}

output "region" {
  description = "aws deployment region"
  value       = var.region
}
