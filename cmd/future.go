/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"verse/pkg/future"

	"github.com/spf13/cobra"
)

// futureCmd represents the future command
var futureCmd = &cobra.Command{
	Use:   "future",
	Short: "A look into the future",
	Long:  `Future is now, go for it.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("future called")
		future.OpenOpenAIPage()
		future.OpenUzufly()
	},
}

func init() {
	rootCmd.AddCommand(futureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// futureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}
