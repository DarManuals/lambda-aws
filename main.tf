terraform {
  required_version = ">= 0.12"
  backend "s3" {}
}

provider "aws" {
  profile = var.profile
  region  = var.region
}

resource "null_resource" "sh1" {
  triggers = {
    created_at = timestamp()
  }
  provisioner "local-exec" {
    command = "make"
  }
}

module "lambda" {
  source    = "./lambda"
  name      = "test1"
  file_name = var.file_name
}

resource "null_resource" "sh2" {
  triggers = {
    hash = module.lambda.hash
    //    last_id = null_resource.sh1.id
  }
  provisioner "local-exec" {
    command = "rm ${var.file_name}"
  }
}


### Variables
variable "profile" {}
variable "region" {}
variable "file_name" {
  default = "lambda.zip"
}