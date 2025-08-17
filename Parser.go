package mare

import (
	"fmt"
	"regexp"
	"strings"
)

// Parser config is used to set the function- and keyword
// delimiters
type ParserConfig struct {
	KeywordPrefix     string                    // Prefix for Keywords
	KeywordSuffix     string                    // Postfix for Keywords
	FunctionDelimiter string                    // Start-Delimiter for Function calls
	functions         map[string]ParserFunction // Functions
}

func (p *ParserConfig) RegisterFunc(funcName string, fn ParserFunction) error {
	_, existing := p.functions[strings.ToUpper(funcName)]
	if existing {
		return fmt.Errorf("%s already registered", funcName)
	}

	p.functions[strings.ToUpper(funcName)] = fn

	return nil
}

func (p *ParserConfig) ExecFunc(funcName string, funcParam string) string {
	parserFunc, isRegistered := p.functions[strings.ToUpper(funcName)]
	if !isRegistered {
		return funcParam
	}

	return parserFunc(funcParam)
}

// Returns a Quoted version of the parser config
// safe to use in regular expressions
func (p *ParserConfig) ToQuoted() ParserConfig {
	return ParserConfig{
		KeywordPrefix:     regexp.QuoteMeta(p.KeywordPrefix),
		KeywordSuffix:     regexp.QuoteMeta(p.KeywordSuffix),
		FunctionDelimiter: regexp.QuoteMeta(p.FunctionDelimiter),
	}
}

type ParserFunction func(string) string

func DefaultConfig() *ParserConfig {
	pc := &ParserConfig{
		KeywordPrefix:     "[",
		KeywordSuffix:     "]",
		FunctionDelimiter: "§",
		functions:         make(map[string]ParserFunction),
	}

	pc.functions["UPPER"] = Upper
	pc.functions["LOWER"] = Lower
	pc.functions["FIRSTLETTER"] = Firstletter

	return pc
}
