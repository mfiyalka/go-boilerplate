package config

import "github.com/joho/godotenv"

// Load loads environment variables from a file specified by the given path.
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
