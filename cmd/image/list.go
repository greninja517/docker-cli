package image

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list the docker images",
	Long:  `used for displaying the docker images`,

	// run used for defining the action of a subcommand
	Run: func(cmd *cobra.Command, args []string) {
		listImages()
	},
}

var number int

func listImages() {
	cmd := exec.Command("docker", "images", "--format", "{{.Repository}}\t{{.ID}}\t{{.Size}}")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running docker images:", err)
		return
	}

	// Print table headers
	fmt.Printf("%-30s %-20s %-10s\n", "REPOSITORY", "IMAGE ID", "SIZE")
	fmt.Println(strings.Repeat("-", 60))

	// Print actual command output
	lines := strings.Split(string(out), "\n")
	line_count := 0 // for controlling the no. of images to show
	for _, line := range lines {

		if line != "" {
			cols := strings.Split(line, "\t")
			if len(cols) == 3 {
				fmt.Printf("%-30s %-20s %-10s\n", cols[0], cols[1], cols[2])
				line_count++
			}
		}
		if (line_count >= number) && (number > 0) {
			break
		}
	}
}

// register the flags to the list sub-command
func init() {
	listCmd.Flags().IntVarP(&number, "number", "n", 0, "Number of images to show")
}
