# AzureCN CDN

## Commands

### Upload Https Certficate

```shell
export AZURE_CN_CDN_KEY_ID={AzureCN CDN KeyId}
export AZURE_CN_CDN_KEY_VALUE={AzureCN CDN KeyValue}
export AZURE_CN_SUBSCRIPTION_ID={AzureCN SubscriptionId}
azure-cn-cdn-cmd upload-https-certificate {Cert Name} {Public Cert Path} {PrivateKey Path}
```