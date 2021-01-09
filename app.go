package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"Hybrid-Storage-Go-Dataplane/resources"
	"Hybrid-Storage-Go-Dataplane/storage"

	"github.com/Azure/go-autorest/autorest/azure"
)

var (
	armEndpoint     = os.Getenv("AZURE_ARM_ENDPOINT")
	tenantID        = os.Getenv("AZURE_TENANT_ID")
	clientID        = os.Getenv("AZURE_SP_CERT_ID")
	certPass        = os.Getenv("AZURE_SP_CERT_PASS")
	certPath        = os.Getenv("AZURE_SP_CERT_PATH")
	subscriptionID  = os.Getenv("AZURE_SUBSCRIPTION_ID")
	location        = os.Getenv("AZURE_LOCATION")
	blobFileAddress = os.Getenv("AZURE_SAMPLE_FILE_PATH")

	storageAccountName    = "samplestacc"
	resourceGroupName     = "azure-sample-rg"
	storageContainerName  = "samplecontainer"
	storageEndpointSuffix = strings.TrimRight(armEndpoint[strings.Index(armEndpoint, ".")+1:], "/")
)

func main() {
	cntx := context.Background()
	environment, _ := azure.EnvironmentFromURL(armEndpoint)
	splitEndpoint := strings.Split(environment.ActiveDirectoryEndpoint, "/")
	splitEndpointlastIndex := len(splitEndpoint) - 1
	if splitEndpoint[splitEndpointlastIndex] == "adfs" || splitEndpoint[splitEndpointlastIndex] == "adfs/" {
		tenantID = "adfs"
	}
	//Create a resource group on Azure Stack
	_, errRgStack := resources.CreateResourceGroup(
		cntx,
		resourceGroupName,
		location,
		certPath,
		armEndpoint,
		tenantID,
		clientID,
		certPass,
		subscriptionID)
	if errRgStack != nil {
		fmt.Println(errRgStack.Error())
		return
	}

	// Create a storge account client
	storageAccountClient := storage.GetStorageAccountsClient(
		tenantID,
		clientID,
		certPass,
		armEndpoint,
		certPath,
		subscriptionID)

	// Create storage account
	_, errSa := storage.CreateStorageAccount(
		cntx,
		storageAccountClient,
		storageAccountName,
		resourceGroupName,
		location)
	if errSa != nil {
		fmt.Println(errSa.Error())
	}

	dataplaneURL, errDP := storage.GetDataplaneURL(
		cntx,
		storageAccountClient,
		storageEndpointSuffix,
		storageAccountName,
		resourceGroupName,
		storageContainerName)
	if errDP != nil {
		fmt.Println(errDP.Error())
	}

	uploadErr := storage.UploadDataToContainer(
		cntx,
		dataplaneURL,
		blobFileAddress)
	if uploadErr != nil {
		fmt.Println(uploadErr.Error())
	}

	return
}
