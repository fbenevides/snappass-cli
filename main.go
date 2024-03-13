package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

func defineSnapUrlCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "url [url]",
		Short: "Configure Snappass API URL",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			config := &Config{BaseUrl: args[0]}
			err := WriteConfig(config)
			if err == nil {
				fmt.Printf("• Snappass URL: %s", config.BaseUrl)
				return nil
			}

			return err
		},
	}
}

func setPasswordCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "set [password]",
		Short: "Sets a new password",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := ReadConfig()
			if err != nil {
				return err
			}

			snapClient := NewClient(config.BaseUrl)
			response, err := snapClient.SetPassword(args[0])
			if err != nil {
				return err
			}

			if response.IsSuccessful() {
				clipboard.WriteAll(response.Link)
				fmt.Printf("• Secret link copied to clipboard. TTL: %d\n", response.Ttl)
				return nil
			}

			return fmt.Errorf("unexpected server error: %d", response.Status)
		},
	}
}

func main() {
	rootCommand := &cobra.Command{
		Use:   "snap",
		Short: "snap is a CLI for Snappass (https://github.com/pinterest/snappass)",
	}

	defineBaseUrlCommand := defineSnapUrlCommand()
	setPasswordCommand := setPasswordCommand()

	rootCommand.AddCommand(defineBaseUrlCommand)
	rootCommand.AddCommand(setPasswordCommand)
	rootCommand.Execute()
}
