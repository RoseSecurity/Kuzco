provider "aws" {
  region  = "us-west-2"
  profile = "default"
}

resource "aws_instance" "example" {
  ami           = "ami-0c55b159cbfafe1f0"
  instance_type = "t2.micro"

  tags {
    Name = "ExampleInstance"
  }

  invalid_attribute = "This is not a valid attribute for aws_instance"
}

