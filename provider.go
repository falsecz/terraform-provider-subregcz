package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SUBREG_USERNAME", nil),
				Description: "A registered username for subreg.cz",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SUBREG_PASSWORD", nil),
				Description: "A registered password for subreg.cz",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"subreg_record": resourceSubregRecord(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	config := Config{
		username: d.Get("username").(string),
		password: d.Get("password").(string),
	}

	return config.Client()
}
