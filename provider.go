package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/sta-travel/cloudns-go"
)

func configureCloudns(d *schema.ResourceData) (interface{}, error) {
	a := cloudns.Apiaccess{
		Authid:       d.Get("authid").(int),
		Subauthid:    d.Get("subauthid").(int),
		Authpassword: d.Get("auth_password").(string),
	}

	_, err := a.Listzones()

	return a, err
}

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"cloudns_zone":   resourceZone(),
			"cloudns_record": resourceRecord(),
		},
		Schema: map[string]*schema.Schema{
			"auth_password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "auth password issued by ClouDNS",
			},
			"authid": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "authid issued by ClouDNS",
			},
			"subauthid": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "subauthid issued by ClouDNS",
			},
		},
		ConfigureFunc: configureCloudns,
	}
}
