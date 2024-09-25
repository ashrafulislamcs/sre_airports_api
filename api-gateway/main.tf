provider "aws" {
  region = "ap-south-1"
}

resource "aws_s3_bucket" "airport_images" {
  bucket = "airport-images-bucket"
  
  versioning {
    enabled = true
  }

  tags = {
    Name        = "Airport Images"
    Environment = "Dev"
  }
}

output "bucket_test" {
  value = aws_s3_bucket.airport_images.bucket
}
