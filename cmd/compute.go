/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/gnanakeethan/gif/internal/file"
	"github.com/gnanakeethan/gif/internal/tea"
	"github.com/gnanakeethan/gif/internal/types"
	"github.com/kr/pretty"
	"github.com/spf13/cobra"

	"gopkg.in/yaml.v3"
)

// computeCmd represents the compute command
var computeCmd = &cobra.Command{
	Use:   "compute",
	Short: "computes impact of manifest definition and its dependencies.",
	Long:  `This command takes a superset manifest yaml defined by IF specification. It computes the impact of the infra and its dependencies. It outputs the computed infra and its dependencies in a YAML file.`,
	Run: func(cmd *cobra.Command, args []string) {
		// load the manifest
		var inputData types.Input
		if fileBytes, err := file.Read(*manifest); err != nil {
			panic(err)
		} else {
			if err := yaml.Unmarshal(fileBytes, &inputData); err != nil {
				panic(err)
			}
		}
		pretty.Println(inputData)
		tea.Compute(inputData)
	},
}
var manifest *string

func init() {
	rootCmd.AddCommand(computeCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	manifest = computeCmd.PersistentFlags().String("manifest", "manifest.yaml", "Path for manifest YAML file. The file defines your infra and its dependencies.")
	computeCmd.MarkPersistentFlagRequired("manifest")
	computeCmd.PersistentFlags().String("output", "", "Path for the output file. It will be a YAML file containing the computed infra and its dependencies.")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//computeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
