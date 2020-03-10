# Lambda

resource "aws_lambda_function" "lfn" {
  function_name = var.name
  filename      = var.file_name
  role          = aws_iam_role.iam_for_lambda.arn
  handler       = "main"
  runtime       = "go1.x"
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

resource "aws_lambda_permission" "apigw_lambda" {
  function_name = aws_lambda_function.lfn.function_name
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  principal     = "apigateway.amazonaws.com"

  # More: http://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-control-access-using-iam-policies-to-invoke-api.html
  source_arn = "arn:aws:execute-api:${var.region}:${var.account_id}:${aws_api_gateway_rest_api.lambda_api.id}/*/${aws_api_gateway_method.lambda_api_method_post.http_method}${aws_api_gateway_resource.lambda_api_res.path}"
}