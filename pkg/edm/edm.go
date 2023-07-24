package edm

import (
	"bufio"
	"errors"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
	"unicode/utf8"

	"github.com/crashdump/edmgen/pkg/filters/content"
	"github.com/crashdump/edmgen/pkg/filters/file"
)

type Opts struct{}

type Filter struct {
	MinLineLength     int
	RequireExtensions []string
	IgnoreFolder      []string
}

type Edm struct {
	rand    *rand.Rand
	opts    Opts
	Content []string
	Paths   []string
}

func New(opts Opts) (Edm, error) {
	return Edm{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
		opts: opts,
	}, nil
}

func (e *Edm) SelectFiles(dir string, filters ...file.Filter) (err error) {
	err = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		for _, filter := range filters {
			match, err := filter(d)
			if err != nil {
				return err
			}
			if !match {
				return nil
			}
		}

		if !d.IsDir() {
			e.Paths = append(e.Paths, path)
		}

		return nil
	})

	return err
}

func (e *Edm) ExamineFiles(filters ...content.Filter) (err error) {
	if len(e.Paths) == 0 {
		return errors.New("no files in scope, please select target files first")
	}

	for _, path := range e.Paths {
		readFile, err := os.Open(path)
		if err != nil {
			log.Printf("[WARN] could not open file: %s", err)
			continue // skip the file
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)

		// Iterate over the file lines
		lines := make([]string, 0)
		for fileScanner.Scan() {
			line := fileScanner.Text()

			// Ignore anything non-UTF8
			if !utf8.ValidString(line) {
				continue
			}

			lines = append(lines, line)
		}

		// We apply the filters at the "file level".
		e.Content = append(e.Content, ApplyFilters(lines, filters...)...)
	}

	return nil
}

func (e *Edm) SampleContent(filters ...content.Filter) []string {
	e.Content = ApplyFilters(e.Content, filters...)
	return e.Content
}

func ApplyFilters(lines []string, filters ...content.Filter) []string {
	for _, filter := range filters {
		lines = filter(lines)
	}
	return lines
}
