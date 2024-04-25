package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/ackers-bud/terraform-provider-zip2b64/client"
)

func resourcezip2b64() *schema.Resource {
	return &schema.Resource{
		Create: Create,
		Update: Update,
		Read:   ReadUrl,
		Delete: Delete,

		Schema: map[string]*schema.Schema{
			"base64file": {
				Type:     schema.TypeString,
				Required: true,
			},

			"filename": {
				Description: "the filename",
				Type:        schema.TypeString,
				Required:    true,
			},

			"id": {
				Description: "The ID of this resource.",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"filecontents_base64": {
				Description: "The Returned body base64 encoded",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func Create(d *schema.ResourceData, meta interface{}) error {

	base64file := d.Get("base64file").(string)
	filenameToExtract := d.Get("filename").(string)

	filecontentsBase64, err := client.ZipExtract(base64file, filenameToExtract)
	if err != nil {
		return fmt.Errorf("error extracting file '%s' from base64 string error: '%v'", filenameToExtract, err)
	}

	err = d.Set("filecontents_base64", filecontentsBase64)
	if err != nil {
		return fmt.Errorf("error setting filecontents_base64 '%v'", err)
	}

	d.SetId(filenameToExtract)
	return nil
}

func Update(d *schema.ResourceData, meta interface{}) error {

	base64file := d.Get("base64file").(string)
	filenameToExtract := d.Get("filename").(string)

	filecontentsBase64, err := client.ZipExtract(base64file, filenameToExtract)
	if err != nil {
		return fmt.Errorf("error extracting file '%s' from base64 string error: '%v'", filenameToExtract, err)
	}

	err = d.Set("filecontents_base64", filecontentsBase64)
	if err != nil {
		return fmt.Errorf("error setting filecontents_base64 '%v'", err)
	}

	d.SetId(filenameToExtract)
	return nil
}

func ReadUrl(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func Delete(d *schema.ResourceData, meta interface{}) error {

	d.SetId("")
	return nil
}
