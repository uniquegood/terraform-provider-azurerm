package client

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/servicebus/mgmt/2021-06-01-preview/servicebus"
	"github.com/hashicorp/terraform-provider-azurerm/internal/common"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/servicebus/sdk/2021-06-01-preview/namespaces"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/servicebus/sdk/2021-06-01-preview/namespacesauthorizationrule"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/servicebus/sdk/2021-06-01-preview/queues"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/servicebus/sdk/2021-06-01-preview/queuesauthorizationrule"
)

type Client struct {
	QueuesClient                  *queues.QueuesClient
	QueuesAuthClient              *queuesauthorizationrule.QueuesAuthorizationRuleClient
	DisasterRecoveryConfigsClient *servicebus.DisasterRecoveryConfigsClient
	NamespacesClient              *servicebus.NamespacesClient
	NamespacesClientP             *namespaces.NamespacesClient
	NamespacesAuthClient          *namespacesauthorizationrule.NamespacesAuthorizationRuleClient
	TopicsClient                  *servicebus.TopicsClient
	SubscriptionsClient           *servicebus.SubscriptionsClient
	SubscriptionRulesClient       *servicebus.RulesClient
}

func NewClient(o *common.ClientOptions) *Client {
	QueuesClient := queues.NewQueuesClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&QueuesClient.Client, o.ResourceManagerAuthorizer)

	QueuesAuthClient := queuesauthorizationrule.NewQueuesAuthorizationRuleClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&QueuesAuthClient.Client, o.ResourceManagerAuthorizer)

	DisasterRecoveryConfigsClient := servicebus.NewDisasterRecoveryConfigsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&DisasterRecoveryConfigsClient.Client, o.ResourceManagerAuthorizer)

	NamespacesClient := servicebus.NewNamespacesClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&NamespacesClient.Client, o.ResourceManagerAuthorizer)

	NamespacesClientP := namespaces.NewNamespacesClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&NamespacesClientP.Client, o.ResourceManagerAuthorizer)

	NamespacesAuthClient := namespacesauthorizationrule.NewNamespacesAuthorizationRuleClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&NamespacesAuthClient.Client, o.ResourceManagerAuthorizer)

	TopicsClient := servicebus.NewTopicsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&TopicsClient.Client, o.ResourceManagerAuthorizer)

	SubscriptionsClient := servicebus.NewSubscriptionsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&SubscriptionsClient.Client, o.ResourceManagerAuthorizer)

	SubscriptionRulesClient := servicebus.NewRulesClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&SubscriptionRulesClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		QueuesClient:                  &QueuesClient,
		QueuesAuthClient:              &QueuesAuthClient,
		DisasterRecoveryConfigsClient: &DisasterRecoveryConfigsClient,
		NamespacesClient:              &NamespacesClient,
		NamespacesClientP:             &NamespacesClientP,
		NamespacesAuthClient:          &NamespacesAuthClient,
		TopicsClient:                  &TopicsClient,
		SubscriptionsClient:           &SubscriptionsClient,
		SubscriptionRulesClient:       &SubscriptionRulesClient,
	}
}
