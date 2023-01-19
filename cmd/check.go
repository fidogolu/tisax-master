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
	"fmt"
	"path"

	"github.com/FabriceT/tisax/internal"
	"github.com/FabriceT/tisax/internal/evaluation"
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check TISAX evaluation",
	Long:  "Checks if a TISAX evaluation is completed",
	Run: func(cmd *cobra.Command, args []string) {

		checked := true
		evaldir, _ := cmd.Flags().GetString("evaldir")
		yamlFile, _ := cmd.Flags().GetString("file")

		evaluation.LoadYAML(yamlFile)

		catalogs := evaluation.GetAllCatalogs()
		for _, catalog := range catalogs {
			fmt.Println("====", catalog.Catalog, "====")
			for _, chapter := range catalog.Chapters {
				fmt.Println(chapter.Isa, ") ", chapter.Chapter)
				questions := chapter.GetAllQuestions()

				for _, q := range questions {
					// On se fiche de l'erreur
					result, _ := q.GetResult(path.Join(evaldir, catalog.Catalog))

					fmt.Printf(" %s) %s - %d %s\n",
						q.Isa,
						q.Name,
						result.MaturityLevel,
						internal.GetMaturityIcon(result.MaturityLevel))

					if result.MaturityLevel < 3 {
						checked = false
					}
				}
			}
		}

		if checked {
			fmt.Println("\nEvaluation successful! \u2705")
		} else {
			fmt.Println("\nCheck failed points! \u274C")

		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
