package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
)

func main() {
	client, err := createECRClient()
	if err != nil {
		log.Fatal("Error creating ECR client:", err)
	}

	registryURL := "REPLACE-REGISTRY-NUMBER.dkr.ecr.eu-central-1.amazonaws.com" // Replace this with your AWS registry URL
	repositoryTags, err := listRepositoryTags(client)
	if err != nil {
		log.Fatal("Error listing repository tags:", err)
	}

	signImages(repositoryTags, registryURL)
}

func createECRClient() (*ecr.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	return ecr.NewFromConfig(cfg), nil
}

func listRepositoryTags(client *ecr.Client) ([]string, error) {
	input := &ecr.DescribeRepositoriesInput{}
	resp, err := client.DescribeRepositories(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	var repositoryTags []string
	for _, repo := range resp.Repositories {
		tags, err := listTagsForRepository(*repo.RepositoryName, client)
		if err != nil {
			log.Printf("Error listing tags for repository %s: %v\n", *repo.RepositoryName, err)
			continue
		}

		for _, tag := range tags {
			repositoryTag := *repo.RepositoryName + ":" + tag
			repositoryTags = append(repositoryTags, repositoryTag)
		}
	}

	return repositoryTags, nil
}

func listTagsForRepository(repoName string, client *ecr.Client) ([]string, error) {
	input := &ecr.DescribeImagesInput{
		RepositoryName: &repoName,
	}

	resp, err := client.DescribeImages(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	var tags []string
	for _, image := range resp.ImageDetails {
		for _, tag := range image.ImageTags {
			tags = append(tags, tag) // Append the tag directly, without using *
		}
	}

	return tags, nil
}

func getRepository(repoTag string) string {
	parts := strings.Split(repoTag, ":")
	if len(parts) > 1 {
		return parts[0]
	}
	return ""
}

func getTag(repoTag string) string {
	parts := strings.Split(repoTag, ":")
	if len(parts) > 1 {
		return parts[1]
	}
	return ""
}

func signImages(repositoryTags []string, registryURL string) {
	for _, repoTag := range repositoryTags {
		image := fmt.Sprintf("%s/%s:%s", registryURL, getRepository(repoTag), getTag(repoTag))
		signImage(image)
	}
}

func signImage(image string) {
	cosignCommand := fmt.Sprintf("docker run --rm --name cosign -e COSIGN_PASSWORD=\"\" -v $(pwd):/cosign-keys/ -v ./config.json:/.docker/config.json bitnami/cosign:2.2.0 sign --key milvus-robotics-ecr.key -y %s", image)

	cmd := exec.Command("sh", "-c", cosignCommand)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("Error signing image %s: %v", image, err)
	} else {
		log.Printf("Successfully signed image: %s", image)
	}
}
