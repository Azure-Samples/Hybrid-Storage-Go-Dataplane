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

1.  Install Go SDK and its dependencies, [install Go SDK](https://github.com/azure/azure-sdk-for-go) 

1.  Install Go SDK dataplane, [install Go SDK Dataplane](https://github.com/Azure/azure-storage-blob-go/) 

1. Clone the sample project repository to your `GOPATH` location. You can add a new path to `GOPATH` location by adding an existing folder path to the `GOPATH` user environment variable. 
    - Create a `src` folder inside this new `GOPATH` folder and `cd` into the `src` folder.
    ```
    mkdir src
    cd src
    ```
    - Clone the sample project repository into your `src` folder.
    ```
    git clone https://github.com/Azure-Samples/Hybrid-Compute-Go-ManagedDisks.git
    ```

1.  Move the Hybrid-Storage-Go-Dataplane folder to your `$GOPATH/src` folder.

1.  Open a Powershell or Bash shell in $GOPATH/src/Hybrid-Storage-Go-Dataplane and enter the following commands:
    ```
    go mod init Hybrid-Storage-Go-Dataplane
    go mod tidy
    go get github.com/Azure/azure-storage-blob-go/azblob@v0.10.0
    ```

    NOTE: The azblob@v0.10.0 version is required for AzureStack.

1.  Create a [service principal using a certificate](https://docs.microsoft.com/en-us/azure/azure-stack/azure-stack-create-service-principals#create-a-service-principal-using-a-certificate) to work against AzureStack. Make sure your service principal has [contributor/owner role](https://docs.microsoft.com/en-us/azure/azure-stack/azure-stack-create-service-principals#assign-role-to-service-principal) on your subscription.

1.  Create a file that you wish to upload to your container, ex: c:\testuploadfile.log

1.  Fill in and export these environment variables into your current shell. 
    ```
    export AZURE_ARM_ENDPOINT={your AzureStack Resource Manager Endpoint}
    export AZURE_TENANT_ID={your tenant id}
    export AZURE_SP_CERT_ID={your service principal certificate client id}
    export AZURE_SP_CERT_PASS={password for your service principal certificate .pfx file}
    export AZURE_SP_CERT_PATH={path to your password protected service principal certificate .pfx file}
    export AZURE_SUBSCRIPTION_ID={your subscription id}
    export AZURE_LOCATION={your resource location}
    export AZURE_SAMPLE_FILE_PATH={address of the file you want to upload to your container, ex: c:/testuploadfile.txt}
    ```

1. Run the sample.
    ```
    go run app.go
    ```
    
## More information

If you don't have a Microsoft Azure subscription you can get a FREE trial account [here](http://go.microsoft.com/fwlink/?LinkId=330212).

---

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/). For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.
