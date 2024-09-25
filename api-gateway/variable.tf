# Variables for Cloud Storage (S3)
variable "bucket_name" {
  description = "Name of the S3 bucket to store airport images"
  type        = string
  default     = "airport-images-bucket"
}

variable "bucket_region" {
  description = "AWS Region for the S3 bucket"
  type        = string
  default     = "ap-south-1"
}
