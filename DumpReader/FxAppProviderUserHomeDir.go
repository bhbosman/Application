package main

import (
	"go.uber.org/fx"
	"log"
	"os"
)

type UserHomeDir string

func FxAppProviderUserHomeDir(path string) fx.Option {
	if path != "" {
		return fx.Provide(
			func(logger *log.Logger) (UserHomeDir, error) {
				logger.Printf("User Path is %v\n", string(path))
				return UserHomeDir(path), nil
			})
	}
	return fx.Provide(
		func(logger *log.Logger) (UserHomeDir, error) {
			userHomeDir, err := os.UserHomeDir()
			if err != nil {
				logger.Printf("Could not find local user folder. Error: %v\n", err)
			}
			logger.Printf("User Path is %v\n", string(userHomeDir))
			return UserHomeDir(userHomeDir), nil
		})

}
