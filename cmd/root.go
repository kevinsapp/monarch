// Package cmd ...
package cmd

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

const cfgFileBaseName = "database"

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "monarch",
	Short: "Monarch is a tool for migrating SQL databases.",
	Long:  `Monarch is a tool for migrating SQL databases.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./database.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Determine project root directory.
		dir, err := rootDir()
		if err != nil {
			log.Fatalf("ERROR: initConfig: %s\n", err)
		}

		// Search for config file in the project root directory with name "database" (without extension).
		viper.AddConfigPath(dir)
		viper.SetConfigName("database")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If no config files is found, log an error.
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("ERROR: initConfig: could not read in config file %q (with extenstion .json, .toml, or .yaml)", cfgFileBaseName)
	}
}

// rootDir returns the project root directory.
func rootDir() (string, error) {
	var err error
	var dir string

	// Determine which file contains this function's caller.
	_, f, _, ok := runtime.Caller(0)
	if ok != true {
		err = errors.New("rootDir: could not determine project root directory")
		return dir, err
	}

	// Determine the file's parent directory.
	dir = filepath.Dir(f)

	// Return the parent of the file's directory (and err).
	return filepath.Dir(dir), err
}
