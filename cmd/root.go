package cmd

import (
	"fmt"
	"log"

	"github.com/arfadmuzali/restui/cmd/restui"
	"github.com/arfadmuzali/restui/internal/help"
	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

var Version = "dev"

var guide bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "restui",
	Version: Version,
	Short:   "RESTUI, API Client in your terminal",
	RunE: func(cmd *cobra.Command, args []string) error {
		if guide {
			out, err := glamour.Render(help.Guide, "dark")
			fmt.Print(out)
			return err
		}
		return restui.Execute()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.restui.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolVarP(&guide, "guide", "g", false, "List of shortcut and tips and trick.")
}
