variable "region" {
  description = "AWS region to deploy the instances"
  default = "us-east-2"
}

variable "instance_name" {
  default     = "testing-test"
}

variable "test_label" {
  default = "yes"
}