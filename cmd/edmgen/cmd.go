package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v2"
	"golang.org/x/exp/slices"

	"github.com/crashdump/edmgen/pkg/edm"
	"github.com/crashdump/edmgen/pkg/filters/content"
	"github.com/crashdump/edmgen/pkg/filters/file"
)

var logger *logging

const STDOUT = "stdout"

func init() {
	cli.HelpFlag = &cli.BoolFlag{Name: "help"}
	cli.VersionFlag = &cli.BoolFlag{Name: "version", Aliases: []string{"v"}}
}

func main() {
	logger = newLogger()
	var flagOutput string
	var flagFormat string

	logger.print("┌────────┐")
	logger.print("│ EDMGEN │")
	logger.print("└────────┘")
	logger.print("")

	app := &cli.App{
		Name:     "edmgen",
		Usage:    "Walk files in a directory and extract (samples) lines",
		Compiled: time.Now(),
		Authors: []*cli.Author{{
			Name:  "Adrien Pujol",
			Email: "ap@cdfr.net",
		}},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Destination: &flagOutput,
				Action: func(ctx *cli.Context, v string) error {
					if v == "" {
						return errors.New("please specify a filename")
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:        "format",
				Aliases:     []string{"f"},
				DefaultText: STDOUT,
				Destination: &flagFormat,
				Action: func(ctx *cli.Context, v string) error {
					formats := []string{STDOUT, "csv", "txt"}
					if slices.Contains[[]string](formats, flagFormat) {
						return fmt.Errorf("output %s not currently supported", flagFormat)
					}
					return nil
				},
			},
		},
		Before: func(cCtx *cli.Context) error {
			if cCtx.Args().Len() < 1 {
				logger.print("You need to specify a path to a folder")
				os.Exit(1)
			}
			return nil
		},
		Action: func(cCtx *cli.Context) error {
			edmc, err := edm.New(edm.Opts{})
			if err != nil {
				logger.printFatal(err.Error())
			}

			path := cCtx.Args().First()

			/*
			 * Search for all the relevant files
			 */
			logger.printHeader("Searching for relevant files...")
			err = edmc.SelectFiles(path,
				file.IgnoreFilename(ignoreFilenames),
				file.IgnoreDirname(ignoreDirnames),
				file.RequireExtensions(requireExtentions),
			)
			if err != nil {
				logger.printFatal(err.Error())
			}
			logger.printfResult("Found %d files.", len(edmc.Paths))
			logger.print("")

			/*
			 * Walk through each of the files and sample lines
			 */
			logger.printHeader("Examining files...")
			err = edmc.ExamineFiles(
				content.IgnoreLine("serialVersionUID"),
				content.LineLength(60, 120, true),
				content.LongestLine,
			)
			if err != nil {
				logger.printFatal(err.Error())
			}
			logger.printfResult("Selected %d lines.", len(edmc.Content))
			logger.print("")

			/*
			 * Sample the lines previously selected down to the result
			 */
			logger.printHeader("Sampling content...")
			lines := edmc.SampleContent(
				content.Uniq,
			)
			logger.printfResult("Sampled down to %d lines", len(lines))
			logger.print("")

			/*
			 * Finally, output the result
			 */
			switch flagFormat {
			case "txt":
				writeFileTxt(flagOutput, lines)

			case "csv":
				writeFileCsv(flagOutput, lines)

			default:
				writeStdout(lines)
			}

			logger.print("Complete!")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		logger.printFatal(err.Error())
	}
}

func writeStdout(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}

func writeFileTxt(filename string, lines []string) {

}

func writeFileCsv(filename string, lines []string) {

}
