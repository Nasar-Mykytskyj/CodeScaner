package main

import (
	"code_scanner/internal/logger"
	"code_scanner/pkg/flags"
	"code_scanner/pkg/security_scan"
	"os"
)

func main() {
	logger := logger.GetGeneralLogger()

	if err := flags.ParseAndValidateFlag(); err != nil {
		logger.Println(err)
		os.Exit(1)
	}

	confFlags := flags.GetConfFlags()

	security_scan.RunVulnerabilityScan(confFlags)
}
