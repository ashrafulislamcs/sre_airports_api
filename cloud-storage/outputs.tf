# Output the name of the S3 bucket
output "bucket_name" {
  description = "The name of the S3 bucket"
  value       = aws_s3_bucket.airport_images.bucket
}

# Output the ARN of the S3 bucket
output "bucket_arn" {
  description = "The ARN of the S3 bucket"
  value       = aws_s3_bucket.airport_images.arn
}

# Output the URL of the S3 bucket
output "bucket_url" {
  description = "The URL for accessing the S3 bucket"
  value       = aws_s3_bucket.airport_images.website_endpoint
}
