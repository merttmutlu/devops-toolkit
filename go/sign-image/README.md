# ecr-image-signer

This script allows you to sign docker images on ECR with cosign.

## Prerequisites

Before you begin, ensure you have the following installed:

-   Go programming language: [Download Go](https://golang.org/dl/)
-   AWS CLI: [Install AWS CLI](https://aws.amazon.com/cli/)

## Installation

To install the required Go packages, run the following commands in your terminal or command prompt:
`go get github.com/aws/aws-sdk-go-v2`
`go get github.com/aws/aws-sdk-go-v2/config`
`go get github.com/aws/aws-sdk-go-v2/service/ecr` 

## Configuration

### Set Registry URL

In your Go code, set the registry URL by replacing `REPLACE-REGISTRY-NUMBER.dkr.ecr.eu-central-1.amazonaws.com` with your actual AWS ECR registry URL.

`registryURL := "REPLACE-REGISTRY-NUMBER.dkr.ecr.eu-central-1.amazonaws.com"` 

### Configure AWS CLI Authentication

Configure AWS CLI authentication using one of the following methods:

#### Method 1: Using `aws configure`

Run the following command and provide your AWS access key ID, secret access key, region, and preferred output format when prompted.
`aws configure` 

#### Method 2: Using Environment Variables

Set the following environment variables with your AWS credentials:

-   `AWS_ACCESS_KEY_ID`: Your AWS access key ID.
-   `AWS_SECRET_ACCESS_KEY`: Your AWS secret access key.
-   `AWS_DEFAULT_REGION`: Your preferred AWS region.

Example:
`export AWS_ACCESS_KEY_ID=YOUR_ACCESS_KEY_ID`
`export AWS_SECRET_ACCESS_KEY=YOUR_SECRET_ACCESS_KEY`
`export AWS_DEFAULT_REGION=eu-central-1` 

## Usage

`go run .`

## License

This project is licensed under the [License Name] License - see the [LICENSE.md](https://chat.openai.com/c/LICENSE.md) file for details.
