/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:   "Verification of contracts",
	Short: "Verifies if the parties in business have mutually agreed upon the same rules",
	Long:  `A cli tool to compare the contracts. The required parameter is the path where the contract.yaml is present`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("verification yet to be implemented.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)
	verifyCmd.Flags().StringVarP(&contractPath, "path", "f", "", "Path to the contract file")

}
