package status

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/weaveworks/weave-gitops/pkg/cmdimpl"
)

var Cmd = &cobra.Command{
	Use:     "status [subcommands]",
	Short:   "status of a resource",
	Args:    cobra.MinimumNArgs(1),
	Example: "wego status application podinfo",
}

var ApplicationCmd = &cobra.Command{
	Use:   "application [name]",
	Short: "status of an application resource",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmdimpl.Status(args, params)
	},
}

var params cmdimpl.AddParamSet

func init() {

	Cmd.Flags().StringVar(&params.Name, "name", "", "Name of remote git repository")
	Cmd.Flags().StringVar(&params.Url, "url", "", "URL of remote git repository")
	Cmd.Flags().StringVar(&params.Path, "path", "./", "Path of files within git repository")
	Cmd.Flags().StringVar(&params.Branch, "branch", "main", "Branch to watch within git repository")
	Cmd.Flags().StringVar(&params.PrivateKey, "private-key", filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa"), "Private key that provides access to git repository")

	Cmd.MarkFlagRequired("name")

	Cmd.AddCommand(ApplicationCmd)
}