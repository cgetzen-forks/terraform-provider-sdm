// This file was generated by protogen. DO NOT EDIT.

package sdm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

func resourceSecretStore() *schema.Resource {
	return &schema.Resource{
		CreateContext: wrapCrudOperation(resourceSecretStoreCreate),
		ReadContext:   wrapCrudOperation(resourceSecretStoreRead),
		UpdateContext: wrapCrudOperation(resourceSecretStoreUpdate),
		DeleteContext: wrapCrudOperation(resourceSecretStoreDelete),
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"aws": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the SecretStore.",
						},
						"region": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"tags": {
							Type:        schema.TypeMap,
							Elem:        tagsElemType,
							Optional:    true,
							Description: "Tags is a map of key, value pairs.",
						},
					},
				},
			},
			"azure_store": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the SecretStore.",
						},
						"tags": {
							Type:        schema.TypeMap,
							Elem:        tagsElemType,
							Optional:    true,
							Description: "Tags is a map of key, value pairs.",
						},
						"vault_uri": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
					},
				},
			},
			"vault_tls": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ca_cert_path": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"client_cert_path": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"client_key_path": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the SecretStore.",
						},
						"namespace": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_address": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"tags": {
							Type:        schema.TypeMap,
							Elem:        tagsElemType,
							Optional:    true,
							Description: "Tags is a map of key, value pairs.",
						},
					},
				},
			},
			"vault_token": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the SecretStore.",
						},
						"namespace": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_address": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"tags": {
							Type:        schema.TypeMap,
							Elem:        tagsElemType,
							Optional:    true,
							Description: "Tags is a map of key, value pairs.",
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
func convertSecretStoreFromResourceData(d *schema.ResourceData) sdm.SecretStore {
	if list := d.Get("aws").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.AWSStore{}
		}
		out := &sdm.AWSStore{
			ID:     d.Id(),
			Name:   convertStringFromMap(raw, "name"),
			Region: convertStringFromMap(raw, "region"),
			Tags:   convertTagsFromMap(raw, "tags"),
		}
		return out
	}
	if list := d.Get("azure_store").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.AzureStore{}
		}
		out := &sdm.AzureStore{
			ID:       d.Id(),
			Name:     convertStringFromMap(raw, "name"),
			Tags:     convertTagsFromMap(raw, "tags"),
			VaultUri: convertStringFromMap(raw, "vault_uri"),
		}
		return out
	}
	if list := d.Get("vault_tls").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.VaultTLSStore{}
		}
		out := &sdm.VaultTLSStore{
			ID:             d.Id(),
			CACertPath:     convertStringFromMap(raw, "ca_cert_path"),
			ClientCertPath: convertStringFromMap(raw, "client_cert_path"),
			ClientKeyPath:  convertStringFromMap(raw, "client_key_path"),
			Name:           convertStringFromMap(raw, "name"),
			Namespace:      convertStringFromMap(raw, "namespace"),
			ServerAddress:  convertStringFromMap(raw, "server_address"),
			Tags:           convertTagsFromMap(raw, "tags"),
		}
		return out
	}
	if list := d.Get("vault_token").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.VaultTokenStore{}
		}
		out := &sdm.VaultTokenStore{
			ID:            d.Id(),
			Name:          convertStringFromMap(raw, "name"),
			Namespace:     convertStringFromMap(raw, "namespace"),
			ServerAddress: convertStringFromMap(raw, "server_address"),
			Tags:          convertTagsFromMap(raw, "tags"),
		}
		return out
	}
	return nil
}

