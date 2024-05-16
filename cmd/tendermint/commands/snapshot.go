package commands

import (
	"strconv"

	"github.com/anchorageoss/sei-tendermint/config"
	"github.com/anchorageoss/sei-tendermint/internal/dbsync"
	"github.com/spf13/cobra"
)

func MakeSnapshotCommand(confGetter func(*cobra.Command) (*config.Config, error)) *cobra.Command {
	return &cobra.Command{
		Use:   "snapshot [height]",
		Short: "Take DBSync snapshot for given height",
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := confGetter(cmd)
			if err != nil {
				return err
			}
			height, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			return dbsync.Snapshot(height, *conf.DBSync, conf.BaseConfig)
		},
	}
}
