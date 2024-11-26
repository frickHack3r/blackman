/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"log"
	"os"

	"github.com/spf13/cobra"
)

// url blackarch repository
const REMOTE_REPO = "https://github.com/BlackArch/blackarch.git"

// local base directory for blackman
var LOCAL_BLACKMAN string

func getLocalBlackman() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return dirname + "/.blackman"
}

func installPackage(packageName string) {
	fmt.Println("Installing package: " + packageName)
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a package",
	Long:  `Install a package from the Blackarch repository`,
	Run: func(cmd *cobra.Command, args []string) {
		LOCAL_BLACKMAN = getLocalBlackman()
		fmt.Println(LOCAL_BLACKMAN)
		fmt.Println(args)
		fmt.Println("install called")
		for _, packageName := range args {
			installPackage(packageName)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	installCmd.Flags().BoolP("install", "i", false, "Install a package")
}
