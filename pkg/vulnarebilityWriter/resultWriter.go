package vulnarebilityWriter

import (
	"bufio"
	"code_scanner/pkg/flags"
	"code_scanner/pkg/models"
	"encoding/json"
	"errors"
	"io"
)

type VulnarebilityWriter interface {
	Write(vulnarebility models.Vulnerability) error
}

type JsonVulnarebilityWriter struct {
	encoder *json.Encoder
}

type TextVulnarebilityWriter struct {
	writer *bufio.Writer
}

func (jvw *JsonVulnarebilityWriter) Write(vulnarebility models.Vulnerability) error {
	if jvw.encoder == nil {
		errors.New("Error JsonVulnarebilityWriter encoder field is null")
	}

	return jvw.encoder.Encode(&vulnarebility)
}

func (tvw *TextVulnarebilityWriter) Write(vulnarebility models.Vulnerability) error {
	if tvw.writer == nil {
		errors.New("Error TextVulnarebilityWriter writer field is null")
	}
	_, err := tvw.writer.WriteString(vulnarebility.String() + "\n")
	if err != nil {
		return err
	}

	return tvw.writer.Flush()
}

func GetVulnarebilityWriter(format flags.OutputFormat, out io.Writer) VulnarebilityWriter {
	switch format {
	case flags.TextFormat:
		return &TextVulnarebilityWriter{bufio.NewWriter(out)}
	case flags.JSONFormat:
		encoder := json.NewEncoder(out)
		encoder.SetIndent("", "  ")
		return &JsonVulnarebilityWriter{encoder}
	default:
		return nil
	}
}
