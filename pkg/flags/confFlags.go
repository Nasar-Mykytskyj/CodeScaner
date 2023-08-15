package flags

import (
	"code_scanner/internal/helper"
	"errors"
	"flag"
	"fmt"
)

type OutputFormat int

type ConfFlags struct {
	SrcPath      string
	OutputPath   string
	OutputFormat OutputFormat
}

const (
	TextFormat OutputFormat = iota
	JSONFormat
)

var confFlags ConfFlags

func init() {
	flag.StringVar(&confFlags.SrcPath, "src", "", "Path to source code directory")
	flag.StringVar(&confFlags.OutputPath, "out", "", "Path for output")
	flag.Var(&confFlags.OutputFormat, "format", "Format of report (text or json)")
}

func GetConfFlags() ConfFlags {
	return confFlags
}

func ParseAndValidateFlag() error {
	flag.Parse()

	empty, err := helper.IsEmptyDir(confFlags.SrcPath)
	if err != nil || empty {
		return errors.New(fmt.Sprintf("Source code directory does not exist or empty err:%v", err))
	}

	if isValid := helper.IsFilePathValid(confFlags.OutputPath); !isValid {
		return errors.New("Path for output result is not valid")
	}

	return nil
}

func (out OutputFormat) String() string {
	return []string{0: "txt", 1: "json"}[out]
}

func (out *OutputFormat) Set(val string) error {
	switch val {
	case "txt":
		*out = TextFormat
	case "json":
		*out = JSONFormat
	default:
		return fmt.Errorf("Invalid format: %s", val)
	}
	return nil
}
