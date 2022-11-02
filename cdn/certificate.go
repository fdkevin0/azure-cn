package cdn

import (
	"encoding/json"
	"net/http"
)

// Upload HTTPS certificate
//
// https://docs.azure.cn/en-us/cdn/cdn-upload-https-certificate
func (c *Client) UploadHttpsCertificate(name, publicCertificate, privateKey string) (resp *http.Response, result *UploadHttpsCertificateResponse, err error) {
	postBody, _ := json.Marshal(&UploadHttpsCertificatePostBody{
		CertificateName:   name,
		PublicCertificate: publicCertificate,
		PrivateKey:        privateKey,
		Format:            "Pem",
	})
	// fmt.Println(string(postBody))
	resp, err = c.Request("POST", c.MakeRequestUrl("/https/certificates?apiVersion=1.0", nil), postBody, &result)
	return resp, result, err
}

type UploadHttpsCertificatePostBody struct {
	CertificateName   string
	PublicCertificate string
	PrivateKey        string
	Format            string
}

type UploadHttpsCertificateResponse struct {
	CertificateID           string
	CertificateName         string
	SubscriptionID          string
	ClientCertificateID     string
	Format                  string
	State                   string
	Issuers                 []string
	Subjects                []string
	SubjectAlternativeNames interface{}
	Thumbprint              string
	SerialNumber            string
	ValidFrom               string
	ValidTo                 string
}