func resourceSecretStoreCreate(ctx context.Context, d *schema.ResourceData, cc *sdm.Client) error {
	localVersion := convertSecretStoreFromResourceData(d)

	resp, err := cc.SecretStores().Create(ctx, localVersion)
	if err != nil {
		return fmt.Errorf("cannot create SecretStore: %w", err)
	}
	d.SetId(resp.SecretStore.GetID())
	switch v := resp.SecretStore.(type) {
	case *sdm.AWSStore:
		localV, _ := localVersion.(*sdm.AWSStore)
		_ = localV
		d.Set("aws", []map[string]interface{}{
			{
				"name":   (v.Name),
				"region": (v.Region),
				"tags":   convertTagsToMap(v.Tags),
			},
		})
	case *sdm.AzureStore:
		localV, _ := localVersion.(*sdm.AzureStore)
		_ = localV
		d.Set("azure_store", []map[string]interface{}{
			{
				"name":      (v.Name),
				"tags":      convertTagsToMap(v.Tags),
				"vault_uri": (v.VaultUri),
			},
		})
	case *sdm.VaultTLSStore:
		localV, _ := localVersion.(*sdm.VaultTLSStore)
		_ = localV
		d.Set("vault_tls", []map[string]interface{}{
			{
				"ca_cert_path":     (v.CACertPath),
				"client_cert_path": (v.ClientCertPath),
				"client_key_path":  (v.ClientKeyPath),
				"name":             (v.Name),
				"namespace":        (v.Namespace),
				"server_address":   (v.ServerAddress),
				"tags":             convertTagsToMap(v.Tags),
			},
		})
	case *sdm.VaultTokenStore:
		localV, _ := localVersion.(*sdm.VaultTokenStore)
		_ = localV
		d.Set("vault_token", []map[string]interface{}{
			{
				"name":           (v.Name),
				"namespace":      (v.Namespace),
				"server_address": (v.ServerAddress),
				"tags":           convertTagsToMap(v.Tags),
			},
		})
	}
	return nil
}

func resourceSecretStoreRead(ctx context.Context, d *schema.ResourceData, cc *sdm.Client) error {
	localVersion := convertSecretStoreFromResourceData(d)
	_ = localVersion

	resp, err := cc.SecretStores().Get(ctx, d.Id())
	var errNotFound *sdm.NotFoundError
	if err != nil && errors.As(err, &errNotFound) {
		d.SetId("")
		return nil
	} else if err != nil {
		return fmt.Errorf("cannot read SecretStore %s: %w", d.Id(), err)
	}
	switch v := resp.SecretStore.(type) {
	case *sdm.AWSStore:
		localV, ok := localVersion.(*sdm.AWSStore)
		if !ok {
			localV = &sdm.AWSStore{}
		}
		_ = localV
		d.Set("aws", []map[string]interface{}{
			{
				"name":   (v.Name),
				"region": (v.Region),
				"tags":   convertTagsToMap(v.Tags),
			},
		})
	case *sdm.AzureStore:
		localV, ok := localVersion.(*sdm.AzureStore)
		if !ok {
			localV = &sdm.AzureStore{}
		}
		_ = localV
		d.Set("azure_store", []map[string]interface{}{
			{
				"name":      (v.Name),
				"tags":      convertTagsToMap(v.Tags),
				"vault_uri": (v.VaultUri),
			},
		})
	case *sdm.VaultTLSStore:
		localV, ok := localVersion.(*sdm.VaultTLSStore)
		if !ok {
			localV = &sdm.VaultTLSStore{}
		}
		_ = localV
		d.Set("vault_tls", []map[string]interface{}{
			{
				"ca_cert_path":     (v.CACertPath),
				"client_cert_path": (v.ClientCertPath),
				"client_key_path":  (v.ClientKeyPath),
				"name":             (v.Name),
				"namespace":        (v.Namespace),
				"server_address":   (v.ServerAddress),
				"tags":             convertTagsToMap(v.Tags),
			},
		})
	case *sdm.VaultTokenStore:
		localV, ok := localVersion.(*sdm.VaultTokenStore)
		if !ok {
			localV = &sdm.VaultTokenStore{}
		}
		_ = localV
		d.Set("vault_token", []map[string]interface{}{
			{
				"name":           (v.Name),
				"namespace":      (v.Namespace),
				"server_address": (v.ServerAddress),
				"tags":           convertTagsToMap(v.Tags),
			},
		})
	}
	return nil
}
func resourceSecretStoreUpdate(ctx context.Context, d *schema.ResourceData, cc *sdm.Client) error {

	resp, err := cc.SecretStores().Update(ctx, convertSecretStoreFromResourceData(d))
	if err != nil {
		return fmt.Errorf("cannot update SecretStore %s: %w", d.Id(), err)
	}
	d.SetId(resp.SecretStore.GetID())
	return resourceSecretStoreRead(ctx, d, cc)
}
func resourceSecretStoreDelete(ctx context.Context, d *schema.ResourceData, cc *sdm.Client) error {
	var errNotFound *sdm.NotFoundError
	_, err := cc.SecretStores().Delete(ctx, d.Id())
	if err != nil && errors.As(err, &errNotFound) {
		return nil
	}
	return err
}
