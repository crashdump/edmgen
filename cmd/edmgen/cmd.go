package main

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/crashdump/edmgen/pkg/edm"
	"github.com/crashdump/edmgen/pkg/filters/content"
	"github.com/crashdump/edmgen/pkg/filters/file"
)

var logger *logging

func main() {
	logger = newLogger()

	logger.print("┌────────┐")
	logger.print("│ EDMGEN │")
	logger.print("└────────┘")
	logger.print("")

	app := &cli.App{
		Name:     "edmgen",
		Usage:    "Walk files in a directory and extract samples for the purporse of Exact Data Match DLPs",
		Compiled: time.Now(),
		Authors: []*cli.Author{{
			Name:  "Adrien Pujol",
			Email: "ap@cdfr.net",
		}},
		Flags: []cli.Flag{
			&cli.StringSliceFlag{Name: "extensions", Aliases: []string{"E"}},
		},
		Before: func(cCtx *cli.Context) error {
			if cCtx.Args().Len() < 1 {
				logger.printFatal("You need to specify a path to a folder")
			}
			return nil
		},
		Action: run,
	}

	if err := app.Run(os.Args); err != nil {
		logger.printFatal(err.Error())
	}
}

func run(cCtx *cli.Context) error {
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
		content.LineLength(40, 100),
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
	logger.printfResult("Sampled down to %d lines", len(edmc.Content))
	logger.print("")

	/*
	 * Finally, output the result
	 */
	for _, line := range lines {
		fmt.Println(line)
	}

	logger.print("Complete!")
	return nil
}
