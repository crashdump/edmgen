// Package edm implement the main edm struct, which offers the following:
//
// SelectFiles: Walks through all the subdirectories of the path specified, collating a list of relevant files.
// ExamineFiles: Reads and sample the lines from the output of `SelectFiles`
// SampleContent: (Optional) Filters the result of `ExamineFiles`
//
// The above consumes filters.
package edm
