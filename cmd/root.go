/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd


import (
	"os"

	"github.com/spf13/cobra"

	"fmt"

	"github.com/luka-03256/go-cli/todo"

)


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-cli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var dataFile string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	/*
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	*/

	err := rootCmd.Execute()
	// deprecated version of error handling
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	/*err := homedir.Dir()
	if err != nil {
		log.Println("Unable to detect home directory. Please set data file using --datafile.")
	}*/

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-cli.yaml)")

	rootCmd.PersistentFlags().StringVar(&dataFile, "datafile", "todos.json", "data file to store todos")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(todo.TodoCmd)
}


