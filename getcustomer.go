package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var nifVar int
var tokenVar string

func init() {
	rootCmd.PersistentFlags().IntVar(&nifVar, "nif", 0, "nif number")
	rootCmd.PersistentFlags().StringVar(&tokenVar, "token", "", "token code")
	rootCmd.AddCommand(getCustomerCmd)
}

// go run . getcustomer --nif 123456789 --token exampletoken
var getCustomerCmd = &cobra.Command{
	Use:   "getcustomer",
	Short: "getcustomer gets a customer",
	Long:  "getcustomer gets a customer",
	Run: func(cmd *cobra.Command, args []string) {
		//check for nif and token args
		nif, err := cmd.Flags().GetInt("nif")
		if err != nil || nif == 0 {
			fmt.Println("failed to get string flag for nif")
			return
		}
		token, err := cmd.Flags().GetString("token")
		if err != nil || nif == 0 {
			fmt.Println("failed to get string flag for token")
			return
		}
		fmt.Println("")
		fmt.Println("nif:", nif)
		fmt.Println("token:", token)
		fmt.Println("")
	},
}
