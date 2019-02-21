package main

import (
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	pepe "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"golang.org/x/net/context"
)

func buildContainer(image string, nodes int) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.35"))
	// cli, err := client.NewClientWithOpts(client.FromEnv)
	fmt.Print(cli.ClientVersion())
	if err != nil {
		panic(err)
	}

	conf := pepe.Config{
		Image: image,
		Tty:   true,
	}

	bindings := nat.PortMap{}

	emptypaths := []string{}

	restartpol := pepe.RestartPolicy{
		Name: "no",
	}

	for i := start; i < nodes+start; i++ {
		fmt.Printf("Building Node %d...\n", i)

		vlan := fmt.Sprintf("wb_vlan%d", i)

		hostConf := pepe.HostConfig{
			NetworkMode:     pepe.NetworkMode(vlan),
			PublishAllPorts: false,
			PortBindings:    bindings,
			ReadonlyPaths:   emptypaths,
			MaskedPaths:     emptypaths,
			DNS:             emptypaths,
			DNSOptions:      emptypaths,
			DNSSearch:       emptypaths,
			RestartPolicy:   restartpol,
		}

		resp, err := cli.ContainerCreate(ctx, &conf, &hostConf, nil, fmt.Sprintf("whiteblock-node%d", i))
		if err != nil {
			panic(err)
		}

		statusCh, errCh := cli.ContainerWait(ctx, resp.ID, pepe.WaitConditionNotRunning)
		select {
		case err := <-errCh:
			if err != nil {
				panic(err)
			}
		case <-statusCh:
		}

		out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
		if err != nil {
			panic(err)
		}

		io.Copy(os.Stdout, out)
	}
}

func startContainer(start, nodes int) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.35"))
	// cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	for i := start; i < nodes+start; i++ {
		fmt.Printf("Starting Node %d...\n", i)
		if err := cli.ContainerStart(ctx, fmt.Sprintf("whiteblock-node%d", i), types.ContainerStartOptions{}); err != nil {
			panic(err)
		}
	}
}

func joinNetwork(start, nodes int) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.35"))
	// cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	for i := start; i < nodes+start; i++ {

		var nodeName = fmt.Sprintf("whiteblock-node%d", i)
		var networkName = fmt.Sprintf("wb_vlan%d", i)

		epSettings := network.EndpointSettings{}

		cli.NetworkConnect(ctx, networkName, nodeName, &epSettings)
	}
}

func deleteContainer(start, nodes int) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.35"))
	// cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Removing Container %d...\n", nodes)

	for i := start; i < nodes+start; i++ {
		cli.ContainerRemove(ctx, fmt.Sprintf("whiteblock-node%d", i), types.ContainerRemoveOptions{
			Force: true,
		})
		if err != nil {
			panic(err)
		}
	}
}
