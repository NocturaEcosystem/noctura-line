package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
)

var version = &cobra.Command{
	Use: "version",
	Short: "Used to get the version of Notcura",
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("noctura-line v0.0")
	},
}

func init() {
	rootcmd.AddCommand(version)
}