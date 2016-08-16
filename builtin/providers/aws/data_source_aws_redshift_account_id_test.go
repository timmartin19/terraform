package aws

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccAWSRedshiftAccountId_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckAwsRedshiftAccountIdConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.aws_redshift_account_id.main", "id", "902366379725"),
				),
			},
			resource.TestStep{
				Config: testAccCheckAwsRedshiftAccountIdExplicitRegionConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.aws_redshift_account_id.regional", "id", "210876761215"),
				),
			},
		},
	})
}

const testAccCheckAwsRedshiftAccountIdConfig = `
data "aws_redshift_account_id" "main" { }
`

const testAccCheckAwsRedshiftAccountIdExplicitRegionConfig = `
data "aws_redshift_account_id" "regional" {
	region = "eu-west-1"
}
`
