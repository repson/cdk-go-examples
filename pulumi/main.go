package main

import (
	do "github.com/pulumi/pulumi-docker/sdk/v3/go/docker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		image, err := do.NewRemoteImage(ctx, "ubuntu", &docker.RemoteImageArgs{
			Name: pulumi.String("ubuntu:precise"),
		})
		if err != nil {
			return err
		}

		container, err := do.NewContainer(ctx, "ubuntu", &docker.ContainerArgs{
			Image: image.Latest(),
		})
		if err != nil {
			return err
		}

		return nil
	})
}
