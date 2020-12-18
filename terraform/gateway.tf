resource "aws_api_gateway_rest_api" "test" {
  name        = "EC2Example"
  description = "Terraform EC2 REST API Example"

  endpoint_configuration {
    types = ["REGIONAL"]
  }
}

