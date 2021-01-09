---
page_type: sample
languages:
- go
products:
- azure
description: "The code provided shows how to do the following"
urlFragment: Hybrid-Storage-Go-Dataplane
---

# Hybrid-Storage-GO-Dataplane

The code provided shows how to do the following:

- Create a resource group
- Create a storage account
- Create a container in the storage account
- Upload a file to the container

To see the code to perform these operations,
check out the `main()` function in [app.go](app.go).


## Running this sample
1.  If you don't already have it, [install Golang](https://golang.org/doc/install).

2.  Install Go SDK and its dependencies, [install Go SDK](https://github.com/azure/azure-sdk-for-go) 

3.  Install Go SDK dataplane, [install Go SDK Dataplane](https://github.com/Azure/azure-storage-blob-go/) 

4.  Clone the repository.

    ```
    git clone https://github.com/Azure-Samples/Hybrid-Storage-Go-Dataplane.git
    ```

5.  Move the Hybrid-Storage-Go-Dataplane folder to your $GOPATH/src folder.

6.  Open a Powershell or Bash shell in $GOPATH/src/Hybrid-Storage-Go-Dataplane and enter the following commands:

    ```
    go mod init Hybrid-Storage-Go-Dataplane
    go mod tidy
    go get github.com/Azure/azure-storage-blob-go/azblob@v0.10.0
    ```

    NOTE: The azblob@v0.10.0 version is required for AzureStack.

7.  Create a [service principal using a certificate](https://docs.microsoft.com/en-us/azure/azure-stack/azure-stack-create-service-principals#create-a-service-principal-using-a-certificate) to work against AzureStack. Make sure your service principal has [contributor/owner role](https://docs.microsoft.com/en-us/azure/azure-stack/azure-stack-create-service-principals#assign-role-to-service-principal) on your subscription.

8.  Create a file that you wish to upload to your container, ex: c:\testuploadfile.log

9.  Fill in and export these environment variables into your current shell. 

    ```
    export AZS_ARM_ENDPOINT={your AzureStack Resource Manager Endpoint}
    export AZS_TENANT_ID={your tenant id}
    export AZS_CERT_CLIENT_ID={your service principal client id associated with a .pfx certificate file}
    export AZS_PKCS12_PASSWORD={password for your .pfx file}
    export AZS_PKCS12_CERT_PATH={path to your password protected .pfx file}
    export AZS_SUBSCRIPTION_ID={your subscription id}
    export AZS_LOCATION={your resource location}
    export AZS_BLOB_FILE_NAME={name of the file you want to upload to your container, ex: testuploadfile.log}
    export AZS_FILE_ADDRESS={address of the file you want to upload to your container, ex: c:\testuploadgile.log}
    ```

10. Run the sample.

    ```
    go run app.go
    ```
    
## More information

If you don't have a Microsoft Azure subscription you can get a FREE trial account [here](http://go.microsoft.com/fwlink/?LinkId=330212).

---

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/). For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.
