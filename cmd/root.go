package cmd

import (
	"errors"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "convertMan",
	SilenceUsage: false,
}

var level string
var outputFormat string
var configDir string

// Execute export
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}

func init() {

	// detect the log level
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {

		level, err := logrus.ParseLevel(level)
		if err != nil {
			return err
		}

		logrus.SetLevel(level)

		return nil
	}

	rootCmd.RunE = func(cmd *cobra.Command, paths []string) error {

		cmd.SilenceUsage = true
		cmd.SilenceErrors = false

		if len(paths) > 0 {

			// read from file

			return nil
		} else {

			// try read from pipe
			fileInfo, _ := os.Stdin.Stat()
			if fileInfo.Mode()&os.ModeCharDevice != 0 {

				// no pipe input, no file input, error tips and usage tips
				cmd.SilenceUsage = false
				return errors.New("input content or input file is neeed. ")

			} else {

				// read from pipe
				return nil

			}

		}

	}

	rootCmd.PersistentFlags().StringVarP(
		&level,
		"verbose",
		"v",
		logrus.WarnLevel.String(),
		"Log level (trace, debug, info, warn, error, fatal, panic) ",
	)

	rootCmd.PersistentFlags().StringVarP(
		&outputFormat,
		"output", "o", "json",
		"options output format: yaml, json",
	)

	home, err := homedir.Dir()
	if err != nil {
		logrus.Fatal(err)
	}

	rootCmd.PersistentFlags().StringVarP(
		&configDir,
		"config", "", path.Join(home, ".convertMan"),
		"location of config files like $CONVERT_MAN_CONFIG ",
	)
}
