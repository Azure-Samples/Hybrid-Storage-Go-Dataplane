package resources

import (
	"context"
	"fmt"
	"log"

	"Hybrid-Storage-Go-Dataplane/iam"

	"github.com/Azure/azure-sdk-for-go/profiles/2020-09-01/resources/mgmt/resources"
	"github.com/Azure/go-autorest/autorest"
)

const (
	errorPrefix = "Cannot create resource group, reason: %v"
)

func getResourceGroupsClient(certPath, armEndpoint, tenantID, clientID, certPass, subscriptionID string) resources.GroupsClient {
	token, err := iam.GetResourceManagementToken(tenantID, clientID, certPass, armEndpoint, certPath)
	if err != nil {
		log.Fatal(fmt.Sprintf(errorPrefix, fmt.Sprintf("Cannot generate token. Error details: %v.", err)))
	}

	groupsClient := resources.NewGroupsClientWithBaseURI(armEndpoint, subscriptionID)
	groupsClient.Authorizer = autorest.NewBearerAuthorizer(token)
	groupsClient.UserAgent = "GoSdkCertDataplaneSample"
	return groupsClient
}

// CreateResourceGroup creates resource group
func CreateResourceGroup(cntx context.Context, rgname, location, certPath, armEndpoint, tenantID, clientID, certPass, subscriptionID string) (name *string, err error) {
	groupClient := getResourceGroupsClient(certPath, armEndpoint, tenantID, clientID, certPass, subscriptionID)
	rg, errRg := groupClient.CreateOrUpdate(cntx, rgname, resources.Group{Location: &location})
	if errRg == nil {
		name = rg.Name
	}
	err = errRg
	return name, err
}
