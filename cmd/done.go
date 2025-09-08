package cmd

import (
    "fmt"
    "log"
    "sort"
    "strconv"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "github.com/luka-03256/go-cli/todo"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
    Use:     "done [index]",
    Aliases: []string{"do"},
    Short:   "Mark a todo item as done",
    Long: `Marks the todo item at the given index as completed.

Example:
  go-cli done 2
  go-cli do 3`,
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

    items, err := todo.ReadItems(viper.GetString("dataFile"))
    if err != nil {
        log.Fatalf("Failed to read todos: %v", err)
    }

    if index > len(items) {
        log.Fatalf("Todo %d does not exist. Total items: %d", index, len(items))
    }

    // Mark as done
    items[index-1].Done = true
    fmt.Printf("✓ Marked \"%s\" as done.\n", items[index-1].Text)

    // Optional: Sort again so done items appear last
    sort.Sort(todo.ByPriority(items))

    // Save updated list
    if err := todo.SaveItems(viper.GetString("dataFile"), items); err != nil {
        log.Fatalf("Failed to save todos: %v", err)
    }
}

func init() {
        rootCmd.AddCommand(doneCmd)
        // Here you will define your flags and configuration settings.

        // Cobra supports Persistent Flags which will work for this command
        // and all subcommands, e.g.:
        // addCmd.PersistentFlags().String("foo", "", "A help for foo")

        // Cobra supports local flags which will only run when this command
        // is called directly, e.g.:
        // addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


