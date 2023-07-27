# DLP Exact Data Match Sampler

[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/crashdump/edmgen/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/crashdump/edmgen?status.svg)](https://godoc.org/github.com/crashdump/edmgen)

Exact Data Match solutions is a Data Loss Prevention (DLP) technique that “fingerprints” sensitive data from structured data sources, such as source code or databases, and then watches for attempts to move the fingerprinted data. If a fingerprint is matched, the DLP will usually block the data movement to stop it from being shared or transferred inappropriately.

While a very useful technology, most DLP solutions lack the tools to sample the data from very large data repositories. This tool aims to close this gap.

## Install

```bash
go install github.com/crashdump/edmgen/cmd/edmgen@latest
```

## Use

```bash
edmgen ./folder_to_fingerprint/ > edm.txt
```

    ┌────────┐
    │ EDMGEN │
    └────────┘

    > Searching for relevant files...
      Found 38437 files.

    > Examining files...
      Got 37849 lines.

    > Sampling content...
      Sampled down to 23396 lines
    
    Complete!

## Phases

`SelectFiles`: Walks through all the subdirectories of the path specified, collating a list of relevant files.
`ExamineFiles`: Reads and sample the lines from the output of `SelectFiles`
`SampleContent`: (Optional) Filters the result of `ExamineFiles`

## Sampling and filtering

Most DLP tools have a limit for the number of lines that can be fingerprinted, hence it is very important to select the right files, and ultimately, lines.

This tools currently offers two types of filters `file` and `content`. 

### File

Can be applied to the phase: `SelectFiles`

* `IgnoreDirname`: Exclude directories and all their content based on their name(s)
* `IgnoreFilename`: Exclude files based on their name(s)
* `RequireExtension`: Only select files with specific extension(s)
* `IgnoreExtension`: Excludes all the files with specific extension(s)

### Content

Can be applied to the phases: `ExamineFiles` and `SampleContent`

* `LineLength`: Only select lines based on their length. Min and Max can be specified.
* `LongestLine`: Only select the longest line in the file.
* `IgnoreLine`: Ignore any line containing a specified string.
* `Uniq`: Deduplicate content. Especially useful during the final `SampleContent` phase.

Note: All filters are implemented as their own self-contained function, which are easily extensible. Implementing your own filter should not require any changes to the core code.

## Performance

Performance will vary depending on the size of the repository and the filters applied but a simple run on the Linux source code takes roughly ~2.5s.

## Contribute

### Build

```bash
go build ./... -o bin/edmgen
```

### Test

Note: This will automatically pull the Linux sources in the `test/linux` directory; they are used as fixtures for the tests.

```bash
go test ./...
```

