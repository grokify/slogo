package slogo

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/OpenSLO/go-sdk/pkg/openslo"
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
	"github.com/OpenSLO/go-sdk/pkg/openslosdk"
)

type Objects []openslo.Object

func Read(r io.Reader) (Objects, error) {
	var out []openslo.Object
	b, err := io.ReadAll(r)
	if err != nil {
		return out, err
	}
	var osfmt openslosdk.ObjectFormat
	if len(b) > 0 && b[0] == '[' {
		osfmt = openslosdk.FormatJSON
	} else {
		osfmt = openslosdk.FormatYAML
	}
	return openslosdk.Decode(bytes.NewReader(b), osfmt)
}

func ReadFile(filename string) (Objects, error) {
	f, err := os.Open(filename)
	if err != nil {
		return Objects{}, err
	}
	defer func() {
		_ = f.Close()
	}()
	return Read(f)
}

func (objs Objects) Validate() error {
	for _, obj := range objs {
		if err := obj.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (objs Objects) WriteFile(filename string, format openslosdk.ObjectFormat) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func() {
		_ = f.Close()
	}()
	return openslosdk.Encode(f, format, objs...)
}

func (objs Objects) WriteFileJSON(filename string) error {
	return objs.WriteFile(filename, openslosdk.FormatJSON)
}

func (objs Objects) WriteFileYAML(filename string) error {
	return objs.WriteFile(filename, openslosdk.FormatYAML)
}

func (objs *Objects) AddSLOs(slos ...v1.SLO) {
	for _, slo := range slos {
		*objs = append(*objs, slo)
	}
}
