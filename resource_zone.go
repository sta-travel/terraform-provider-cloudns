package main

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/sta-travel/cloudns-go"
	"strings"
)

func validateZoneType(v interface{}, k string) (ws []string, es []error) {
	var possibleTypes = []string{"master", "slave", "parked", "geodns"}
	var possibleTypesStr = strings.Join(possibleTypes, ", ")
	var errs []error
	var warns []string

	value, ok := v.(string)

	if !ok {
		errs = append(errs, fmt.Errorf("expected Type to be a String!"))
		return warns, errs
	}

	for _, t := range possibleTypes {
		if value == t {
			return warns, errs
		}
	}

	errs = append(errs, fmt.Errorf("You have set an unknown zone type (%s), possible values are: %s", value, possibleTypesStr))
	return warns, errs
}

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
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Type of the zone. Can bei either of 'master/slave/parked/geodns'",
				ValidateFunc: validateZoneType,
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

func mkzone(d *schema.ResourceData) cloudns.Zone {
	z := cloudns.Zone{
		Domain: d.Get("domain").(string),
		Ztype:  d.Get("type").(string),
		Ns:     d.Get("nameservers").([]string),
	}
	return z
}

func resourceZoneCreate(d *schema.ResourceData, m interface{}) error {
	a, _ := m.(cloudns.Apiaccess)
	z := mkzone(d)
	_, err := z.Create(&a)
	return err
}

func resourceZoneRead(d *schema.ResourceData, m interface{}) error {
	a, _ := m.(cloudns.Apiaccess)
	z := mkzone(d)
	zr, err := z.Read(&a)
	d.Set("domain", zr.Domain)
	d.Set("type", zr.Ztype)
	d.Set("nameservers", zr.Ns)
	return err
}

func resourceZoneUpdate(d *schema.ResourceData, m interface{}) error {
	a, _ := m.(cloudns.Apiaccess)
	z := mkzone(d)
	_, err := z.Update(&a)
	return err
}

func resourceZoneDelete(d *schema.ResourceData, m interface{}) error {
	a, _ := m.(cloudns.Apiaccess)
	z := mkzone(d)
	_, err := z.Destroy(&a)
	return err
}
