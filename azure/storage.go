package azure

import (
	"context"
	"fmt"
	"iot/log"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/Azure/azure-storage-blob-go/azblob"
	_ "github.com/joho/godotenv/autoload"
)

func handleErrors(err error) {
	if err != nil {
		if serr, ok := err.(azblob.StorageError); ok { // This error is a Service-specific
			switch serr.ServiceCode() { // Compare serviceCode to ServiceCodeXxx constants
			case azblob.ServiceCodeContainerAlreadyExists:
				// fmt.Println("컨테이너 존재합니다.")
				log.Error.Println("컨테이너 존재하여 생성 생략합니다.")
				return
			}
		}
	}
}

func Storage(containerName string, guid string, fileName string) {
	// From the Azure portal, get your storage account name and key and set environment variables.
	accountName, accountKey := os.Getenv("AZURE_STORAGE_ACCOUNT"), os.Getenv("AZURE_STORAGE_ACCESS_KEY")
	if len(accountName) == 0 || len(accountKey) == 0 {
		log.Error.Println("Either the AZURE_STORAGE_ACCOUNT or AZURE_STORAGE_ACCESS_KEY environment variable is not set")
	}

	// Create a default request pipeline using your storage account name and account key.
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		log.Error.Println("Invalid credentials with error: " + err.Error())
	}
	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	// From the Azure portal, get your storage account blob service URL endpoint.
	URL, _ := url.Parse(
		fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName))

	// Create a ContainerURL object that wraps the container URL and a request
	// pipeline to make requests.
	containerURL := azblob.NewContainerURL(*URL, p)

	// Create the container
	// fmt.Printf("Creating a container named %s\n", containerName)
	ctx := context.Background() // This example uses a never-expiring context
	_, err = containerURL.Create(ctx, azblob.Metadata{}, azblob.PublicAccessNone)
	handleErrors(err)
	re := regexp.MustCompile(`(\d{2})(\d{2})(\d{2})`)
	date := re.ReplaceAllString(strings.Split(fileName, "_")[0], "20$1-$2-$3")
	uploadFile := fmt.Sprintf("./files/%s", fileName)
	newFileName := fmt.Sprintf("%s/%s/%s_%s", guid, date, fileName)
	// fmt.Println(newFileName,"newFileName")
	// Here's how to upload a blob.
	blobURL := containerURL.NewBlockBlobURL(newFileName)
	file, err := os.Open(uploadFile)
	handleErrors(err)

	// fmt.Printf("ContentType: %s\n", fileName)
	var ContentType azblob.BlobHTTPHeaders
	ContentType.ContentType = "video/mp4"

	// fmt.Printf("Uploading the file with blob name: %s\n", fileName)
	_, err = azblob.UploadFileToBlockBlob(ctx, file, blobURL, azblob.UploadToBlockBlobOptions{
		BlobHTTPHeaders: ContentType,
		BlockSize:       4 * 1024 * 1024,
		Parallelism:     16,
	})
	handleErrors(err)

	log.Trace.Println("성공적으로 업로드 되었습니다.")

	// 파일 삭제
	if err := os.Remove(uploadFile); err != nil {
		log.Error.Println(err)
	}
}
