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
	"strings"

	"github.com/FabriceT/tisax/internal"
	"github.com/FabriceT/tisax/internal/evaluation"
	"github.com/FabriceT/tisax/internal/markdown"
	"github.com/FabriceT/tisax/internal/models"
	"github.com/spf13/cobra"
)

const syntheseTableHeader = `
 ISA | Question | Maturité
:---:|----------|:-------:
`

var outputfile string
var summary bool
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "export tisax report to a PDF File",
	Long:  "Write a evaluation report to a PDF file",
	Run: func(cmd *cobra.Command, args []string) {
		evaldir, _ := cmd.Flags().GetString("evaldir")
		yamlFile, _ := cmd.Flags().GetString("file")

		evaluation.LoadYAML(yamlFile)

		/* Markdown content for results table
		## Synthèse

		ISA | Question | Maturité
		----|----------|:-------:
		X.X | What?    |   0  :D
		Table: qqchos

		*/

		var summaryBuilder strings.Builder
		summaryBuilder.WriteString("## Synthèse\n")

		// We add head.md if it exists.
		markdown.IncludeMDFile(evaldir + "/head.md")

		catalogs := evaluation.GetAllCatalogs()
		for _, catalog := range catalogs {

			markdown.AddCatalog(catalog)
			SummaryBeginTable(&summaryBuilder)

			for _, chapter := range catalog.Chapters {
				markdown.AddChapter(chapter)
				for _, assessment := range chapter.Assessments {
					for _, question := range assessment.Questions {
						result, _ := question.GetResult(path.Join(evaldir, catalog.Catalog))
						markdown.AddQuestion(question, result.MaturityLevel, result.Text)
						SummaryAddResult(&summaryBuilder, question, result)
					}
					markdown.AddNewLine()
				}
			}

			SummaryEndTable(&summaryBuilder, fmt.Sprintf("Maturité: %s\n", catalog.Catalog))
		}

		if summary {
			// On inclus la Synthese
			markdown.IncludeMDContent(summaryBuilder.String())
		}

		markdown.IncludeMDFile(evaldir + "/footer.md")
		markdown.Save(outputfile)
	},
}

func SummaryBeginTable(builder *strings.Builder) {
	builder.WriteString(syntheseTableHeader)
}

func SummaryEndTable(builder *strings.Builder, text string) {
	fmt.Fprintf(builder, "Table: %s\n", text)
}

func SummaryAddResult(builder *strings.Builder, question models.QuestionEntry, result models.EvaluationResult) {
	if result.MaturityLevel == -1 {
		// Not evaluated
		//|-----|----------|-----------|
		//| ISA | Question |  -        |
		//|-----|----------|-----------|
		fmt.Fprintf(builder, "%s | %s | -\n",
			question.Isa,
			question.Name)
	} else {
		//|-----|----------|-----------|
		//| ISA | Question | Note Icon |
		//|-----|----------|-----------|
		fmt.Fprintf(builder, "%s | %s | %d %s\n",
			question.Isa,
			question.Name,
			result.MaturityLevel,
			internal.GetMaturityIcon(result.MaturityLevel))
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&outputfile, "output", "evaluation.html", "where to store HTML result")
	rootCmd.PersistentFlags().BoolVar(&summary, "summary", true, "Add summary to document")
	rootCmd.AddCommand(exportCmd)
}
