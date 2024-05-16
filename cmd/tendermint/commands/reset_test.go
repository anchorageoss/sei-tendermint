package commands

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	cfg "github.com/anchorageoss/sei-tendermint/config"
	"github.com/anchorageoss/sei-tendermint/libs/log"
	"github.com/anchorageoss/sei-tendermint/privval"
	"github.com/anchorageoss/sei-tendermint/types"
)

func Test_ResetAll(t *testing.T) {
	config := cfg.TestConfig()
	dir := t.TempDir()
	config.SetRoot(dir)
	logger := log.NewNopLogger()
	cfg.EnsureRoot(dir)
	require.NoError(t, initFilesWithConfig(context.Background(), config, logger, types.ABCIPubKeyTypeEd25519))
	pv, err := privval.LoadFilePV(config.PrivValidator.KeyFile(), config.PrivValidator.StateFile())
	require.NoError(t, err)
	pv.LastSignState.Height = 10
	require.NoError(t, pv.Save())
	require.NoError(t, ResetAll(config.DBDir(), config.PrivValidator.KeyFile(),
		config.PrivValidator.StateFile(), logger, types.ABCIPubKeyTypeEd25519))
	require.DirExists(t, config.DBDir())
	require.NoFileExists(t, filepath.Join(config.DBDir(), "block.db"))
	require.NoFileExists(t, filepath.Join(config.DBDir(), "state.db"))
	require.NoFileExists(t, filepath.Join(config.DBDir(), "evidence.db"))
	require.NoFileExists(t, filepath.Join(config.DBDir(), "tx_index.db"))
	require.FileExists(t, config.PrivValidator.StateFile())
	pv, err = privval.LoadFilePV(config.PrivValidator.KeyFile(), config.PrivValidator.StateFile())
	require.NoError(t, err)
	require.Equal(t, int64(0), pv.LastSignState.Height)
}

func Test_ResetState(t *testing.T) {
	config := cfg.TestConfig()
	dir := t.TempDir()
	config.SetRoot(dir)
	logger := log.NewNopLogger()
	cfg.EnsureRoot(dir)
	require.NoError(t, initFilesWithConfig(context.Background(), config, logger, types.ABCIPubKeyTypeEd25519))
	pv, err := privval.LoadFilePV(config.PrivValidator.KeyFile(), config.PrivValidator.StateFile())
	require.NoError(t, err)
	pv.LastSignState.Height = 10
	require.NoError(t, pv.Save())
	require.NoError(t, ResetState(config.DBDir(), logger))
	require.DirExists(t, config.DBDir())
	require.NoFileExists(t, filepath.Join(config.DBDir(), "block.db"))
	require.NoFileExists(t, filepath.Join(config.DBDir(), "state.db"))
	require.NoFileExists(t, filepath.Join(config.DBDir(), "evidence.db"))
	require.NoFileExists(t, filepath.Join(config.DBDir(), "tx_index.db"))
	require.FileExists(t, config.PrivValidator.StateFile())
	pv, err = privval.LoadFilePV(config.PrivValidator.KeyFile(), config.PrivValidator.StateFile())
	require.NoError(t, err)
	// private validator state should still be in tact.
	require.Equal(t, int64(10), pv.LastSignState.Height)
}

func Test_UnsafeResetAll(t *testing.T) {
	config := cfg.TestConfig()
	dir := t.TempDir()
	config.SetRoot(dir)
	logger := log.NewNopLogger()
	cfg.EnsureRoot(dir)
	require.NoError(t, initFilesWithConfig(context.Background(), config, logger, types.ABCIPubKeyTypeEd25519))
	pv, err := privval.LoadFilePV(config.PrivValidator.KeyFile(), config.PrivValidator.StateFile())
	require.NoError(t, err)
	pv.LastSignState.Height = 10
	require.NoError(t, pv.Save())
	require.NoError(t, ResetAll(config.DBDir(), config.PrivValidator.KeyFile(),
		config.PrivValidator.StateFile(), logger, types.ABCIPubKeyTypeEd25519))
	require.DirExists(t, config.DBDir())
	require.NoFileExists(t, filepath.Join(config.DBDir(), "block.db"))
	require.NoFileExists(t, filepath.Join(config.DBDir(), "state.db"))
	require.NoFileExists(t, filepath.Join(config.DBDir(), "evidence.db"))
	require.NoFileExists(t, filepath.Join(config.DBDir(), "tx_index.db"))
	require.FileExists(t, config.PrivValidator.StateFile())
	pv, err = privval.LoadFilePV(config.PrivValidator.KeyFile(), config.PrivValidator.StateFile())
	require.NoError(t, err)
	require.Equal(t, int64(0), pv.LastSignState.Height)
}
