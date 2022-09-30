// Copyright 2020 StrongDM Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Package sdm implements an API client to strongDM restful API.
package sdm

// This file was generated by protogen. DO NOT EDIT.

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	plumbing "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

const (
	defaultAPIHost   = "api.strongdm.com:443"
	apiVersion       = "2021-08-23"
	defaultUserAgent = "strongdm-sdk-go/3.3.0"
)

var _ = metadata.Pairs

// Client is the strongDM API client implementation.
type Client struct {
	testOptionsMu sync.RWMutex
	testOptions   map[string]interface{}

	apiHost               string
	apiToken              string
	apiSecret             []byte
	apiInsecureTransport  bool
	apiTLSConfig          *tls.Config
	exposeRateLimitErrors bool
	userAgent             string
	disableSigning        bool

	grpcConn *grpc.ClientConn

	maxRetries           int
	baseRetryDelay       time.Duration
	maxRetryDelay        time.Duration
	accountAttachments   *AccountAttachments
	accountGrants        *AccountGrants
	accounts             *Accounts
	controlPanel         *ControlPanel
	nodes                *Nodes
	remoteIdentities     *RemoteIdentities
	remoteIdentityGroups *RemoteIdentityGroups
	resources            *Resources
	roles                *Roles
	secretStores         *SecretStores
}

// New creates a new strongDM API client.
func New(token, secret string, opts ...ClientOption) (*Client, error) {
	token = strings.TrimSpace(token)
	secret = strings.TrimSpace(secret)
	decodedSecret, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return nil, convertErrorToPorcelain(fmt.Errorf("invalid secret: %w", err))
	}

	client := &Client{
		apiHost:        defaultAPIHost,
		maxRetries:     defaultMaxRetries,
		baseRetryDelay: defaultBaseRetryDelay,
		maxRetryDelay:  defaultMaxRetryDelay,
		testOptions:    map[string]interface{}{},
		apiToken:       token,
		apiSecret:      decodedSecret,
		userAgent:      defaultUserAgent,
	}

	for _, opt := range opts {
		opt(client)
	}

	var dialOpt grpc.DialOption
	if client.apiInsecureTransport {
		dialOpt = grpc.WithInsecure()
	} else if client.apiTLSConfig != nil {
		dialOpt = grpc.WithTransportCredentials(credentials.NewTLS(client.apiTLSConfig))
	} else {
		dialOpt = grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			RootCAs:            nil,
			InsecureSkipVerify: false,
		}))
	}
	cc, err := grpc.Dial(client.apiHost, dialOpt)
	if err != nil {
		return nil, convertErrorToPorcelain(fmt.Errorf("cannot dial API server: %w", err))
	}
	client.grpcConn = cc
	client.accountAttachments = &AccountAttachments{
		client: plumbing.NewAccountAttachmentsClient(client.grpcConn),
		parent: client,
	}
	client.accountGrants = &AccountGrants{
		client: plumbing.NewAccountGrantsClient(client.grpcConn),
		parent: client,
	}
	client.accounts = &Accounts{
		client: plumbing.NewAccountsClient(client.grpcConn),
		parent: client,
	}
	client.controlPanel = &ControlPanel{
		client: plumbing.NewControlPanelClient(client.grpcConn),
		parent: client,
	}
	client.nodes = &Nodes{
		client: plumbing.NewNodesClient(client.grpcConn),
		parent: client,
	}
	client.remoteIdentities = &RemoteIdentities{
		client: plumbing.NewRemoteIdentitiesClient(client.grpcConn),
		parent: client,
	}
	client.remoteIdentityGroups = &RemoteIdentityGroups{
		client: plumbing.NewRemoteIdentityGroupsClient(client.grpcConn),
		parent: client,
	}
	client.resources = &Resources{
		client: plumbing.NewResourcesClient(client.grpcConn),
		parent: client,
	}
	client.roles = &Roles{
		client: plumbing.NewRolesClient(client.grpcConn),
		parent: client,
	}
	client.secretStores = &SecretStores{
		client: plumbing.NewSecretStoresClient(client.grpcConn),
		parent: client,
	}
	return client, nil
}

// A ClientOption is an optional argument to New that can override the created
// client's default behavior.
type ClientOption func(c *Client)

// WithHost causes a Client to make it's calls against the provided host instead
// of against api.strongdm.com.
func WithHost(host string) ClientOption {
	return func(c *Client) {
		c.apiHost = host
	}
}

// WithInsecure enables a Client to talk to an http server instead of an https
// server. This is potentially useful when communicating through a proxy, but
// should be used with care.
func WithInsecure() ClientOption {
	return func(c *Client) {
		c.apiInsecureTransport = true
	}
}

// WithTLSConfig allows customization of the TLS configuration used to
// communicate with the API server.
func WithTLSConfig(cfg *tls.Config) ClientOption {
	return func(c *Client) {
		c.apiTLSConfig = cfg
	}
}

// WithUserAgentExtra modifies the user agent string to include additional identifying
// information for server-side analytics. The intended use is by extension libraries,
// like a terraform provider wrapping this client.
func WithUserAgentExtra(userAgentExtra string) ClientOption {
	return func(c *Client) {
		c.userAgent += " " + userAgentExtra
	}
}

// WithRateLimitRetries configures whether encountered rate limit errors should
// cause this client to sleep and retry (if enabled), or whether those errors should be
// exposed to the code using this client (if disabled). By default, it is enabled.
func WithRateLimitRetries(enabled bool) ClientOption {
	return func(c *Client) {
		c.exposeRateLimitErrors = !enabled
	}
}

