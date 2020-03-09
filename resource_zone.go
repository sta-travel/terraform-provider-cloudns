package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/sta-travel/cloudns-go"
)

func resourceZone() *schema.Resource {
	return &schema.Resource{
		Create: resourceZoneCreate,
		Read:   resourceZoneRead,
		Update: resourceZoneUpdate,
		Delete: resourceZoneDelete,

		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The Domain name for the zone, also the unique ID for zone. The format is 'domain.tld' without a trailing dot.",
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of the zone. Can bei either of 'master/slave/parked/geodns'",
			},
			"nameservers": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of Name servers to use, can be omitted when the zone is of type 'master'",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceZoneCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceZoneRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceZoneUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceZoneDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
