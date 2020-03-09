package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"cloudns_zone":   resourceZone(),
			"cloudns_record": resourceRecord(),
		},
		Schema: map[string]*schema.Schema{
			"auth_password": {
				Type:     schema.TypeString,
				Required: true,
			},
			"authid": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"subauthid": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
