package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/knockoutez/terraform-provider-crudcrud/client"
)

func resourceUnicorn() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			// "_id": {
			// 	Type:        schema.TypeString,
			// 	Required:     true,
			// 	Description: "The id of the unicorn resource",
			// 	ForceNew:     true,
			// },
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the unicorn resource",
			},
			"age": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The age of the unicorn resource",
			},
			"colour": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The colour of the unicorn resource",
			},
		},
		Create: resourceCreateItem,
		Read:   resourceReadItem,
		Update: resourceUpdateItem,
		Delete: resourceDeleteItem,
		Exists: resourceExistsItem,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateItem(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	item := client.Unicorn{
		Name:   d.Get("name").(string),
		Age:    d.Get("age").(int),
		Colour: d.Get("colour").(string),
	}

	unicorn,err := apiClient.NewItem(&item)

	if err != nil {
		return err
	}
	fmt.Println(unicorn.Id,"item id")
	d.SetId(unicorn.Id)
	return nil
}

func resourceReadItem(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	item, err := apiClient.GetItem(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding Item with ID %s", itemId)
		}
	}

	d.SetId(item.Id)
	d.Set("name", item.Name)
	d.Set("age", item.Age)
	d.Set("colour", item.Colour)
	return nil
}

func resourceUpdateItem(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	item := client.Unicorn{
		Name:        d.Get("name").(string),
		Age:         d.Get("age").(int),
		Colour:      d.Get("colour").(string),
	}

	err := apiClient.UpdateItem(d.Id(),&item)
	if err != nil {
		return err
	}
	return nil
}

func resourceDeleteItem(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteItem(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsItem(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.GetItem(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
