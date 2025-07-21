package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "newsbot",
	Short: "Newsbot CLI",
	Long:  "CLI утилита для работы с NewsBot",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Используйте одну из команд: migrate, fetch")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
