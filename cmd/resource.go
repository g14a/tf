/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/g14a/tf/boilerplate"
	"github.com/g14a/tf/file"
	"github.com/g14a/tf/terraform"
	"github.com/spf13/cobra"
)

// resourceCmd represents the resource command
var resourceCmd = &cobra.Command{
	Use:   "resource",
	Short: "Select resource information in your Terraform configuration",
	Run: func(cmd *cobra.Command, args []string) {
		provider, _ := cmd.Flags().GetString("provider")
		boilerPlate, _ := cmd.Flags().GetBool("boilerplate")
		resource, _ := cmd.Flags().GetString("resource")

		if boilerPlate {
			boilerplate.SelectResourceBP(provider, resource)
			if resource == "" {
				terraform.SelectResourceTree(provider, resource, boilerPlate)
			}
		} else {
			file.Prompt()
			if provider != "" {
				// if provider is provided, bypass the provider
				// prompt and directly select resources
				terraform.SelectResourceTree(provider, resource, false)
			} else {
				// if provider is not provided in flags, give the provider prompt
				provider := terraform.ProvidersPrompt()
				terraform.SelectResourceTree(provider, resource, false)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(resourceCmd)
	resourceCmd.Flags().StringP("provider", "p", "", "Specify provider directly\ne.g. tf resource --provider aws\n")
	resourceCmd.Flags().StringP("resource", "r", "", "Specify resource directly\ne.g tf resource -p aws -r aws_instance\n")
	resourceCmd.Flags().BoolP("boilerplate", "b", false, "Boilerplate configuration for the resource\ne.g tf resource -p aws -r aws_instance -b\n")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resourceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resourceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
