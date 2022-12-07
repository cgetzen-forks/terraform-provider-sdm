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
			"cyberark_conjur": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "CyberarkConjurStore is currently unstable, and its API may change, or it may be removed, without a major version bump.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"app_url": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
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
					},
				},
			},
			"cyberark_pam": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "CyberarkPAMStore is currently unstable, and its API may change, or it may be removed, without a major version bump.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"app_url": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
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
					},
				},
			},
			"cyberark_pam_experimental": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "CyberarkPAMExperimentalStore is currently unstable, and its API may change, or it may be removed, without a major version bump.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"app_url": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
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
					},
				},
			},
			"delinea_store": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "DelineaStore is currently unstable, and its API may change, or it may be removed, without a major version bump.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the SecretStore.",
						},
						"server_url": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"tags": {
							Type:        schema.TypeMap,
							Elem:        tagsElemType,
							Optional:    true,
							Description: "Tags is a map of key, value pairs.",
						},
						"tenant_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"gcp_store": {
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
						"project_id": {
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
			"vault_approle": {
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
func convertSecretStoreToPlumbing(d *schema.ResourceData) sdm.SecretStore {
	if list := d.Get("aws").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.AWSStore{}
		}
		out := &sdm.AWSStore{
			ID:     d.Id(),
			Name:   convertStringToPlumbing(raw["name"]),
			Region: convertStringToPlumbing(raw["region"]),
			Tags:   convertTagsToPlumbing(raw["tags"]),
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
			Name:     convertStringToPlumbing(raw["name"]),
			Tags:     convertTagsToPlumbing(raw["tags"]),
			VaultUri: convertStringToPlumbing(raw["vault_uri"]),
		}
		return out
	}
	if list := d.Get("cyberark_conjur").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.CyberarkConjurStore{}
		}
		out := &sdm.CyberarkConjurStore{
			ID:     d.Id(),
			AppURL: convertStringToPlumbing(raw["app_url"]),
			Name:   convertStringToPlumbing(raw["name"]),
			Tags:   convertTagsToPlumbing(raw["tags"]),
		}
		return out
	}
	if list := d.Get("cyberark_pam").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.CyberarkPAMStore{}
		}
		out := &sdm.CyberarkPAMStore{
			ID:     d.Id(),
			AppURL: convertStringToPlumbing(raw["app_url"]),
			Name:   convertStringToPlumbing(raw["name"]),
			Tags:   convertTagsToPlumbing(raw["tags"]),
		}
		return out
	}
	if list := d.Get("cyberark_pam_experimental").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.CyberarkPAMExperimentalStore{}
		}
		out := &sdm.CyberarkPAMExperimentalStore{
			ID:     d.Id(),
			AppURL: convertStringToPlumbing(raw["app_url"]),
			Name:   convertStringToPlumbing(raw["name"]),
			Tags:   convertTagsToPlumbing(raw["tags"]),
		}
		return out
	}
	if list := d.Get("delinea_store").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.DelineaStore{}
		}
		out := &sdm.DelineaStore{
			ID:         d.Id(),
			Name:       convertStringToPlumbing(raw["name"]),
			ServerUrl:  convertStringToPlumbing(raw["server_url"]),
			Tags:       convertTagsToPlumbing(raw["tags"]),
			TenantName: convertStringToPlumbing(raw["tenant_name"]),
		}
		return out
	}
	if list := d.Get("gcp_store").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.GCPStore{}
		}
		out := &sdm.GCPStore{
			ID:        d.Id(),
			Name:      convertStringToPlumbing(raw["name"]),
			ProjectID: convertStringToPlumbing(raw["project_id"]),
			Tags:      convertTagsToPlumbing(raw["tags"]),
		}
		return out
	}
	if list := d.Get("vault_approle").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.VaultAppRoleStore{}
		}
		out := &sdm.VaultAppRoleStore{
			ID:            d.Id(),
			Name:          convertStringToPlumbing(raw["name"]),
			Namespace:     convertStringToPlumbing(raw["namespace"]),
			ServerAddress: convertStringToPlumbing(raw["server_address"]),
			Tags:          convertTagsToPlumbing(raw["tags"]),
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
			CACertPath:     convertStringToPlumbing(raw["ca_cert_path"]),
			ClientCertPath: convertStringToPlumbing(raw["client_cert_path"]),
			ClientKeyPath:  convertStringToPlumbing(raw["client_key_path"]),
			Name:           convertStringToPlumbing(raw["name"]),
			Namespace:      convertStringToPlumbing(raw["namespace"]),
			ServerAddress:  convertStringToPlumbing(raw["server_address"]),
			Tags:           convertTagsToPlumbing(raw["tags"]),
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
			Name:          convertStringToPlumbing(raw["name"]),
			Namespace:     convertStringToPlumbing(raw["namespace"]),
			ServerAddress: convertStringToPlumbing(raw["server_address"]),
			Tags:          convertTagsToPlumbing(raw["tags"]),
		}
		return out
	}
	return nil
}

