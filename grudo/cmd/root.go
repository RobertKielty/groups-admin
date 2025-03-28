/*
Copyright © 2025 Robert Kielty <robert.kielty@cncf.io>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "grudo",
	Short: "grudo is a pseudo admin tool for groups.io",
	Long: `grudo fakes admin account operations on groups.io by using your
owner mailing list subscriptions as a template administrator account that you 
can transfer to other users who need to have the same access as your account.

Most grudo commands work with two groups.io accounts. 

The source account is the account that you use to log into groups.io. Your 
account acts as the pseudo admin account, which means on every list where 
you are subscribed owner, use grudo to act as an administrator for that 
subset of lists on your groups.io installation. You must supply the password
for this account to grudo.

grudo provides the following sub-commands:
	list - list your mailing list subscriptions
	copy - copy your mailing list subscriptions to another user
	delete - remove a member from your mailing list subscriptions

`,
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// FIXME unresolved ref.
	rootCmd.PersistentFlags().StringVar(&cfgFile,
		"config",
		"",
		"config file (default is $HOME/.config/grudo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("targetUser", "tu", false, "Target groups.io user that your account will work on.")
}
