package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return Provider()
		},
	})
}

func Provider() *schema.Provider {
	rand.Seed(time.Now().Unix())

	return &schema.Provider{
		Schema: map[string]*schema.Schema{},

		ResourcesMap: map[string]*schema.Resource{
			"baseball_bat": resourceBat(),
		},
	}
}

func CreateBaseballClient(venue string) (string, error) {
	log.Printf("[DEBUG] called baseball client create")
	return "", nil
}

func resourceBat() *schema.Resource {
	return &schema.Resource{
		Create: resourceBatCreate,
		Read:   resourceBatRead,
		Update: resourceBatUpdate,
		Delete: resourceBatDelete,

		Schema: map[string]*schema.Schema{
			"targets": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"results": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceBatCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] called baseball_bat create")
	d.SetId(fmt.Sprintf("%d", rand.Int()))

	targets := d.Get("targets").([]interface{})
	results := make([]string, len(targets))

	for i, target := range targets {
		results[i] = target.(string)
	}

  d.Set("results", results)

	return nil
}

func resourceBatRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] called read")
	return nil
}

func resourceBatUpdate(d *schema.ResourceData, meta interface{}) error {
	targets := d.Get("targets").([]interface{})
	results := make([]string, len(targets))

	for i, target := range targets {
		results[i] = target.(string)
	}

  d.Set("results", results)

	return nil
}

func resourceBatDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] called delete")
	targets := d.Get("targets").([]interface{})
	for target := range targets {
		log.Printf(fmt.Sprintf("[INFO] Beating %s with a baseball bat on the way out"), target)
	}
	return nil
}
