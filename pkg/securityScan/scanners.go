package securityScan

import (
	"code_scanner/pkg/models"
	"regexp"
	"strings"
)

var regexCrossSite *regexp.Regexp
var regexSQL *regexp.Regexp

type ScanData struct {
	Line       string
	FilePath   string
	LineNumber int64
	FileType   string
}

type Scanner interface {
	Scan(ScanData) models.Vulnerability
}

type CrossSiteScanner struct {
}

type SQLScanner struct {
}

type SensitiveDataScanner struct {
}

func init() {
	regexCrossSite = regexp.MustCompile(`(Alert\((.*)\))`)
	regexSQL = regexp.MustCompile("\"([^\"]*\\bSELECT\\b[^\"]*\\bWHERE\\b[^\"]*%s[^\"]*)\"")
}

func (scanner *CrossSiteScanner) Scan(scanData ScanData) models.Vulnerability {
	if regexCrossSite.MatchString(scanData.Line) {
		return models.Vulnerability{
			File: scanData.FilePath,
			Line: scanData.LineNumber,
			Type: models.CrossSiteScripting,
		}
	}

	return models.Vulnerability{}
}

func (scanner *SQLScanner) Scan(scanData ScanData) models.Vulnerability {
	if regexSQL.MatchString(scanData.Line) {
		return models.Vulnerability{
			File: scanData.FilePath,
			Line: scanData.LineNumber,
			Type: models.SQlInjection,
		}
	}

	return models.Vulnerability{}
}

func (scanner *SensitiveDataScanner) Scan(scanData ScanData) models.Vulnerability {
	sensitiveData := []string{"Checkmarx", "Hellman & Friedman", "$1.15b"}

	counter := 0

	for _, val := range sensitiveData {
		if strings.Contains(scanData.Line, val) {
			counter++
		}
	}

	if counter == len(sensitiveData) {
		return models.Vulnerability{
			File: scanData.FilePath,
			Line: scanData.LineNumber,
			Type: models.SensitiveData,
		}
	}

	return models.Vulnerability{}
}

func GetScanners(fileType string) []Scanner {
	switch fileType {
	case ".js":
		return ([]Scanner{&CrossSiteScanner{}, &SensitiveDataScanner{}, &SQLScanner{}})
	default:
		return ([]Scanner{&SensitiveDataScanner{}, &SQLScanner{}})
	}
}
