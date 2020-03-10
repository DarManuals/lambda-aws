resource "aws_lambda_function" "lfn" {
  function_name = var.name
  filename      = var.file_name
  //  source_code_hash = filebase64sha256(var.file_name)

  handler = "main"
  role    = aws_iam_role.iam_for_lambda.arn
  runtime = "go1.x"

  tags = {
    created_at_tag = timestamp()
  }
}

resource "aws_iam_role" "iam_for_lambda" {
  name               = "iam_for_lambda"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

variable "name" {}
variable "file_name" {}

output "name" {
  value = aws_lambda_function.lfn.function_name
}
output "hash" {
  value = aws_lambda_function.lfn.source_code_hash
}