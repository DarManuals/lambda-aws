# API

resource "aws_api_gateway_rest_api" "lambda_api" {
  name = "Lambda API for [${var.name}]"
}

resource "aws_api_gateway_resource" "lambda_api_res" {
  rest_api_id = aws_api_gateway_rest_api.lambda_api.id
  parent_id   = aws_api_gateway_rest_api.lambda_api.root_resource_id
  path_part   = var.name
}

resource "aws_api_gateway_method" "lambda_api_method_post" {
  rest_api_id   = aws_api_gateway_rest_api.lambda_api.id
  resource_id   = aws_api_gateway_resource.lambda_api_res.id
  http_method   = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "integration" {
  rest_api_id             = aws_api_gateway_rest_api.lambda_api.id
  resource_id             = aws_api_gateway_resource.lambda_api_res.id
  http_method             = aws_api_gateway_method.lambda_api_method_post.http_method
  uri                     = aws_lambda_function.lfn.invoke_arn
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
}

resource "aws_api_gateway_stage" "dev" {
  stage_name    = var.stage_name
  rest_api_id   = aws_api_gateway_rest_api.lambda_api.id
  deployment_id = aws_api_gateway_deployment.dev.id
}

resource "aws_api_gateway_deployment" "dev" {
  stage_name  = ""
  rest_api_id = aws_api_gateway_rest_api.lambda_api.id
}