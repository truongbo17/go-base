package cli

import (
	"github.com/spf13/cobra"
	"go-base/migrations"
)

var MigrateCmd = &cobra.Command{
	Use:     "migrate",
	Short:   "Run database migrations",
	Example: "ggb migrate",
	Run: func(cmd *cobra.Command, args []string) {
		migrateCmd(args)
	},
}

func migrateCmd(args []string) {
	migrations.Migrate()
}
