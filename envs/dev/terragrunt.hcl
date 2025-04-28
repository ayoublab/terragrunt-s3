include "root" {
  path   = find_in_parent_folders("root.hcl")
  expose = true
}


locals {
  region = include.root.locals.region
}

terraform {
  source = "../../modules/s3"
}

inputs = {
  bucket_name = "terragrunt-s3-dev-bucket-test-morocco"
  region      = local.region   
  tags = {
    environment = "dev"
    owner       = "myself"
  }
}
