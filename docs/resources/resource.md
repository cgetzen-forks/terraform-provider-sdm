---
page_title: "SDM: sdm_resource"
description: |-
  Provides settings for Resource.
layout: “sdm”
sidebar_current: “docs-sdm-resource-resource"
---
# Resource: sdm_resource

A Resource is a database, server, cluster, website, or cloud that strongDM
 delegates access to.
## Example Usage

```hcl
resource "sdm_resource" "redis-test" {
    redis {
        name = "redis-test"
        hostname = "example.com"
        port_override = 4020
        tags = {
            region = "us-west"
            env = "dev"
        }
    }
}

resource "sdm_resource" "postgres-test" {
    postgres {
        name = "postgres-test"
        hostname = "example.com"
        database = "my-db"
        username = "admin"
        password = "hunter2"
        port = 5432
        tags = {
            region = "us-west"
            env = "dev"
        }
    }
}
```
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.
## Providing Credentials

Credentials can be provided to resources in two forms:
- As raw text, which will not be returned to the terraform client on import or on loading state from StrongDM, but may be stored in the terraform state itself.
- As a path to a credential in a Secret Store, which will be returned on import. e.g. /path/to/secret?key=password&encoding=base64

All credentials must be either raw or Secret Store paths, depending on whether the resource has a Secret Store ID provided. In both cases, some credentials may be optional depending on the resource subtype.
## Argument Reference
The following arguments are supported by the Resource resource:
* aks:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `certificate_authority` - (Optional) The CA to authenticate TLS connections with.
	* `client_certificate` - (Optional) The certificate to authenticate TLS connections with.
	* `client_key` - (Optional) The key to authenticate TLS connections with.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `remote_identity_group_id` - (Optional) The ID of the remote identity group to use for remote identity connections.
	* `remote_identity_healthcheck_username` - (Optional) The username to use for healthchecks, when clients otherwise connect with their own remote identity username.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* aks_basic_auth:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port` - (Required) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* aks_service_account:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `remote_identity_group_id` - (Optional) The ID of the remote identity group to use for remote identity connections.
	* `remote_identity_healthcheck_username` - (Optional) The username to use for healthchecks, when clients otherwise connect with their own remote identity username.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `token` - (Required, either in plaintext, or as a secret store path) The API token to authenticate with.
* aks_service_account_user_impersonation:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `token` - (Required, either in plaintext, or as a secret store path) The API token to authenticate with.
* aks_user_impersonation:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `certificate_authority` - (Optional) The CA to authenticate TLS connections with.
	* `client_certificate` - (Optional) The certificate to authenticate TLS connections with.
	* `client_key` - (Optional) The key to authenticate TLS connections with.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* amazon_eks:
	* `access_key` - (Required, either in plaintext, or as a secret store path) The Access Key ID to use to authenticate.
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `certificate_authority` - (Optional) The CA to authenticate TLS connections with.
	* `cluster_name` - (Required) The name of the cluster to connect to.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) The endpoint to dial.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `region` - (Required) The AWS region to connect to e.g. us-east-1.
	* `remote_identity_group_id` - (Optional) The ID of the remote identity group to use for remote identity connections.
	* `remote_identity_healthcheck_username` - (Optional) The username to use for healthchecks, when clients otherwise connect with their own remote identity username.
	* `role_arn` - (Optional) The role to assume after logging in.
	* `role_external_id` - (Optional) The external ID to associate with assume role requests. Does nothing if a role ARN is not provided.
	* `secret_access_key` - (Required, either in plaintext, or as a secret store path) The Secret Access Key to use to authenticate.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* amazon_eks_instance_profile:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `certificate_authority` - (Optional) The CA to authenticate TLS connections with.
	* `cluster_name` - (Required) The name of the cluster to connect to.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) The endpoint to dial.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `region` - (Required) The AWS region to connect to e.g. us-east-1.
	* `remote_identity_group_id` - (Optional) The ID of the remote identity group to use for remote identity connections.
	* `remote_identity_healthcheck_username` - (Optional) The username to use for healthchecks, when clients otherwise connect with their own remote identity username.
	* `role_arn` - (Optional) The role to assume after logging in.
	* `role_external_id` - (Optional) The external ID to associate with assume role requests. Does nothing if a role ARN is not provided.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* amazon_eks_instance_profile_user_impersonation:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `certificate_authority` - (Optional) The CA to authenticate TLS connections with.
	* `cluster_name` - (Required) The name of the cluster to connect to.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) The endpoint to dial.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `region` - (Required) The AWS region to connect to e.g. us-east-1.
	* `remote_identity_group_id` - (Optional) The ID of the remote identity group to use for remote identity connections.
	* `remote_identity_healthcheck_username` - (Optional) The username to use for healthchecks, when clients otherwise connect with their own remote identity username.
	* `role_arn` - (Optional) The role to assume after logging in.
	* `role_external_id` - (Optional) The external ID to associate with assume role requests. Does nothing if a role ARN is not provided.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* amazon_eks_user_impersonation:
	* `access_key` - (Required, either in plaintext, or as a secret store path) The Access Key ID to use to authenticate.
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `certificate_authority` - (Optional) The CA to authenticate TLS connections with.
	* `cluster_name` - (Required) The name of the cluster to connect to.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) The endpoint to dial.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `region` - (Required) The AWS region to connect to e.g. us-east-1.
	* `role_arn` - (Optional) The role to assume after logging in.
	* `role_external_id` - (Optional) The external ID to associate with assume role requests. Does nothing if a role ARN is not provided.
	* `secret_access_key` - (Required, either in plaintext, or as a secret store path) The Secret Access Key to use to authenticate.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* amazon_es:
	* `access_key` - (Optional) The Access Key ID to use to authenticate.
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Optional) The endpoint to dial e.g. search-?.region.es.amazonaws.com"
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `region` - (Required) The AWS region to connect to e.g. us-east-1.
	* `role_arn` - (Optional) The role to assume after logging in.
	* `role_external_id` - (Optional) The external ID to associate with assume role requests. Does nothing if a role ARN is not provided.
	* `secret_access_key` - (Optional) The Secret Access Key to use to authenticate.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* amazonmq_amqp_091:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) If set, TLS must be used to connect to this resource.
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* athena:
	* `access_key` - (Required, either in plaintext, or as a secret store path) The Access Key ID to use to authenticate.
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `output` - (Required) The AWS S3 output location.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `region` - (Optional) The AWS region to connect to e.g. us-east-1.
	* `role_arn` - (Optional) The role to assume after logging in.
	* `role_external_id` - (Optional) The external ID to associate with assume role requests. Does nothing if a role ARN is not provided.
	* `secret_access_key` - (Required, either in plaintext, or as a secret store path) The Secret Access Key to use to authenticate.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* aurora_mysql:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `database` - (Required) The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `require_native_auth` - (Optional) Whether native auth (mysql_native_password) is used for all connections (for backwards compatibility)
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `use_azure_single_server_usernames` - (Optional) If true, appends the hostname to the username when hitting a database.azure.com address
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* aurora_postgres:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `database` - (Required) The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) If set, the database configured cannot be changed by users. This setting is not recommended for most use cases, as some clients will insist their database has changed when it has not, leading to user confusion.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* aws:
	* `access_key` - (Required, either in plaintext, or as a secret store path) The Access Key ID to use to authenticate.
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_region` - (Required) The AWS region healthcheck requests should attempt to connect to.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `role_arn` - (Optional) The role to assume after logging in.
	* `role_external_id` - (Optional) The external ID to associate with assume role requests. Does nothing if a role ARN is not provided.
	* `secret_access_key` - (Required, either in plaintext, or as a secret store path) The Secret Access Key to use to authenticate.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* aws_console:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `enable_env_variables` - (Optional) If true, prefer environment variables to authenticate connection even if EC2 roles are configured.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `region` - (Required) The AWS region to connect to.
	* `remote_identity_group_id` - (Optional) The ID of the remote identity group to use for remote identity connections.
	* `remote_identity_healthcheck_username` - (Optional) The username to use for healthchecks, when clients otherwise connect with their own remote identity username.
	* `role_arn` - (Required, either in plaintext, or as a secret store path) The role to assume after logging in.
	* `role_external_id` - (Optional) The external ID to associate with assume role requests. Does nothing if a role ARN is not provided.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `session_expiry` - (Optional) The length of time in seconds AWS console sessions will live before needing to reauthenticate.
	* `subdomain` - (Required) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* aws_console_static_key_pair:
	* `access_key` - (Required, either in plaintext, or as a secret store path) The Access Key ID to authenticate with.
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `region` - (Required) The AWS region to connect to.
	* `remote_identity_group_id` - (Optional) The ID of the remote identity group to use for remote identity connections.
	* `remote_identity_healthcheck_username` - (Optional) The username to use for healthchecks, when clients otherwise connect with their own remote identity username.
	* `role_arn` - (Required, either in plaintext, or as a secret store path) The role to assume after logging in.
	* `role_external_id` - (Optional) The external ID to associate with assume role requests. Does nothing if a role ARN is not provided.
	* `secret_access_key` - (Required, either in plaintext, or as a secret store path) The Secret Access Key to authenticate with.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `session_expiry` - (Optional) The length of time in seconds AWS console sessions will live before needing to reauthenticate.
	* `subdomain` - (Required) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* azure:
	* `app_id` - (Required, either in plaintext, or as a secret store path) The application ID to authenticate with.
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tenant_id` - (Required, either in plaintext, or as a secret store path) The tenant ID to authenticate to.
* azure_certificate:
	* `app_id` - (Required, either in plaintext, or as a secret store path) The application ID to authenticate with.
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `client_certificate` - (Required, either in plaintext, or as a secret store path) The service Principal certificate file, both private and public key included.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tenant_id` - (Required, either in plaintext, or as a secret store path) The tenant ID to authenticate to.
* azure_mysql:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `database` - (Required) The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `require_native_auth` - (Optional) Whether native auth (mysql_native_password) is used for all connections (for backwards compatibility)
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `use_azure_single_server_usernames` - (Optional) If true, appends the hostname to the username when hitting a database.azure.com address
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* azure_postgres:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `database` - (Required) The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) If set, the database configured cannot be changed by users. This setting is not recommended for most use cases, as some clients will insist their database has changed when it has not, leading to user confusion.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* azure_postgres_managed_identity:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `database` - (Required) The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) If set, the database configured cannot be changed by users. This setting is not recommended for most use cases, as some clients will insist their database has changed when it has not, leading to user confusion.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `use_azure_single_server_usernames` - (Optional) If true, appends the hostname to the username when hitting a database.azure.com address
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* big_query:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) The endpoint to dial.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `private_key` - (Required, either in plaintext, or as a secret store path) The JSON Private key to authenticate with.
	* `project` - (Required) The project to connect to.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) The username to authenticate with.
