package aws

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

// See http://docs.aws.amazon.com/redshift/latest/mgmt/db-auditing.html#db-auditing-enable-logging
var redshiftAccountIdPerRegionMap = map[string]string{
	"us-east-1":      "193672423079",
	"us-west-1":      "262260360010",
	"us-west-2":      "902366379725",
	"ap-south-1":     "865932855811",
	"ap-northeast-2": "760740231472",
	"ap-southeast-1": "361669875840",
	"ap-southeast-2": "762762565011",
	"ap-northeast-1": "404641285394",
	"eu-central-1":   "053454850223",
	"eu-west-1":      "210876761215",
}

func dataSourceAwsRedshiftAccountId() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAwsRedshiftAccountIdRead,

		Schema: map[string]*schema.Schema{
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceAwsRedshiftAccountIdRead(d *schema.ResourceData, meta interface{}) error {
	region := meta.(*AWSClient).region
	if v, ok := d.GetOk("region"); ok {
		region = v.(string)
	}

	if accid, ok := redshiftAccountIdPerRegionMap[region]; ok {
		d.SetId(accid)
		return nil
	}

	return fmt.Errorf("Unknown region (%q)", region)
}
