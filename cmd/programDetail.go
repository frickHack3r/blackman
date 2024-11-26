package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const DETAILS = `################################################################################
#                ____                     _ __                                 #
#     ___  __ __/ / /__ ___ ______ ______(_) /___ __                           #
#    / _ \/ // / / (_-</ -_) __/ // / __/ / __/ // /                           #
#   /_//_/\_,_/_/_/___/\__/\__/\_,_/_/ /_/\__/\_, /                            #
#                                            /___/ team                        #
#                                                                              #
# Version                                                                      #
#   v1.0.0                                                                     #
#                                                                              #
# DESCRIPTION                                                                  #
# Download and compile packages as Emerge does                                 #
#                                                                              #
# AUTHOR                                                                       #
# nrz@nullsecurity.net                                                         #
# noptrix@nullsecurity.net                                                     #
# Frick Hack3r <frickhack3r@outlook.com>                                       #
#                                                                              #
################################################################################`

var detailsCmd = &cobra.Command{
	Use:   "details",
	Short: "Print the details of Blackman",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", DETAILS)
	},
}

func init() {
	rootCmd.AddCommand(detailsCmd)
}
