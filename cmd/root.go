package cmd

import (
	"os"

	"github.com/paddymorgan84/kubectl-pdborked/pdbs"
	"github.com/paddymorgan84/kubectl-pdborked/ui"
	"github.com/spf13/cobra"
)

var namespace string = ""

var pdborkedCmd = &cobra.Command{
	Use:   "pdborked",
	Short: "A kubectl plugin extension to identify PDBs with no allowed disruptions",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := pdbs.GetBorkedPdbs(namespace, new(ui.TableRenderer))

		if err != nil {
			return err
		}

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := pdborkedCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	pdborkedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	pdborkedCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "The namespace to run against")
}
