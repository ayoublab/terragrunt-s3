locals {
  region = "eu-central-1"
}

remote_state {
  backend = "s3"
  config = {
    bucket         = "morocco-terragtunt-tf-state-bucket"
    key            = "dev/s3/terraform.tfstate"
    region         = local.region
    encrypt        = true
    dynamodb_table = "terraform-locks"
  }
}
