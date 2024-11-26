/*
Copyright © 2024 Frick Hack3r <frickhack3r@outlook.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "blackman",
	Short: "A brief description of your application",
	Long:  `Blackman is a tool for managing Blackarch repositories. It allows you to search, download, and install packages from the Blackarch repositories. Blackarch is an Arch Linux-based penetration testing distribution that provides a large collection of security and penetration testing tools.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
