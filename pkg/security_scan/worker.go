package security_scan

import (
	"code_scanner/pkg/models"
	"context"
	"reflect"
	"sync"
)

type WorkerPool struct {
	workersNumber int
	scanDataChan  chan ScanData
	results       chan models.Vulnerability
	Done          chan struct{}
}

func NewWorkerPool(number int, scanDataChan chan ScanData, resultsChan chan models.Vulnerability) *WorkerPool {
	return &WorkerPool{
		workersNumber: number,
		scanDataChan:  scanDataChan,
		results:       resultsChan,
		Done:          make(chan struct{}),
	}
}

func worker(wg *sync.WaitGroup, scanDataChan <-chan ScanData, results chan<- models.Vulnerability) {
	defer wg.Done()
	for scanData := range scanDataChan {
		scanners := GetScanners(scanData.FileType)

		for _, scanner := range scanners {
			res := scanner.Scan(scanData)

			if !reflect.DeepEqual(models.Vulnerability{}, res) {
				results <- scanner.Scan(scanData)
			}
		}
	}
}

func (wp WorkerPool) Run(ctx context.Context) {
	var wg sync.WaitGroup

	for i := 0; i < wp.workersNumber; i++ {
		wg.Add(1)
		go worker(&wg, wp.scanDataChan, wp.results)
	}

	wg.Wait()
}

func (wp WorkerPool) Results() <-chan models.Vulnerability {
	return wp.results
}
