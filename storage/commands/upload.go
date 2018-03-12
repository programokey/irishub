package commands

import (
	"github.com/spf13/cobra"
	"fmt"
	ipfs "github.com/kidinamoto01/go-ipfs-api"
	"flag"
	"github.com/pkg/errors"
	//"github.com/spf13/viper"
	"strings"
)


var (UploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "this will add file to IPFS",
	RunE:  runUpload,
}
	FlagFilePath   = "path"
)

func init() {

	fsService := flag.NewFlagSet("", flag.ContinueOnError)
	fsService.String(FlagFilePath, "", "upload file path")
}


func runUpload(cmd *cobra.Command, args []string) error {

	//path := viper.GetString(FlagFilePath)
	if len(args) != 1 {
		return fmt.Errorf("`init` takes one argument, a basecoin account address. Generate one using `basecli keys new mykey`")
	}
	path := ""

	input := args[0]
	kv := strings.Split(input, "=")
	if len(kv) == 2 {
		path = kv[1]
	} else if len(kv) == 1 && kv[0] != "" {
		kv = strings.Split(input, " ")
		if len(kv) == 2{
			path = kv[1]
		}
	}



	if path != "" {
		s := ipfs.NewShell(shellUrl)

		mhash, err := s.AddFile(path)
		if err!= nil{
			return errors.New("cannot find the file")
		}else{
			// "QmUfZ9rAdhV5ioBzXKdUTh2ZNsz9bzbkaLVyQ8uc8pj21F")
			fmt.Println("get output",mhash)

		}
	}


	return nil
}
