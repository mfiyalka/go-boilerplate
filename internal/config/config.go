package config

import "github.com/joho/godotenv"

func Load(path string) error { //nolint:revive
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
