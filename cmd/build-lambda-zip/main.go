// Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved

package main

import (
	"archive/zip"
	"flag"
	"fmt" //nolint: staticcheck
	"log"
	"os"
	"path/filepath"
)

const usage = `build-lambda-zip - Puts an executable and supplemental files into a zip file that works with AWS Lambda.
usage:
  build-lambda-zip [options] handler-exe [paths...]
options:
  -o, --output  output file path for the zip. (default: ${handler-exe}.zip)
  -h, --help    prints usage
`

func main() {
	var outputZip string
	flag.StringVar(&outputZip, "o", "", "")
	flag.StringVar(&outputZip, "output", "", "")
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
	}
	flag.Parse()
	if len(flag.Args()) == 0 {
		log.Fatal("no input provided")
	}
	inputExe := flag.Arg(0)
	if outputZip == "" {
		outputZip = fmt.Sprintf("%s.zip", filepath.Base(inputExe))
	}
	if err := compressExeAndArgs(outputZip, inputExe, flag.Args()[1:]); err != nil {
		log.Fatalf("failed to compress file: %v", err)
	}
	log.Printf("wrote %s", outputZip)
}

func writeExe(writer *zip.Writer, pathInZip string, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// indicates Unix
// -rwxrwxrwx file permissions

func compressExeAndArgs(outZipPath string, exePath string, args []string) error {
	_ = "STUB: not implemented"
	return nil
}
