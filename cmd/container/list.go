/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package container

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list the running containers in the host machine",
	Long:  `Provide different flags for more control over the command`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := listContainers(&listAll); err != nil {
			fmt.Println("Error Encountered: ", err)
		}
	},
}

var listAll bool

// give me the method that list all the container in the host machine
func listContainers(listAll *bool) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("failed to create Docker client: %w", err)
	}

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{
		All: *listAll,
	})

	if err != nil {
		return fmt.Errorf("failed to list containers: %w", err)
	}

	if len(containers) == 0 {
		fmt.Println("No containers found.")
		return nil
	}
	count := 1
	for _, container := range containers {
		fmt.Printf("No: %d\nID: %s\tImage: %s\tStatus: %s\tName: %s\n", count, container.ID[:10], container.Image, container.Status, container.Names[0])
		count++
	}

	return nil
}

// register the flags to this command
func init() {
	listCmd.Flags().BoolVarP(&listAll, "all", "a", false, "list all containers in the host machine")
}
