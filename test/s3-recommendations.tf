resource "aws_s3_bucket" "cloudtrail_logs" {
  bucket              = "my-cloudtrail-logs-bucket"
  object_lock_enabled = true

  tags = {
    Name        = "My CloudTrail Bucket"
    Environment = "Dev"
    Region      = "us-west-2"
  }
}
