package image

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm IMAGE_ID",
	Short: "remove the image for the given image ID",
	Long:  `different flags for different actions while removing the image`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		imageID := args[0]
		if err := RemoveDockerImage(imageID); err != nil {
			fmt.Println("Error Encountered: ", err)
		}
	},
}

func RemoveDockerImage(imageID string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("failed to create Docker client: %w", err)
	}
	defer cli.Close()

	ctx := context.Background()

	removedImages, err := cli.ImageRemove(ctx, imageID, image.RemoveOptions{
		Force:         false,
		PruneChildren: true,
	})
	if err != nil {
		return fmt.Errorf("failed to remove image: %w", err)
	}

	for _, removed := range removedImages {
		if removed.Deleted != "" {
			fmt.Printf("Deleted: %s\n", removed)
		}
		if removed.Untagged != "" {
			fmt.Printf("Untagged: %s\n", removed.Untagged)
		}
	}

	return nil
}

func init() {
	// your flages here if you want any
}
