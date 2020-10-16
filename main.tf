
provider "aws" {
  region = var.region
}

resource "aws_instance" "web_server" {
  ami = "ami-07efac79022b86107"
  instance_type = "t2.micro"
  vpc_security_group_ids = [aws_security_group.secgroup1.id]

  tags = {
    Name = var.instance_name
  }

  user_data = <<-EOF
              #!/bin/bash
              echo "website is online" > index.html
              nohup busybox httpd -f -p "${var.server_port}" &
              EOF
}

  resource "aws_security_group" "secgroup1" {
    name = "terraform-allow-traffic"

    ingress {
      from_port = var.server_port
      to_port = var.server_port
      protocol = "tcp"
      cidr_blocks = ["0.0.0.0/0"]
    }
  }

  output "public_ip" {
    value = aws_instance.web_server.public_ip
  }

  variable "server_port" {
    description = "The port the server will use for HTTP requests"
    type        = number
    default     = 8080
  }

output "instance_id" {
  value = aws_instance.web_server.id
}