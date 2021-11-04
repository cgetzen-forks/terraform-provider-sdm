// This file was generated by protogen. DO NOT EDIT.

package sdm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

func resourceNode() *schema.Resource {
	return &schema.Resource{
		Create: wrapCrudOperation(resourceNodeCreate),
		Read:   wrapCrudOperation(resourceNodeRead),
		Update: wrapCrudOperation(resourceNodeUpdate),
		Delete: wrapCrudOperation(resourceNodeDelete),
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"gateway": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Gateway represents a StrongDM CLI installation running in gateway mode.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bind_address": {
							Type:     schema.TypeString,
							Optional: true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								return new == ""
							},
							ForceNew:    true,
							Description: "The hostname/port tuple which the gateway daemon will bind to. If not provided on create, set to \"0.0.0.0:<listen_address_port>\".",
						},
						"gateway_filter": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "GatewayFilter can be used to restrict the peering between relays and gateways.",
						},
						"listen_address": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The public hostname/port tuple at which the gateway will be accessible to clients.",
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								return new == ""
							},
							Description: "Unique human-readable name of the Gateway. Node names must include only letters, numbers, and hyphens (no spaces, underscores, or other special characters). Generated if not provided on create.",
						},
						"tags": {
							Type: schema.TypeMap,

							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Optional:    true,
							Description: "Tags is a map of key, value pairs.",
						},
						"token": {
							Type:      schema.TypeString,
							Computed:  true,
							Sensitive: true,
						},
					},
				},
			},
			"relay": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Relay represents a StrongDM CLI installation running in relay mode.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gateway_filter": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "GatewayFilter can be used to restrict the peering between relays and gateways.",
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								return new == ""
							},
							Description: "Unique human-readable name of the Relay. Node names must include only letters, numbers, and hyphens (no spaces, underscores, or other special characters). Generated if not provided on create.",
						},
						"tags": {
							Type: schema.TypeMap,

							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Optional:    true,
							Description: "Tags is a map of key, value pairs.",
						},
						"token": {
							Type:      schema.TypeString,
							Computed:  true,
							Sensitive: true,
						},
					},
				},
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(60 * time.Second),
		},
	}
}
func convertNodeFromResourceData(d *schema.ResourceData) sdm.Node {
	if list := d.Get("gateway").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.Gateway{}
		}
		out := &sdm.Gateway{
			ID:            d.Id(),
			BindAddress:   convertStringFromMap(raw, "bind_address"),
			GatewayFilter: convertStringFromMap(raw, "gateway_filter"),
			ListenAddress: convertStringFromMap(raw, "listen_address"),
			Name:          convertStringFromMap(raw, "name"),
			Tags:          convertTagsFromMap(raw, "tags"),
		}
		return out
	}
	if list := d.Get("relay").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.Relay{}
		}
		out := &sdm.Relay{
			ID:            d.Id(),
			GatewayFilter: convertStringFromMap(raw, "gateway_filter"),
			Name:          convertStringFromMap(raw, "name"),
			Tags:          convertTagsFromMap(raw, "tags"),
		}
		return out
	}
	return nil
}

func resourceNodeCreate(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutCreate))
	defer cancel()
	localVersion := convertNodeFromResourceData(d)

	resp, err := cc.Nodes().Create(ctx, localVersion)
	if err != nil {
		return fmt.Errorf("cannot create Node: %w", err)
	}
	d.SetId(resp.Node.GetID())
	switch v := resp.Node.(type) {
	case *sdm.Gateway:
		localV, _ := localVersion.(*sdm.Gateway)
		_ = localV
		d.Set("gateway", []map[string]interface{}{
			{
				"bind_address":   (v.BindAddress),
				"gateway_filter": (v.GatewayFilter),
				"listen_address": (v.ListenAddress),
				"name":           (v.Name),
				"tags":           convertTagsToMap(v.Tags),
				"token":          resp.Token,
			},
		})
	case *sdm.Relay:
		localV, _ := localVersion.(*sdm.Relay)
		_ = localV
		d.Set("relay", []map[string]interface{}{
			{
				"gateway_filter": (v.GatewayFilter),
				"name":           (v.Name),
				"tags":           convertTagsToMap(v.Tags),
				"token":          resp.Token,
			},
		})
	}
	return nil
}

func resourceNodeRead(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	localVersion := convertNodeFromResourceData(d)
	_ = localVersion

	resp, err := cc.Nodes().Get(ctx, d.Id())
	var errNotFound *sdm.NotFoundError
	if err != nil && errors.As(err, &errNotFound) {
		d.SetId("")
		return nil
	} else if err != nil {
		return fmt.Errorf("cannot read Node %s: %w", d.Id(), err)
	}
	switch v := resp.Node.(type) {
	case *sdm.Gateway:
		localV, ok := localVersion.(*sdm.Gateway)
		if !ok {
			localV = &sdm.Gateway{}
		}
		_ = localV
		d.Set("gateway", []map[string]interface{}{
			{
				"bind_address":   (v.BindAddress),
				"gateway_filter": (v.GatewayFilter),
				"listen_address": (v.ListenAddress),
				"name":           (v.Name),
				"tags":           convertTagsToMap(v.Tags),
				"token":          d.Get("gateway.0.token"),
			},
		})
	case *sdm.Relay:
		localV, ok := localVersion.(*sdm.Relay)
		if !ok {
			localV = &sdm.Relay{}
		}
		_ = localV
		d.Set("relay", []map[string]interface{}{
			{
				"gateway_filter": (v.GatewayFilter),
				"name":           (v.Name),
				"tags":           convertTagsToMap(v.Tags),
				"token":          d.Get("relay.0.token"),
			},
		})
	}
	return nil
}
func resourceNodeUpdate(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutUpdate))
	defer cancel()

	resp, err := cc.Nodes().Update(ctx, convertNodeFromResourceData(d))
	if err != nil {
		return fmt.Errorf("cannot update Node %s: %w", d.Id(), err)
	}
	d.SetId(resp.Node.GetID())
	return resourceNodeRead(d, cc)
}
func resourceNodeDelete(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutDelete))
	defer cancel()
	var errNotFound *sdm.NotFoundError
	_, err := cc.Nodes().Delete(ctx, d.Id())
	if err != nil && errors.As(err, &errNotFound) {
		return nil
	}
	return err
}
