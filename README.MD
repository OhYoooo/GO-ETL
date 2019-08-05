# Common-Etl

## Introduction
A common ETL sample to fetch data and ack from google cloud pubsub, do some transaction then save to google datastore entities. This project need to be using Go 11 or above, as it is using `go mod` as package management tool.

## Init
- put the project **out of** `$GOPATH/src`
- run `go mod init common-etl`
- run `go build`

## Run
- run `go run main.go`

## Env Variables
examples:
```bash
export GCLOUD_PROJECT_ID="sample"
export GOOGLE_APPLICATION_CREDENTIALS="/usr/local/google/keys"
export LOG_LOCATION="/log/location"
export ENVIRONMENT="test"
export PORT="8080"
export SUBSCRIPTION_NAME="sample_name"
```
