package commands

import (
	"github.com/spf13/cobra"
	"fmt"
	ipfs "github.com/kidinamoto01/go-ipfs-api"
	flag "github.com/spf13/pflag"
	"github.com/pkg/errors"
	//"github.com/spf13/viper"
	//"strings"
	"github.com/spf13/viper"
	storage "github.com/irisnet/iris-hub/storage"
)


var (UploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "this will add file to IPFS",
	RunE:  runUpload,
}
	FlagUploadFilePath   = "upload-path"
)


var (DownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "this will download file from IPFS",
	RunE:  runDownload,
}
	FlagDownloadFilePath   = "download-path"
)

func init() {

	fsService := flag.NewFlagSet("", flag.ContinueOnError)
	fsService.String(FlagUploadFilePath, "", "upload file path")
	fsService.String(FlagDownloadFilePath, "", "download file path")
	UploadCmd.Flags().AddFlagSet(fsService)
	DownloadCmd.Flags().AddFlagSet(fsService)
}


func runUpload(cmd *cobra.Command, args []string) error {

	path := viper.GetString(FlagUploadFilePath)
	fmt.Println(path)

	if path != "" {
		s := ipfs.NewShell(shellUrl)

		mhash, err := s.AddFile(path)
		if err != nil{
			return errors.New("cannot find the file")
		}else {
			// "QmUfZ9rAdhV5ioBzXKdUTh2ZNsz9bzbkaLVyQ8uc8pj21F")
			fmt.Println("get output",mhash)

		}
	}


	return nil
}



func runDownload(cmd *cobra.Command, args []string) error {

	path := viper.GetString(FlagDownloadFilePath)
	fmt.Println(path)

	if path != "" {

		if len(args) != 1 {
			return storage.ErrMissingInput()
			//return fmt.Errorf("`download` takes one argument, a basecoin account address. Generate one using `basecli keys new mykey`")
		}

		fileHash := args[0]

		s := ipfs.NewShell(shellUrl)

		//QmaWsjUEsUV8bLCfkzyecyC3hnubeyZ9PVhNLiZ5JAvAfr
		err := s.Get(fileHash,path)
		if err != nil{
			return errors.New("cannot find the file")
		}else {

			fmt.Println("download success")

		}
	}else{
		return storage.ErrMissingInput()
	}


	return nil
}