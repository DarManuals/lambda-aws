remote_state {
  backend = "s3"
  config = {
    bucket = "lambda-terra"
    key = "${path_relative_to_include()}/terraform.tfstate"
    region = "eu-central-1"
    encrypt = true
    dynamodb_table = "my-lock-table2"
    profile = "my"
  }
}

inputs = {
  region = "eu-central-1"
  profile = "my"
}