// AccountAttachments assign an account to a role.
func (c *Client) AccountAttachments() *AccountAttachments {
	return c.accountAttachments
}

// AccountGrants assign a resource directly to an account, giving the account the permission to connect to that resource.
func (c *Client) AccountGrants() *AccountGrants {
	return c.accountGrants
}

// Accounts are users that have access to strongDM. There are two types of accounts:
// 1. **Users:** humans who are authenticated through username and password or SSO.
// 2. **Service Accounts:** machines that are authenticated using a service token.
func (c *Client) Accounts() *Accounts {
	return c.accounts
}

// ControlPanel contains all administrative controls.
func (c *Client) ControlPanel() *ControlPanel {
	return c.controlPanel
}

// Nodes make up the strongDM network, and allow your users to connect securely to your resources. There are two types of nodes:
// - **Gateways** are the entry points into network. They listen for connection from the strongDM client, and provide access to databases and servers.
// - **Relays** are used to extend the strongDM network into segmented subnets. They provide access to databases and servers but do not listen for incoming connections.
func (c *Client) Nodes() *Nodes {
	return c.nodes
}

// RemoteIdentities assign a resource directly to an account, giving the account the permission to connect to that resource.
func (c *Client) RemoteIdentities() *RemoteIdentities {
	return c.remoteIdentities
}

// A RemoteIdentityGroup is a named grouping of Remote Identities for Accounts.
// An Account's relationship to a RemoteIdentityGroup is defined via RemoteIdentity objects.
func (c *Client) RemoteIdentityGroups() *RemoteIdentityGroups {
	return c.remoteIdentityGroups
}

// Resources are databases, servers, clusters, websites, or clouds that strongDM
// delegates access to.
func (c *Client) Resources() *Resources {
	return c.resources
}

// A Role has a list of access rules which determine which Resources the members
// of the Role have access to. An Account can be a member of multiple Roles via
// AccountAttachments.
func (c *Client) Roles() *Roles {
	return c.roles
}

// SecretStores are servers where resource secrets (passwords, keys) are stored.
func (c *Client) SecretStores() *SecretStores {
	return c.secretStores
}

// Sign returns the signature for the given byte array
func (c *Client) Sign(methodName string, message []byte) string {
	if c.disableSigning {
		return ""
	}
	// Current UTC date
	y, m, d := time.Now().UTC().Date()
	currentUTCDate := fmt.Sprintf("%04d-%02d-%02d", y, m, d)

	signingKey := hmacHelper(c.apiSecret, []byte(currentUTCDate))
	signingKey = hmacHelper(signingKey, []byte("sdm_api_v1"))

	hash := sha256.New()
	hash.Write([]byte(methodName))
	hash.Write([]byte{'\n'})
	hash.Write(message)
	hashedMessage := hash.Sum(nil)

	return base64.StdEncoding.EncodeToString(hmacHelper(signingKey, hashedMessage))
}

func hmacHelper(key, msg []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(msg)
	return mac.Sum(nil)
}

func (c *Client) wrapContext(ctx context.Context, req proto.Message, methodName string) context.Context {
	msg, _ := proto.Marshal(req)
	return metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{
		"x-sdm-authentication": c.apiToken,
		"x-sdm-signature":      c.Sign(methodName, msg),
		"x-sdm-api-version":    apiVersion,
		"x-sdm-user-agent":     c.userAgent,
	}))
}

func (c *Client) testOption(key string) interface{} {
	c.testOptionsMu.RLock()
	defer c.testOptionsMu.RUnlock()
	return c.testOptions[key]
}

// These defaults are taken from AWS. Customization of these values
// is a future step in the API.
const (
	defaultMaxRetries     = 3
	defaultBaseRetryDelay = 30 * time.Millisecond
	defaultMaxRetryDelay  = 300 * time.Second
)

func (c *Client) jitterSleep(iter int) {
	durMax := c.baseRetryDelay * time.Duration(2<<iter)
	if durMax > c.maxRetryDelay {
		durMax = c.maxRetryDelay
	}
	// This is a full jitter, ranging from no delay to the maximum
	// this jittering aims to prevent clients that start and conflict
	// at the same time from retrying at the same intervals
	dur := rand.Intn(int(durMax))
	// TODO: use resource.Retry instead of time.Sleep in Terraform provider
	// see https://github.com/bflad/tfproviderlint/tree/main/passes/R018
	time.Sleep(time.Duration(dur)) //lintignore:R018
}

func (c *Client) shouldRetry(iter int, err error) bool {
	if iter >= c.maxRetries-1 {
		return false
	}
	// Internal and unknown errors should be retried
	// Other error types are likely not brief temporary errors
	if s, ok := status.FromError(err); ok {
		if !c.exposeRateLimitErrors && s.Code() == codes.ResourceExhausted {
			var rateLimitError *RateLimitError
			if err != nil && errors.As(convertErrorToPorcelain(err), &rateLimitError) {
				waitFor := time.Until(rateLimitError.RateLimit.ResetAt)
				// If timezones or clock drift causes this calculation to fail,
				// wait at most one minute.
				if waitFor < 0 || waitFor > time.Minute {
					waitFor = time.Minute
				}
				time.Sleep(waitFor)
			}
			return true
		}
		return s.Code() == codes.Internal || s.Code() == codes.Unavailable
	}
	return true
}
