/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"github.com/spf13/cobra"
)

var (
	yamlFile            string
	evaluationDirectory string
)

var rootCmd = &cobra.Command{
	Use:   "tisax",
	Short: "Manage TISAX evalutation",
	Long:  "Do something with TISAX evalution.",
	//Run: func(cmd *cobra.Command, args []string) { fmt.Println("Hello from root cmd") },
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVar(&evaluationDirectory, "evaldir", "./evaluation", "evaluation directory (default ./evaluation")
	rootCmd.PersistentFlags().StringVar(&yamlFile, "file", "tisax.yml", "evaluation file (default tisax.yml")
}
