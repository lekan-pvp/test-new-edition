package main

import (
	_ "embed"
	"encoding/json"
	"github.com/lekan-pvp/short/cmd/staticlint/internal/analyzer"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/structtag"
	"honnef.co/go/tools/staticcheck"
)

//go:embed linter.json
var data []byte

// ConfigData describes a struct of configuration file.
type ConfigData struct {
	Staticcheck []string
}

func main() {
	var cfg ConfigData
	if err := json.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}
	mychecks := []*analysis.Analyzer{
		printf.Analyzer,
		shadow.Analyzer,
		structtag.Analyzer,
		analyzer.OsExitCheck,
	}
	checks := make(map[string]bool)

	for _, v := range cfg.Staticcheck {
		checks[v] = true
	}

	for _, v := range staticcheck.Analyzers {
		if checks[v.Analyzer.Name] {
			mychecks = append(mychecks, v.Analyzer)
		}
	}
	multichecker.Main(
		mychecks...,
	)
}
