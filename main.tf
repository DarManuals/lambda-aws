terraform {
  required_version = ">= 0.12"
  backend "s3" {}
}

provider "aws" {
  profile = var.profile
  region  = var.region
}

resource "null_resource" "sh0" {
  triggers = {
    created_at = timestamp()
  }
  provisioner "local-exec" {
    working_dir = "tmp"
    command     = "rm -rvf code; git clone git@github.com:${var.github_repo}.git code"
  }
}

resource "null_resource" "sh1" {
  triggers = {
    repo_changed = null_resource.sh0.id
    created_at   = timestamp()
  }
  provisioner "local-exec" {
    command = "make"
  }
}

module "lambda" {
  source     = "./lambda"
  name       = "test1"
  stage_name = "dev"
  file_name  = var.file_name
  account_id = data.aws_caller_identity.current.account_id
  region     = var.region
}

resource "null_resource" "sh2" {
  triggers = {
    last_id = null_resource.sh1.id
    lambda  = module.lambda.hash
  }
  provisioner "local-exec" {
    command = "rm ${var.file_name}; rm -rvf tmp/*"
  }
}


### Data
data "aws_caller_identity" "current" {}

### Variables
variable "profile" {}
variable "region" {}
variable "file_name" {
  default = "lambda.zip"
}
variable "github_repo" {
  default     = "DarManuals/lambda-aws"
  description = "github repo like: <user>/<repo_name>"
}

### Out
output "lambda_url" {
  value = module.lambda.lambda_url
}