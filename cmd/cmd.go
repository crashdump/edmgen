package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"edmgen/pkg/edm"
	"edmgen/pkg/filters/content"
	"edmgen/pkg/filters/file"
	"github.com/urfave/cli/v2"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stderr, "", 0)
}

func main() {
	logger.Println("EDMGEN")
	logger.Println("------")

	app := &cli.App{
		Name:     "edmgen",
		Usage:    "Walk files in a directory and samples content for Cloudflare's Exact Data Match DLP",
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
				logger.Fatal("You need to specify a path to a folder")
			}
			return nil
		},
		Action: func(cCtx *cli.Context) error {
			edmc, err := edm.New(edm.Opts{})
			if err != nil {
				logger.Fatal(err)
			}

			path := cCtx.Args().First()

			logger.Print("* Searching for relevant files...")
			err = edmc.SelectFiles(path,
				file.IgnoreFilename(ignoreFilenames),
				file.IgnoreDirname(ignoreDirnames),
				file.RequireExtensions(requireExtentions),
			)
			if err != nil {
				logger.Fatal(err)
			}
			logger.Printf(" got %d files.\n", len(edmc.Paths))

			logger.Println("* Examining files..")
			err = edmc.ExamineFiles(
				content.LineLength(40, 100),
				content.LongestLine,
			)
			if err != nil {
				logger.Fatal(err)
			}
			logger.Printf(" got %d lines.\n", len(edmc.Content))

			logger.Println("* Sampling content..")
			lines := edmc.SampleContent(
				content.Uniq,
			)
			logger.Printf(" sampled down to %d lines.\n", len(edmc.Content))

			for _, line := range lines {
				fmt.Println(line)
			}

			logger.Println("* Complete!")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		logger.Fatal(err)
	}
}