* cassandra:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) If set, TLS must be used to connect to this resource.
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* citus:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `database` - (Required) The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) If set, the database configured cannot be changed by users. This setting is not recommended for most use cases, as some clients will insist their database has changed when it has not, leading to user confusion.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* clustrix:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `database` - (Required) The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `require_native_auth` - (Optional) Whether native auth (mysql_native_password) is used for all connections (for backwards compatibility)
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `use_azure_single_server_usernames` - (Optional) If true, appends the hostname to the username when hitting a database.azure.com address
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* cockroach:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `database` - (Required) The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) If set, the database configured cannot be changed by users. This setting is not recommended for most use cases, as some clients will insist their database has changed when it has not, leading to user confusion.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* db_2_i:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port` - (Required) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) If set, TLS must be used to connect to this resource.
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* db_2_luw:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `database` - (Required) The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* document_db_host:
	* `auth_database` - (Required) The authentication database to use.
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) The username to authenticate with.
* document_db_replica_set:
	* `auth_database` - (Required) The authentication database to use.
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `connect_to_replica` - (Optional) Set to connect to a replica instead of the primary node.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) Hostname must contain the hostname/port pairs of all instances in the replica set separated by commas.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) The password to authenticate with.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `replica_set` - (Required) The name of the mongo replicaset.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) The username to authenticate with.
* druid:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) The username to authenticate with.
* dynamo_db:
	* `access_key` - (Required, either in plaintext, or as a secret store path) The Access Key ID to use to authenticate.
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) The endpoint to dial e.g. dynamodb.region.amazonaws.com
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `region` - (Required) The region to authenticate requests against e.g. us-east-1
	* `role_arn` - (Optional) The role to assume after logging in.
	* `role_external_id` - (Optional) The external ID to associate with assume role requests. Does nothing if a role ARN is not provided.
	* `secret_access_key` - (Required, either in plaintext, or as a secret store path) The Secret Access Key to use to authenticate.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* elastic:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) If set, TLS must be used to connect to this resource.
	* `username` - (Optional) The username to authenticate with.
* elasticache_redis:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) If set, TLS must be used to connect to this resource.
	* `username` - (Optional) The username to authenticate with.
* gcp:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `keyfile` - (Required, either in plaintext, or as a secret store path) The service account keyfile to authenticate with.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `scopes` - (Required) Space separated scopes that this login should assume into when authenticating.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* google_gke:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `certificate_authority` - (Optional) The CA to authenticate TLS connections with.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) The endpoint to dial.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `remote_identity_group_id` - (Optional) The ID of the remote identity group to use for remote identity connections.
	* `remote_identity_healthcheck_username` - (Optional) The username to use for healthchecks, when clients otherwise connect with their own remote identity username.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `service_account_key` - (Required, either in plaintext, or as a secret store path) The service account key to authenticate with.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* google_gke_user_impersonation:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `certificate_authority` - (Optional) The CA to authenticate TLS connections with.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) The endpoint to dial.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `service_account_key` - (Required, either in plaintext, or as a secret store path) The service account key to authenticate with.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* greenplum:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `database` - (Required) The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) If set, the database configured cannot be changed by users. This setting is not recommended for most use cases, as some clients will insist their database has changed when it has not, leading to user confusion.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* http_auth:
	* `auth_header` - (Required, either in plaintext, or as a secret store path) The content to set as the authorization header.
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `default_path` - (Optional) Automatically redirect to this path upon connecting.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `headers_blacklist` - (Optional) Header names (e.g. Authorization), to omit from logs.
	* `healthcheck_path` - (Required) This path will be used to check the health of your site.
	* `host_override` - (Optional) The host header will be overwritten with this field if provided.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Required) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `url` - (Required) The base address of your website without the path.
* http_basic_auth:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `default_path` - (Optional) Automatically redirect to this path upon connecting.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `headers_blacklist` - (Optional) Header names (e.g. Authorization), to omit from logs.
	* `healthcheck_path` - (Required) This path will be used to check the health of your site.
	* `host_override` - (Optional) The host header will be overwritten with this field if provided.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) The password to authenticate with.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Required) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `url` - (Required) The base address of your website without the path.
	* `username` - (Optional) The username to authenticate with.
* http_no_auth:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `default_path` - (Optional) Automatically redirect to this path upon connecting.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `headers_blacklist` - (Optional) Header names (e.g. Authorization), to omit from logs.
	* `healthcheck_path` - (Required) This path will be used to check the health of your site.
	* `host_override` - (Optional) The host header will be overwritten with this field if provided.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Required) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `url` - (Required) The base address of your website without the path.
* kubernetes:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `certificate_authority` - (Optional) The CA to authenticate TLS connections with.
	* `client_certificate` - (Optional) The certificate to authenticate TLS connections with.
	* `client_key` - (Optional) The key to authenticate TLS connections with.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `remote_identity_group_id` - (Optional) The ID of the remote identity group to use for remote identity connections.
	* `remote_identity_healthcheck_username` - (Optional) The username to use for healthchecks, when clients otherwise connect with their own remote identity username.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* kubernetes_basic_auth:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port` - (Required) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* kubernetes_service_account:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `remote_identity_group_id` - (Optional) The ID of the remote identity group to use for remote identity connections.
	* `remote_identity_healthcheck_username` - (Optional) The username to use for healthchecks, when clients otherwise connect with their own remote identity username.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `token` - (Required, either in plaintext, or as a secret store path) The API token to authenticate with.
