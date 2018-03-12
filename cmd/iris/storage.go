package main

import (
	//"fmt"

	"github.com/spf13/cobra"
	"github.com/irisnet/iris-hub/storage/commands"

)


// StorageCmd - The service of ipfs
var StorageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Show storage info",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}


func prepareStorageCommands() {

	StorageCmd.AddCommand(
		commands.VersionCmd,
		)
}


