package main

import (
	"code_scanner/pkg/flags"
	"code_scanner/pkg/logger"
	"code_scanner/pkg/securityScan"
	"os"
)

func main() {
	logger := logger.GetGeneralLogger()

	if err := flags.ParseAndValidateFlag(); err != nil {
		logger.Println(err)
		os.Exit(1)
	}

	confFlags := flags.GetConfFlags()

	securityScan.RunVulnerabilityScan(confFlags)
}