func resourceSecretStoreCreate(ctx context.Context, d *schema.ResourceData, cc *sdm.Client) error {
	localVersion := convertSecretStoreToPlumbing(d)

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
				"tags":   convertTagsToPorcelain(v.Tags),
			},
		})
	case *sdm.AzureStore:
		localV, _ := localVersion.(*sdm.AzureStore)
		_ = localV
		d.Set("azure_store", []map[string]interface{}{
			{
				"name":      (v.Name),
				"tags":      convertTagsToPorcelain(v.Tags),
				"vault_uri": (v.VaultUri),
			},
		})
	case *sdm.CyberarkConjurStore:
		localV, _ := localVersion.(*sdm.CyberarkConjurStore)
		_ = localV
		d.Set("cyberark_conjur", []map[string]interface{}{
			{
				"app_url": (v.AppURL),
				"name":    (v.Name),
				"tags":    convertTagsToPorcelain(v.Tags),
			},
		})
	case *sdm.CyberarkPAMStore:
		localV, _ := localVersion.(*sdm.CyberarkPAMStore)
		_ = localV
		d.Set("cyberark_pam", []map[string]interface{}{
			{
				"app_url": (v.AppURL),
				"name":    (v.Name),
				"tags":    convertTagsToPorcelain(v.Tags),
			},
		})
	case *sdm.CyberarkPAMExperimentalStore:
		localV, _ := localVersion.(*sdm.CyberarkPAMExperimentalStore)
		_ = localV
		d.Set("cyberark_pam_experimental", []map[string]interface{}{
			{
				"app_url": (v.AppURL),
				"name":    (v.Name),
				"tags":    convertTagsToPorcelain(v.Tags),
			},
		})
	case *sdm.DelineaStore:
		localV, _ := localVersion.(*sdm.DelineaStore)
		_ = localV
		d.Set("delinea_store", []map[string]interface{}{
			{
				"name":        (v.Name),
				"server_url":  (v.ServerUrl),
				"tags":        convertTagsToPorcelain(v.Tags),
				"tenant_name": (v.TenantName),
			},
		})
	case *sdm.GCPStore:
		localV, _ := localVersion.(*sdm.GCPStore)
		_ = localV
		d.Set("gcp_store", []map[string]interface{}{
			{
				"name":       (v.Name),
				"project_id": (v.ProjectID),
				"tags":       convertTagsToPorcelain(v.Tags),
			},
		})
	case *sdm.VaultAppRoleStore:
		localV, _ := localVersion.(*sdm.VaultAppRoleStore)
		_ = localV
		d.Set("vault_approle", []map[string]interface{}{
			{
				"name":           (v.Name),
				"namespace":      (v.Namespace),
				"server_address": (v.ServerAddress),
				"tags":           convertTagsToPorcelain(v.Tags),
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
				"tags":             convertTagsToPorcelain(v.Tags),
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
				"tags":           convertTagsToPorcelain(v.Tags),
			},
		})
	}
	return nil
}

func resourceSecretStoreRead(ctx context.Context, d *schema.ResourceData, cc *sdm.Client) error {
	localVersion := convertSecretStoreToPlumbing(d)
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
				"tags":   convertTagsToPorcelain(v.Tags),
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
				"tags":      convertTagsToPorcelain(v.Tags),
				"vault_uri": (v.VaultUri),
			},
		})
	case *sdm.CyberarkConjurStore:
		localV, ok := localVersion.(*sdm.CyberarkConjurStore)
		if !ok {
			localV = &sdm.CyberarkConjurStore{}
		}
		_ = localV
		d.Set("cyberark_conjur", []map[string]interface{}{
			{
				"app_url": (v.AppURL),
				"name":    (v.Name),
				"tags":    convertTagsToPorcelain(v.Tags),
			},
		})
	case *sdm.CyberarkPAMStore:
		localV, ok := localVersion.(*sdm.CyberarkPAMStore)
		if !ok {
			localV = &sdm.CyberarkPAMStore{}
		}
		_ = localV
		d.Set("cyberark_pam", []map[string]interface{}{
			{
				"app_url": (v.AppURL),
				"name":    (v.Name),
				"tags":    convertTagsToPorcelain(v.Tags),
			},
		})
	case *sdm.CyberarkPAMExperimentalStore:
		localV, ok := localVersion.(*sdm.CyberarkPAMExperimentalStore)
		if !ok {
			localV = &sdm.CyberarkPAMExperimentalStore{}
		}
		_ = localV
		d.Set("cyberark_pam_experimental", []map[string]interface{}{
			{
				"app_url": (v.AppURL),
				"name":    (v.Name),
				"tags":    convertTagsToPorcelain(v.Tags),
			},
		})
	case *sdm.DelineaStore:
		localV, ok := localVersion.(*sdm.DelineaStore)
		if !ok {
			localV = &sdm.DelineaStore{}
		}
		_ = localV
		d.Set("delinea_store", []map[string]interface{}{
			{
				"name":        (v.Name),
				"server_url":  (v.ServerUrl),
				"tags":        convertTagsToPorcelain(v.Tags),
				"tenant_name": (v.TenantName),
			},
		})
	case *sdm.GCPStore:
		localV, ok := localVersion.(*sdm.GCPStore)
		if !ok {
			localV = &sdm.GCPStore{}
		}
		_ = localV
		d.Set("gcp_store", []map[string]interface{}{
			{
				"name":       (v.Name),
				"project_id": (v.ProjectID),
				"tags":       convertTagsToPorcelain(v.Tags),
			},
		})
	case *sdm.VaultAppRoleStore:
		localV, ok := localVersion.(*sdm.VaultAppRoleStore)
		if !ok {
			localV = &sdm.VaultAppRoleStore{}
		}
		_ = localV
		d.Set("vault_approle", []map[string]interface{}{
			{
				"name":           (v.Name),
				"namespace":      (v.Namespace),
				"server_address": (v.ServerAddress),
				"tags":           convertTagsToPorcelain(v.Tags),
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
				"tags":             convertTagsToPorcelain(v.Tags),
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
				"tags":           convertTagsToPorcelain(v.Tags),
			},
		})
	}
	return nil
}
func resourceSecretStoreUpdate(ctx context.Context, d *schema.ResourceData, cc *sdm.Client) error {

	resp, err := cc.SecretStores().Update(ctx, convertSecretStoreToPlumbing(d))
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
