/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package container

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm [CONTAINER ID]",
	Short: "removes the container for the given container ID",
	Long:  `Provides different flags for tweaking the behaviour of the rm command`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		containerID := args[0]
		if err := removeContainer(containerID); err != nil {
			fmt.Println("Error Encountered: ", err)
		}
	},
}

func removeContainer(prefix string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("failed to create Docker client: %w", err)
	}
	defer cli.Close()

	ctx := context.Background()

	// Step 1: List all containers (running + stopped)
	containers, err := cli.ContainerList(ctx, container.ListOptions{
		All: true},
	)

	if err != nil {
		return fmt.Errorf("failed to list containers: %w", err)
	}

	// prefix matching with the container_ID
	var matched []string
	for _, container := range containers {
		if strings.HasPrefix(container.ID, prefix) {
			matched = append(matched, container.ID)
		}
	}

	// handle the results of prefix matching
	switch len(matched) {
	case 0:
		return fmt.Errorf("no container found with id: %s", prefix)
	case 1:
		// Proceed with deletion
		ctrName, _ := getContainerNameById(cli, ctx, matched[0])
		err := cli.ContainerRemove(ctx, matched[0], container.RemoveOptions{})
		if err != nil {
			return fmt.Errorf("failed to remove container %s: %w", ctrName, err)
		}
		fmt.Printf("Removed container: %s\n", ctrName)
		return nil

	default:
		return fmt.Errorf("ambiguous container id: %s (matches %d containers)", prefix, len(matched))
	}
}

func getContainerNameById(cli *client.Client, ctx context.Context, id string) (string, error) {
	containerJSON, err := cli.ContainerInspect(ctx, id)
	if err != nil || containerJSON.Name == "" {
		return "", errors.New("failed to get container name")
	}
	return strings.TrimPrefix(containerJSON.Name, "/"), nil
}

func init() {
	// your flags here if any
}
