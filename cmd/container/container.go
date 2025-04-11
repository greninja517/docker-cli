/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package container

import (
	"github.com/spf13/cobra"
)

// containerCmd represents the container command
var ContainerCmd = &cobra.Command{
	Use:   "container",
	Short: "Handles all the commands related to containers",
	Long:  `Provides commands like list, delete, stop, etc. for the containers`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func addSubCommands() {
	ContainerCmd.AddCommand(listCmd)
	ContainerCmd.AddCommand((rmCmd))
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// containerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// containerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubCommands()
}
