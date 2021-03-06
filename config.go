package main

import (
	"github.com/falsecz/terraform-provider-subregcz/subreg"
)

type Config struct {
	username string
	password string
}

// Client() returns a new client for accessing Namecheap.
func (c *Config) Client() (*subreg.SubregCz, error) {

	auth := subreg.BasicAuth{Login: c.username, Password: c.password}

	client := subreg.NewSubregCz("", true, &auth)
	return client, nil

}
