/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package image

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var ImageCmd = &cobra.Command{
	Use:   "image",
	Short: "docker image commands like list,delete...",
	Long:  `Provides all the commands related to docker image.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// register all the underlying image sub-commands
func addSubCommands() {
	ImageCmd.AddCommand(listCmd)
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubCommands()
}
