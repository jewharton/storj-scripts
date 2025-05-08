package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/zeebo/errs"

	"storj.io/common/identity"
	"storj.io/common/identity/testidentity"
	"storj.io/common/storj"
)

var (
	outputDir string
	index     int
	signed    bool
)

var cmd = &cobra.Command{
	Use:   "pregen-identity",
	Short: "Exports pregenerated identities",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		gen := testidentity.PregeneratedIdentity
		if signed {
			gen = testidentity.PregeneratedSignedIdentity
		}

		pregenIdent, err := gen(index, storj.LatestIDVersion())
		if err != nil {
			return errs.Wrap(err)
		}

		identConfig := identity.Config{
			CertPath: filepath.Join(outputDir, "identity.cert"),
			KeyPath:  filepath.Join(outputDir, "identity.key"),
		}

		if err = identConfig.Save(pregenIdent); err != nil {
			return errs.Wrap(err)
		}

		absOutputDir, err := filepath.Abs(outputDir)
		if err != nil {
			return errs.Wrap(err)
		}

		nodeID := pregenIdent.ID.String()
		err = os.WriteFile(filepath.Join(outputDir, "node_id.txt"), []byte(nodeID), 0644)
		if err != nil {
			return errs.Wrap(err)
		}

		fmt.Printf("Identity (node ID: %s) output to %q\n", nodeID, absOutputDir)

		return nil
	},
}

func main() {
	cmd.Flags().StringVar(&outputDir, "output-dir", "", "The directory into which the pregenerated identity is output")
	cmd.Flags().IntVar(&index, "index", 1, "The index of the pregenerated identity")
	cmd.Flags().BoolVar(&signed, "signed", false, "Whether the pregenerated identity should be signed")

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}
