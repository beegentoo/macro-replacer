package mare

import "regexp"

// Parser config is used to set the function- and keyword
// delimiters
type ParserConfig struct {
	KeywordPrefix     string // Prefix for Keywords
	KeywordSuffix     string // Postfix for Keywords
	FunctionDelimiter string // Start-Delimiter for Function calls
}

// Returns a Quoted version of the parser config
// safe to use in regular expressions
func (p ParserConfig) ToQuoted() ParserConfig {
	return ParserConfig{
		KeywordPrefix:     regexp.QuoteMeta(p.KeywordPrefix),
		KeywordSuffix:     regexp.QuoteMeta(p.KeywordSuffix),
		FunctionDelimiter: regexp.QuoteMeta(p.FunctionDelimiter),
	}
}

var defaultConfig ParserConfig = ParserConfig{
	KeywordPrefix:     "[",
	KeywordSuffix:     "]",
	FunctionDelimiter: "§",
}
