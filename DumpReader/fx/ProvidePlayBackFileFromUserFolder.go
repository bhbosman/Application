package fx

import (
	"go.uber.org/fx"
	"log"
	"os"
	"path"
)

type PlayBackFile string

func ProvidePlayBackFileFromUserFolder(fileName string) fx.Option {
	return fx.Provide(func(logger *log.Logger, userHomeDir UserHomeDir) (PlayBackFile, error) {
		fileName := path.Join(string(userHomeDir), fileName)
		if _, err := os.Stat(fileName); err != nil {
			if os.IsNotExist(err) {
				logger.Printf("File does not exist. Error: %v\n", err)
				return "", err
			}
		}
		logger.Printf("PlaybackFileName is %v\n", fileName)
		return PlayBackFile(fileName), nil
	})
}

//