* kubernetes_service_account_user_impersonation:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `token` - (Required, either in plaintext, or as a secret store path) The API token to authenticate with.
* kubernetes_user_impersonation:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `certificate_authority` - (Optional) The CA to authenticate TLS connections with.
	* `client_certificate` - (Optional) The certificate to authenticate TLS connections with.
	* `client_key` - (Optional) The key to authenticate TLS connections with.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* maria:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `database` - (Required) The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `require_native_auth` - (Optional) Whether native auth (mysql_native_password) is used for all connections (for backwards compatibility)
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `use_azure_single_server_usernames` - (Optional) If true, appends the hostname to the username when hitting a database.azure.com address
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* memcached:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* memsql:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `database` - (Required) The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `require_native_auth` - (Optional) Whether native auth (mysql_native_password) is used for all connections (for backwards compatibility)
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `use_azure_single_server_usernames` - (Optional) If true, appends the hostname to the username when hitting a database.azure.com address
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* mongo_host:
	* `auth_database` - (Required) The authentication database to use.
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) If set, TLS must be used to connect to this resource.
	* `username` - (Optional) The username to authenticate with.
* mongo_legacy_host:
	* `auth_database` - (Required) The authentication database to use.
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `replica_set` - (Optional) The name of the mongo replicaset.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) If set, TLS must be used to connect to this resource.
	* `username` - (Optional) The username to authenticate with.
