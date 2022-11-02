package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fdkevin0/azure-cn/cdn"
)

func main() {
	cdnClient := cdn.NewClient(
		os.Getenv("AZURE_CN_CDN_KEY_ID"),
		os.Getenv("AZURE_CN_CDN_KEY_VALUE"),
		os.Getenv("AZURE_CN_SUBSCRIPTION_ID"),
	)

	if len(os.Args) == 1 {
		log.Fatal("Please input command")
	}

	switch os.Args[1] {
	case "list-endpoints":
		resp, result, err := cdnClient.ListEndpoints()
		if err != nil {
			log.Fatal(err)
		}
		PrintJson(resp, result)
	case "upload-https-certificate":
		var (
			pubCert []byte
			privKey []byte
			err     error
			resp    *http.Response
			result  *cdn.UploadHttpsCertificateResponse
		)
		if pubCert, err = os.ReadFile(os.Args[3]); err != nil {
			log.Fatal(err)
		}
		if privKey, err = os.ReadFile(os.Args[4]); err != nil {
			log.Fatal(err)
		}
		if resp, result, err = cdnClient.UploadHttpsCertificate(os.Args[2], string(pubCert), string(privKey)); err != nil {
			log.Println("X-Correlation-Id:", resp.Header.Get("X-Correlation-Id"))
			log.Fatalln(err)
		}
		PrintJson(result)
	}
}

func PrintJson(a ...any) {
	for _, v := range a {
		b, _ := json.MarshalIndent(v, "", "  ")
		fmt.Println(string(b))
	}
}
