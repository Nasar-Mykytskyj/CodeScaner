package main

import (
	"code_scanner/pkg/flags"
	"code_scanner/pkg/helper"
	"code_scanner/pkg/logger"
	"fmt"
	"os"
)

func main() {
	logger := logger.GetGeneralLogger()

	if err := flags.ParseAndValidateFlag(); err != nil {
		logger.Println(err)
		os.Exit(1)
	}

	confFlags := flags.GetConfFlags()

	fmt.Println(helper.GetAllFilesInDirectory(confFlags.SrcPath))
}
