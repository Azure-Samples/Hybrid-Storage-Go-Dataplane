package main

import (
    "context"
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "strings"

    "Hybrid-Storage-Go-Dataplane/resources"
    "Hybrid-Storage-Go-Dataplane/storage"
)

var (
    armEndpoint     = os.Getenv("AZS_ARM_ENDPOINT")
    tenantID        = os.Getenv("AZS_TENANT_ID")
    clientID        = os.Getenv("AZS_CERT_CLIENT_ID")
    pfxPassword     = os.Getenv("AZS_PKCS12_PASSWORD")
    certPath        = os.Getenv("AZS_PKCS12_CERT_PATH")
    subscriptionID  = os.Getenv("AZS_SUBSCRIPTION_ID")
    location        = os.Getenv("AZS_LOCATION")
    blobFileName    = os.Getenv("AZS_BLOB_FILE_NAME")
    blobFileAddress = os.Getenv("AZS_FILE_ADDRESS")

    storageAccountName    = fmt.Sprintf("samplestacc%s", strconv.Itoa(rand.Intn(1000)))
    resourceGroupName     = fmt.Sprintf("stackrg%s", strconv.Itoa(rand.Intn(1000)))
    storageContainerName  = fmt.Sprintf("samplecontainer%s", strconv.Itoa(rand.Intn(1000)))
    storageEndpointSuffix = strings.TrimRight(armEndpoint[strings.Index(armEndpoint, ".")+1:], "/")
)

func main() {
    cntx := context.Background()

    //Create a resource group on Azure Stack
    _, errRgStack := resources.CreateResourceGroup(
        cntx,
        resourceGroupName,
        location,
        certPath,
        armEndpoint,
        tenantID,
        clientID,
        pfxPassword,
        subscriptionID)
    if errRgStack != nil {
        fmt.Println(errRgStack.Error())
        return
    }

    // Create a storge account client
    storageAccountClient := storage.GetStorageAccountsClient(
        tenantID,
        clientID,
        pfxPassword,
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
        blobFileName,
        blobFileAddress)
    if uploadErr != nil {
        fmt.Println(uploadErr.Error())
    }

    return
}
