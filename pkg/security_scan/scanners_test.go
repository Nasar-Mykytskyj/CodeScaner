package security_scan

import (
	"code_scanner/pkg/models"
	"reflect"
	"testing"
)

func TestScanners(t *testing.T) {
	tests := []struct {
		name     string
		scanData ScanData
		scanner  Scanner
		result   models.Vulnerability
	}{
		{
			name: "TestSQLScanner",
			scanData: ScanData{
				"err = db.Get(&jason, \"SELECT * FROM person WHERE first_name=%s\", \"Jason\")",
				"test",
				2,
				"json",
			},
			scanner: &SQLScanner{},
			result: models.Vulnerability{
				"test",
				2,
				models.SQlInjection,
			},
		},
		{
			name: "TestCrossSiteScanner",
			scanData: ScanData{
				"opts = { autoBom: !opts }, window.Alert(\"test\")",
				"test",
				3,
				"js",
			},
			scanner: &CrossSiteScanner{},
			result: models.Vulnerability{
				"test",
				3,
				models.CrossSiteScripting,
			},
		},
		{
			name: "TestSensitiveDataScanner",
			scanData: ScanData{
				"if (typeof opts === 'undefined') opts = \"Checkmarx\" \"Hellman & Friedman\" \"$1.15b\"\n",
				"test",
				3,
				"js",
			},
			scanner: &SensitiveDataScanner{},
			result: models.Vulnerability{
				"test",
				3,
				models.SensitiveData,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.scanner.Scan(tt.scanData)

			if !reflect.DeepEqual(result, tt.result) {
				t.Fatalf("Test %s is failed received value %v is not expected %v", tt.name, result, tt.result)
			}
		})
	}
}
