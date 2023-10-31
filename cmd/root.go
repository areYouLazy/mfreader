package cmd

import (
	"fmt"
	"os"

	"github.com/areYouLazy/mfreader/mifare"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mfreader",
	Short: "Print MIFARE dumps in human readable format",
	Long: `mfreader, a tool inspired by mfdread ( https://github.com/zhovner/mfdread )
to analyze and print mifare dumps in human readable format`,
	Run: func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetString("file")

		//open file
		fh, err := os.ReadFile(file)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// get flags
		json, _ := cmd.Flags().GetBool("json")
		manufacturer, _ := cmd.Flags().GetBool("manufacturer")

		if json && manufacturer {
			mifare.ParseAnsJSONPrintMIFAREManufacturerBinaryDump(fh)
			return
		}

		if json {
			mifare.ParseAnsJSONPrintMIFAREBinaryDump(fh)
			return
		}

		if manufacturer {
			mifare.ParseAndPrettyPrintMIFAREManufacturerBinaryDump(fh)
			return
		}

		mifare.ParseAndPrettyPrintMIFAREBinaryDump(fh)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mfreader.yaml)")
	rootCmd.PersistentFlags().StringP("file", "f", "", "dump file to analyze")
	rootCmd.MarkPersistentFlagRequired("file")

	rootCmd.PersistentFlags().BoolP("json", "j", false, "output in JSON format")
	rootCmd.PersistentFlags().BoolP("manufacturer", "m", false, "print manufacturer data")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// initConfig reads in config file and ENV variables if set.
// func initConfig() {
// 	if cfgFile != "" {
// 		// Use config file from the flag.
// 		viper.SetConfigFile(cfgFile)
// 	} else {
// 		// Find home directory.
// 		home, err := os.UserHomeDir()
// 		cobra.CheckErr(err)

// 		// Search config in home directory with name ".mfreader" (without extension).
// 		viper.AddConfigPath(home)
// 		viper.SetConfigType("yaml")
// 		viper.SetConfigName(".mfreader")
// 	}

// 	// read in environment variables that match
// 	viper.AutomaticEnv()

// 	// If a config file is found, read it in.
// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
// 	}
// }
