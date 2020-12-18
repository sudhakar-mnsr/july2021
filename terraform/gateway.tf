resource "aws_api_gateway_rest_api" "test" {
  name        = "EC2Example"
  description = "Terraform EC2 REST API Example"

  endpoint_configuration {
    types = ["REGIONAL"]
  }
}

// Method request configuration
resource "aws_api_gateway_method" "test" {
   rest_api_id   = aws_api_gateway_rest_api.test.id
   resource_id   = aws_api_gateway_rest_api.test.root_resource_id
   http_method   = "GET"

   authorization = "NONE"
 }
