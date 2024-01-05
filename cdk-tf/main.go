package main

import (
	"github.com/aws/constructs-go/constructs/v10/constructs"
	"github.com/aws/jsii-runtime-go"

	"github.com/cdktf/cdktf-provider-docker-go/docker/v3/container"
	dockerprovider "github.com/cdktf/cdktf-provider-docker-go/docker/v3/provider"
	"github.com/hashicorp/terraform-cdk-go/cdktf/terraformStack"
)

func NewMyStack(scope constructs.Construct, id string) terraformStack.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	dockerprovider.NewProvider(stack, jsii.String("docker"), &dockerprovider.ProviderConfig{})

	dockerImage := image.NewImage(stack, jsii.String("docker_image"), &image.ImageConfig{
		Name:        jsii.String("nginx:latest"),
		KeepLocally: jsii.Bool(false),
	})

	container.NewContainer(stack, jsii.String("docker_container"), &container.ContainerConfig{
		Image: dockerImage.Name(),
		Name:  jsii.String("nginx"),
		Ports: &[]*container.ContainerPorts{{
			Internal: jsii.Number(80), External: jsii.Number(8000),
		}},
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "cdktf")

	app.Synth()
}
