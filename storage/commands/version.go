package commands

import (
	"github.com/spf13/cobra"
	"fmt"
	ipfs "github.com/kidinamoto01/go-ipfs-api"
)


const (
	Path = "ipfs"
	shellUrl     = "localhost:5001"
)

var VersionCmd = &cobra.Command{
	Use:   "version ",
	Short: "this will return the IPFS version",
	RunE:  runVersion,
}

func runVersion(cmd *cobra.Command, args []string) error {
	//root := viper.GetString(cli.HomeFlag)
	//resetRoot(root, false)
	//return nil

	s := ipfs.NewShell(shellUrl)

	version, commit,err := s.Version()
	if err!= nil{
		panic(err)
	}else{
		// "QmUfZ9rAdhV5ioBzXKdUTh2ZNsz9bzbkaLVyQ8uc8pj21F")
		fmt.Println("get version",version,commit)

	}
	return nil
}

//func GetVersion(){
//
//	s := ipfs.NewShell(shellUrl)
//
//	version, commit,err := s.Version()
//	if err!= nil{
//		panic(err)
//	}else{
//		// "QmUfZ9rAdhV5ioBzXKdUTh2ZNsz9bzbkaLVyQ8uc8pj21F")
//		fmt.Println("get version",version,commit)
//
//	}
//}