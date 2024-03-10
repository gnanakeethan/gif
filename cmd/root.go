/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/elewis787/boa"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gif",
	Short: "Golang port of Impact Engine by Green Software Foundation",
	Long: `GIF is a Golang port of Impact Engine by Green Software Foundation. 

It allows you to run Impact Engine without requiring you to install any external 
dependencies that can potentially harm your system. It will feature a fully verified / vetted
plugin management system in future that is interopable with Impact Engine's plugin system.`,
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
	styles := boa.DefaultStyles()
	styles.Title.BorderForeground(lipgloss.AdaptiveColor{Light: `#E3BD2D`, Dark: `#E3BD2D`})
	styles.Border.BorderForeground(lipgloss.AdaptiveColor{Light: `#E3BD2D`, Dark: `#E3BD2D`})
	styles.SelectedItem.Foreground(lipgloss.AdaptiveColor{Light: `#353C3B`, Dark: `#353C3B`}).
		Background(lipgloss.AdaptiveColor{Light: `#E3BD2D`, Dark: `#E3BD2D`})

	b := boa.New(boa.WithStyles(styles))

	rootCmd.SetUsageFunc(b.UsageFunc)
	rootCmd.SetHelpFunc(b.HelpFunc)

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gif.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