* mongo_legacy_replicaset:
	* `auth_database` - (Required) The authentication database to use.
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `connect_to_replica` - (Optional) Set to connect to a replica instead of the primary node.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `replica_set` - (Required) The name of the mongo replicaset.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) If set, TLS must be used to connect to this resource.
	* `username` - (Optional) The username to authenticate with.
* mongo_replica_set:
	* `auth_database` - (Required) The authentication database to use.
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `connect_to_replica` - (Optional) Set to connect to a replica instead of the primary node.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `replica_set` - (Required) The name of the mongo replicaset.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) If set, TLS must be used to connect to this resource.
	* `username` - (Optional) The username to authenticate with.
* mongo_sharded_cluster:
	* `auth_database` - (Required) The authentication database to use.
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) The password to authenticate with.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) If set, TLS must be used to connect to this resource.
	* `username` - (Optional) The username to authenticate with.
* mtls_mysql:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `certificate_authority` - (Optional) The CA to authenticate TLS connections with.
	* `client_certificate` - (Required, either in plaintext, or as a secret store path) The certificate to authenticate TLS connections with.
	* `client_key` - (Required, either in plaintext, or as a secret store path) The key to authenticate TLS connections with.
	* `database` - (Required) The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `require_native_auth` - (Optional) Whether native auth (mysql_native_password) is used for all connections (for backwards compatibility)
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `server_name` - (Optional) Server name for TLS verification (unverified by StrongDM if empty)
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `use_azure_single_server_usernames` - (Optional) If true, appends the hostname to the username when hitting a database.azure.com address
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* mtls_postgres:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `certificate_authority` - (Optional) The CA to authenticate TLS connections with.
	* `client_certificate` - (Required, either in plaintext, or as a secret store path) The certificate to authenticate TLS connections with.
	* `client_key` - (Required, either in plaintext, or as a secret store path) The key to authenticate TLS connections with.
	* `database` - (Required) The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) If set, the database configured cannot be changed by users. This setting is not recommended for most use cases, as some clients will insist their database has changed when it has not, leading to user confusion.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `server_name` - (Optional) Server name for TLS verification (unverified by StrongDM if empty)
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* mysql:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `database` - (Required) The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `require_native_auth` - (Optional) Whether native auth (mysql_native_password) is used for all connections (for backwards compatibility)
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `use_azure_single_server_usernames` - (Optional) If true, appends the hostname to the username when hitting a database.azure.com address
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* neptune:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) The neptune endpoint to connect to as in endpoint.region.neptune.amazonaws.com
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* neptune_iam:
	* `access_key` - (Required, either in plaintext, or as a secret store path) The Access Key ID to use to authenticate.
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) The neptune endpoint to connect to as in endpoint.region.neptune.amazonaws.com
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `region` - (Required) The AWS region to connect to.
	* `role_arn` - (Optional) The role to assume after logging in.
	* `role_external_id` - (Optional) The external ID to associate with assume role requests. Does nothing if a role ARN is not provided.
	* `secret_access_key` - (Required, either in plaintext, or as a secret store path) The Secret Access Key to use to authenticate.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* oracle:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `database` - (Required) The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port` - (Required) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) If set, TLS must be used to connect to this resource.
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* postgres:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `database` - (Required) The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) If set, the database configured cannot be changed by users. This setting is not recommended for most use cases, as some clients will insist their database has changed when it has not, leading to user confusion.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* presto:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `database` - (Required) The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) If set, TLS must be used to connect to this resource.
	* `username` - (Optional) The username to authenticate with.
* rabbitmq_amqp_091:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) If set, TLS must be used to connect to this resource.
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* raw_tcp:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* rdp:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `downgrade_nla_connections` - (Optional) When set, network level authentication will not be used. May resolve unexpected authentication errors to older servers. When set, healthchecks cannot detect if a provided username / password pair is correct.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* redis:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) If set, TLS must be used to connect to this resource.
	* `username` - (Optional) The username to authenticate with.
* redshift:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `database` - (Required) The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) If set, the database configured cannot be changed by users. This setting is not recommended for most use cases, as some clients will insist their database has changed when it has not, leading to user confusion.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* single_store:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `database` - (Required) The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `require_native_auth` - (Optional) Whether native auth (mysql_native_password) is used for all connections (for backwards compatibility)
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `use_azure_single_server_usernames` - (Optional) If true, appends the hostname to the username when hitting a database.azure.com address
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* snowflake:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `database` - (Required) The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `schema` - (Optional) The schema to provide on authentication.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* snowsight:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_username` - (Required) The StrongDM user email to use for healthchecks.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `saml_metadata` - (Required, either in plaintext, or as a secret store path) The Metadata for your snowflake IDP integration
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Required) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* sql_server:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `database` - (Required) The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) If set, the database configured cannot be changed by users. This setting is not recommended for most use cases, as some clients will insist their database has changed when it has not, leading to user confusion.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `schema` - (Optional) The Schema to use to direct initial requests.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* sql_server_azure_ad:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `client_id` - (Required, either in plaintext, or as a secret store path) The Azure AD application (client) ID with which to authenticate.
	* `database` - (Required) The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) If set, the database configured cannot be changed by users. This setting is not recommended for most use cases, as some clients will insist their database has changed when it has not, leading to user confusion.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `schema` - (Optional) The Schema to use to direct initial requests.
	* `secret` - (Required, either in plaintext, or as a secret store path) The Azure AD client secret (application password) with which to authenticate.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tenant_id` - (Required, either in plaintext, or as a secret store path) The Azure AD directory (tenant) ID with which to authenticate.
* sql_server_kerberos_ad:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `database` - (Required) The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `keytab` - (Required, either in plaintext, or as a secret store path) The keytab file in base64 format containing an entry with the principal name (username@realm) and key version number with which to authenticate.
	* `krb_config` - (Required, either in plaintext, or as a secret store path) The Kerberos 5 configuration file (krb5.conf) specifying the Active Directory server (KDC) for the configured realm.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) If set, the database configured cannot be changed by users. This setting is not recommended for most use cases, as some clients will insist their database has changed when it has not, leading to user confusion.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `realm` - (Required, either in plaintext, or as a secret store path) The Active Directory domain (realm) to which the configured username belongs.
	* `schema` - (Optional) The Schema to use to direct initial requests.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `server_spn` - (Required, either in plaintext, or as a secret store path) The Service Principal Name of the Microsoft SQL Server instance in Active Directory.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* ssh:
	* `allow_deprecated_key_exchanges` - (Optional) Whether deprecated, insecure key exchanges are allowed for use to connect to the target ssh server.
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `key_type` - (Optional) The key type to use e.g. rsa-2048 or ed25519
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) The port to dial to initiate a connection from the egress node to this resource.
	* `port_forwarding` - (Optional) Whether port forwarding is allowed through this server.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* ssh_cert:
	* `allow_deprecated_key_exchanges` - (Optional) Whether deprecated, insecure key exchanges are allowed for use to connect to the target ssh server.
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `key_type` - (Optional) The key type to use e.g. rsa-2048 or ed25519
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) The port to dial to initiate a connection from the egress node to this resource.
	* `port_forwarding` - (Optional) Whether port forwarding is allowed through this server.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `remote_identity_group_id` - (Optional) The ID of the remote identity group to use for remote identity connections.
	* `remote_identity_healthcheck_username` - (Optional) The username to use for healthchecks, when clients otherwise connect with their own remote identity username.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* ssh_customer_key:
	* `allow_deprecated_key_exchanges` - (Optional) Whether deprecated, insecure key exchanges are allowed for use to connect to the target ssh server.
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) The port to dial to initiate a connection from the egress node to this resource.
	* `port_forwarding` - (Optional) Whether port forwarding is allowed through this server.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `private_key` - (Required, either in plaintext, or as a secret store path) The private key used to authenticate with the server.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* sybase:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* sybase_iq:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* teradata:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) The username to authenticate with.
* trino:
	* `bind_interface` - (Optional) The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
	* `database` - (Required) The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) The host to dial to initiate a connection from the egress node to this resource.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) The password to authenticate with.
	* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
	* `port_override` - (Optional) The local port used by clients to connect to this resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) The username to authenticate with.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the Resource resource:
* `id` - A unique identifier for the Resource resource.
* ssh:
	* `public_key` - The public key to append to a server's authorized keys. This will be generated after resource creation.
## Import
Resource can be imported using the id, e.g.,

```
$ terraform import sdm_resource.example rs-12345678
```
