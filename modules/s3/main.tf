terraform {
  backend "s3" {}
  
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.region
}


resource "aws_s3_bucket" "this" {
  bucket = var.bucket_name
  tags   = var.tags
}


resource "aws_s3_bucket_server_side_encryption_configuration" "this" {
  bucket = aws_s3_bucket.this.id

  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}


resource "aws_s3_bucket_versioning" "this" {
  bucket = aws_s3_bucket.this.id

  versioning_configuration {
    status = "Enabled"
  }
}


resource "aws_s3_bucket_public_access_block" "this" {
  bucket = aws_s3_bucket.this.id

  block_public_acls       = true
  ignore_public_acls      = true
  block_public_policy     = true
  restrict_public_buckets = true
}


variable "bucket_name" {
  description = "The name of the S3 bucket"
  type        = string
}

variable "region" {
  description = "The AWS region to deploy the S3 bucket"
  type        = string
  default     = "us-east-1"
}

variable "tags" {
  description = "Tags for the S3 bucket"
  type        = map(string)
  default     = {}
}
