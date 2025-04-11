/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package container

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "removes the container for the given container ID",
	Long:  `Provides different flags for tweaking the behaviour of the rm command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rm called")
	},
}

func init() {

}
