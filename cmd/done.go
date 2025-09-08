/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/luka-03256/go-cli/todo"

)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: doneRun,
}


func doneRun(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatal("Error: You must specify the index of the todo to mark as done.\nUsage: go-cli done <index>")
	}

	index, err := strconv.Atoi(args[0])
	if err != nil || index < 1 {
		log.Fatalf("Invalid index: %v", args[0])
	}

	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Fatalf("Failed to read todos: %v", err)
	}
	if index > len(items) {
		log.Fatalf("Todo %d does not exist.", index)
	}

	items[index-1].Done = true

	err = todo.SaveItems(dataFile, items)
	if err != nil {
		log.Fatalf("Failed to save todos: %v", err)
	}
	fmt.Printf("Marked todo %d as done.\n", index)
}


func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
