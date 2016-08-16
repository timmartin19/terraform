---
layout: "aws"
page_title: "AWS: aws_redshift_account_id"
sidebar_current: "docs-aws-datasource-redshift-account-id"
description: |-
  Get AWS Redshift Account ID
---

# aws\_redshift\_account\_id

Use this data source to get the Account ID of the [AWS Redshift Account](http://docs.aws.amazon.com/redshift/latest/mgmt/db-auditing.html#db-auditing-enable-logging)
in a given region for the purpose of negotiating with other AWS services.

## Example Usage

```
data "aws_redshift_account_id" "main" { }

resource "aws_s3_bucket" "bucket" {
	bucket = "tf-redshift-logging-test-bucket"
	force_destroy = true
	policy = <<EOF
{
	"Version": "2008-10-17",
	"Statement": [
		{
        			"Sid": "Put bucket policy needed for audit logging",
        			"Effect": "Allow",
        			"Principal": {
        				"AWS": "arn:aws:iam:${data.aws_redshift_account_id.main.id}:user/logs"
        			},
        			"Action": "s3:PutObject",
        			"Resource": "arn:aws:s3:::tf-redshift-logging-test-bucket/*"
        		},
        		{
        			"Sid": "Get bucket policy needed for audit logging ",
        			"Effect": "Allow",
        			"Principal": {
        				"AWS": "arn:aws:iam:${data.aws_redshift_account_id.main.id}:user/logs"
        			},
        			"Action": "s3:GetBucketAcl",
        			"Resource": "arn:aws:s3:::tf-redshift-logging-test-bucket"
        		}
	]
}
EOF
}
```

## Argument Reference

* `region` - (Optional) Region of a given AWS Redshift Account


## Attributes Reference

* `id` - Account ID
