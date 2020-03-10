output "hash" {
  value = aws_lambda_function.lfn.source_code_hash
}

output "lambda_url" {
  value = "${aws_api_gateway_stage.dev.invoke_url}/${aws_api_gateway_resource.lambda_api_res.path_part}"
}