package securityScan

import (
	"bufio"
	"code_scanner/internal/helper"
	"code_scanner/internal/logger"
	"code_scanner/pkg/flags"
	"code_scanner/pkg/models"
	"code_scanner/pkg/vulnarebilityWriter"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func RunVulnerabilityScan(conf flags.ConfFlags) {
	logger := logger.GetGeneralLogger()
	files, err := helper.GetAllFilesInDirectory(conf.SrcPath)

	if err != nil {
		logger.Println("Error during receiving files list from directory err:", err)
	}

	fmt.Printf("Start scan files: %v", files)

	scanChan := make(chan ScanData)
	resultChan := make(chan models.Vulnerability)

	defer func() {
		close(scanChan)
		close(resultChan)
	}()

	go ReadWriteResults(resultChan, conf.OutputPath, conf.OutputFormat)

	workerPool := NewWorkerPool(runtime.GOMAXPROCS(0), scanChan, resultChan)
	go workerPool.Run(context.TODO())

	for _, path := range files {
		ReadFileLineByLine(path, scanChan)
	}
}

func ReadFileLineByLine(path string, scanChan chan ScanData) {
	file, err := os.Open(path)
	if err != nil {
		logger.GetGeneralLogger().Println("Error opening file:", err)
		return
	}

	defer file.Close()

	fileType := filepath.Ext(path)
	scanner := bufio.NewScanner(file)

	var lineNumber int64 = 0

	for scanner.Scan() {
		data := ScanData{
			Line:       scanner.Text(),
			FilePath:   path,
			LineNumber: lineNumber,
			FileType:   fileType,
		}

		scanChan <- data

		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		logger.GetGeneralLogger().Println(err)
	}
}

func ReadWriteResults(resultChan chan models.Vulnerability, path string, format flags.OutputFormat) {
	file, err := os.Create(path)
	if err != nil {
		logger.GetGeneralLogger().Printf("Error creating results file: %v\n", err)
		return
	}

	defer file.Close()

	writer := vulnarebilityWriter.GetVulnarebilityWriter(format, file)

	for res := range resultChan {
		if err := writer.Write(res); err != nil {
			logger.GetGeneralLogger().Printf("Error during writing to result file err: %v", err)
		}
	}
}
