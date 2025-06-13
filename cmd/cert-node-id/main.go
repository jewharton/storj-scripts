package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zeebo/errs"

	"storj.io/common/identity"
)

var cmd = &cobra.Command{
	Use:   "cert-node-id <cert-path>",
	Short: "Output the node ID of an identity cert",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		nodeID, err := identity.NodeIDFromCertPath(args[0])
		if err != nil {
			return errs.Wrap(err)
		}
		fmt.Println(nodeID)
		return nil
	},
}

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}
