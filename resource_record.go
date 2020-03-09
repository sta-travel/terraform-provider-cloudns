package main

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/sta-travel/cloudns-go"
	"strings"
)

func validateType(v interface{}, k string) (ws []string, es []error) {
	var possibleTypes = []string{"A", "AAAA", "MX", "CNAME", "TXT", "NS", "SRV", "WR", "RP", "SSHFP", "ALIAS", "CAA", "TLSA"}
	var possibleTypesStr = strings.Join(possibleTypes, ", ")
	var errs []error
	var warns []string
	value, ok := v.(string)
	if !ok {
		errs = append(errs, fmt.Errorf("expected Type to be a String"))
		return warns, errs
	}

	for _, t := range possibleTypes {
		if value == t {
			return warns, errs
		}
	}

	errs = append(errs, fmt.Errorf("You have set an unknown record type (%s), possible values are: %s", value, possibleTypesStr))
	return warns, errs

}

func validateTtl(v interface{}, k string) (ws []string, es []error) {
	var possibleTtls = []int{60, 300, 900, 1800, 3600, 21600, 43200, 86400, 172800, 259200, 604800, 1209600, 2592000}
	var possibleTtlsStr = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(possibleTtls)), ", "), "[]")

	var errs []error
	var warns []string
	value, ok := v.(int)
	if !ok {
		errs = append(errs, fmt.Errorf("expected TTL to be Int"))
		return warns, errs
	}

	for _, t := range possibleTtls {
		if value == t {
			return warns, errs
		}
	}

	errs = append(errs, fmt.Errorf("You have set TTL to a value not accepted by the API (%s), possible values are: %s", value, possibleTtlsStr))
	return warns, errs

}

func resourceRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceRecordCreate,
		Read:   resourceRecordRead,
		Update: resourceRecordUpdate,
		Delete: resourceRecordDelete,

		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
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
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Record type, can be one of: A, AAAA, MX, CNAME, TXT, NS, SRV, WR, RP, SSHFP, ALIAS, CAA, TLSA",
				ValidateFunc: validateType,
			},
			"ttl": &schema.Schema{
				Type:         schema.TypeInt,
				Required:     true,
				Description:  "TTL Value for the record, valid values are: 60,300,900,1800,3600,21600,43200,86400,172800,259200,604800,1209600,2592000",
				ValidateFunc: validateTtl,
			},
			"record": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The actual Value you want the record being set to. Must match the Type of the record.",
			},
		},
	}
}

func mkrec(d *schema.ResourceData) cloudns.Record {
	r := cloudns.Record{
		Domain: d.Get("domain").(string),
		ID:     d.Get("id").(string),
		Host:   d.Get("host").(string),
		Rtype:  d.Get("type").(string),
		Record: d.Get("record").(string),
		TTL:    d.Get("ttl").(int),
	}
	return r
}

func resourceRecordCreate(d *schema.ResourceData, m interface{}) error {
	r := mkrec(d)
	a, _ := m.(cloudns.Apiaccess)
	rc, err := r.Create(&a)
	d.Set("id", rc.ID)
	return err
}

func resourceRecordRead(d *schema.ResourceData, m interface{}) error {
	r := mkrec(d)
	a, _ := m.(cloudns.Apiaccess)
	rr, err := r.Read(&a)
	d.Set("id", rr.ID)
	return err
}

func resourceRecordUpdate(d *schema.ResourceData, m interface{}) error {
	r := mkrec(d)
	a, _ := m.(cloudns.Apiaccess)
	ru, err := r.Update(&a)
	d.Set("id", ru.ID)
	return err
}

func resourceRecordDelete(d *schema.ResourceData, m interface{}) error {
	r := mkrec(d)
	a, _ := m.(cloudns.Apiaccess)
	_, err := r.Destroy(&a)
	return err
}
