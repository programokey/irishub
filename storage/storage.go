package storage

import (
"fmt"

"github.com/spf13/cobra"
)

// Version - Iris Version
const Path = "ipfs"

// StorageCmd - The service of ipfs
var StorageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Show storage info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", Path)
	},
}
