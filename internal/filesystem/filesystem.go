// Package filesystem contains the filesystem interface.
package filesystem

import (
	"codegen/internal/flagkeys"

	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

func FileSystem() afero.Fs {
	return viper.Get(flagkeys.FileSystem).(afero.Fs)
}

func init() {
	viper.SetDefault(flagkeys.FileSystem, afero.NewOsFs())
}
