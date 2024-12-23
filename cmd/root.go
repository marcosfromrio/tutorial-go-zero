package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
		Use:   "zero",
		Short: "Zero is a simple CLI tool for managing your projects",
		Long: "Zero is a simple CLI tool for managing your projects",
		Run: func(cmd *cobra.Command, args []string) {},

}

func Execute(){
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}