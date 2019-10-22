package main

import (
	"fmt"
	"strconv"
	"strings"
	"terraform-provider-subreg-cz/subreg"

	"github.com/hashicorp/terraform/helper/schema"
)

const ncDefaultTTL int = 1800
const ncDefaultMXPref int = 10

func resourceSubregRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceSubgegCzRecordCreate,
		Update: resourceSubgegCzRecordUpdate,
		Read:   resourceSubgegCzRecordRead,
		Importer: &schema.ResourceImporter{
			State: resourceSubgegCzRecordImport,
		},
		Delete: resourceSubgegCzRecordDelete,
		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Default:  "",
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"content": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"prio": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  900,
			},
		},
	}
}

func resourceSubgegCzRecordCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*subreg.SubregCz)
	domain := d.Get("domain").(string)

	newRecord := subreg.AddDNSRecordRecord{

		Name:    d.Get("name").(string),
		Type_:   d.Get("type").(string),
		Content: d.Get("content").(string),
		Prio:    int32(d.Get("prio").(int)),
		Ttl:     int32(d.Get("ttl").(int)),
	}

	// api returns only ok not id :/
	if _, err := client.AddDNSRecord(&subreg.AddDNSRecord{
		Domain: domain,
		Record: &newRecord,
	}); err != nil {
		return err
	}

	result, err := client.GetDNSZone(&subreg.GetDNSZone{Domain: domain})
	if err != nil {
		return err
	}

	// find new added record and get id
	for _, record := range result.Response.Data.Records {
		if record.Type_ == newRecord.Type_ &&
			record.Name == newRecord.Name &&
			record.Content == newRecord.Content {

			d.SetId(strconv.Itoa(int(record.Id)))
			return resourceSubgegCzRecordRead(d, meta)
		}

	}
	return fmt.Errorf("Cannot find newly added record in api")

}

func resourceSubgegCzRecordUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*subreg.SubregCz)
	domain := d.Get("domain").(string)

	hashID, err := strconv.ParseInt(d.Id(), 10, 32)
	if err != nil {
		return fmt.Errorf("Failed to parse id: %s", err)
	}
	// name nejde menit
	record := subreg.ModifyDNSRecordRecord{
		Id:      int32(hashID),
		Type_:   d.Get("type").(string),
		Content: d.Get("content").(string),
		Prio:    int32(d.Get("prio").(int)),
		Ttl:     int32(d.Get("ttl").(int)),
	}

	_, err = client.ModifyDNSRecord(&subreg.ModifyDNSRecord{Domain: domain, Record: &record})
	if err != nil {
		return err
	}

	return resourceSubgegCzRecordRead(d, meta)
}
func resourceSubgegCzRecordRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*subreg.SubregCz)
	domain := d.Get("domain").(string)

	hashID, err := strconv.ParseInt(d.Id(), 10, 32)
	if err != nil {
		return fmt.Errorf("Failed to parse id: %s", err)
	}

	r2 := subreg.GetDNSZone{Domain: domain}

	result2, err := client.GetDNSZone(&r2)
	if err != nil {
		return err
	}
	records := result2.Response.Data.Records
	for _, record := range records {
		if record.Id == int32(hashID) {

			d.Set("name", record.Name)
			d.Set("type", record.Type_)
			d.Set("content", record.Content)
			d.Set("prio", record.Prio)
			d.Set("ttl", record.Ttl)

			return nil
		}

	}
	return fmt.Errorf("Cannot find record with id: %d", hashID)

}

func resourceSubgegCzRecordDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*subreg.SubregCz)
	domain := d.Get("domain").(string)

	hashID, err := strconv.ParseInt(d.Id(), 10, 32)
	if err != nil {
		return fmt.Errorf("Failed to parse id: %s", err)
	}
	_, err = client.DeleteDNSRecord(&subreg.DeleteDNSRecord{
		Domain: domain,
		Record: &subreg.DeleteDNSRecordRecord{Id: int32(hashID)}})
	if err != nil {
		return fmt.Errorf("Failed to delete subreg record record: %s", err)
	}

	return nil
}

func resourceSubgegCzRecordImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), "_")

	if len(parts) != 2 {
		return nil, fmt.Errorf("Error Importing subregcz record. Please make sure the record ID is in the form DOMAIN_RECORDID (i.e. example.com_1234")
	}

	d.SetId(parts[1])
	d.Set("domain", parts[0])

	if err := resourceSubgegCzRecordRead(d, meta); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
