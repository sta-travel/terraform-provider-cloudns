package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/sta-travel/cloudns-go"
)

func resourceRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceRecordCreate,
		Read:   resourceRecordRead,
		Update: resourceRecordUpdate,
		Delete: resourceRecordDelete,

		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: "remote id of the record, will be automatically added in most cases",
			},
			"domain": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Domain name for the record",
			},
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Host name for the record",
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Record type, can be one of: A, AAAA, MX, CNAME, TXT, NS, SRV, WR, RP, SSHFP, ALIAS, CAA, TLSA",
			},
			"ttl": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "TTL Value for the record, valid values are: 60,300,900,1800,3600,21600,43200,86400,172800,259200,604800,1209600,2592000",
			},
			"record": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The actual Value you want the record being set to. Must match the Type of the record.",
			},
		},
	}
}

func resourceRecordCreate(d *schemaResourceData, m interface{}) error {
	return nil
}

func resourceRecordRead(d *schemaResourceData, m interface{}) error {
	return nil
}

func resourceRecordUpdate(d *schemaResourceData, m interface{}) error {
	return nil
}

func resourceRecordDelete(d *schemaResourceData, m interface{}) error {
	return nil
}
