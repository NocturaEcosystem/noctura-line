package cmds

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var verbose bool
var rootcmd = &cobra.Command{
	Use: "noctura",
	Short: "CLI tool for Noctura ecosystem",
	Long: "CLI tool used for Noctura ecosystems. You can spawn desktop shells, utility, etc. with this CLI",
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("This is the cli tool for noctura devices... Use '--help' to get the list of available commands")
	},
}

func Execute() {
	if err := rootcmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootcmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
}
