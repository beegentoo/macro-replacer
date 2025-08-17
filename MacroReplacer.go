package mare

import (
	"fmt"
	"regexp"
)

// Process a replacement using the default configuration:
// * Replacement Keywords are prepended with [ and postfixed with ]
// * Function calls are prepended with §
//
// Available functions are:
// * Upper: Convert to upper case string
// * Lower: Convert to lower case string
// * Firstletter: Return the first letter
//
// Example:
//
//	Keyword1 = "Hello"
//	Keyword2 = "World"
//	Input String: "[Keyword1] §upper([Keyword2])"
//
//	Will result in: "Hello WORLD"
//
// Important: Nesting of functions is not supported
func Process(str string, replacements map[string]any) string {
	return ProcessWithConfig(str, replacements, DefaultConfig())
}

func ProcessWithConfig(str string, replacements map[string]any, cfg *ParserConfig) string {
	processed := str

	// First loop: replace
	for currEntry := range replacements {
		re, err := makeReplRe(currEntry, cfg)
		if err != nil {
			continue
		}

		processed = re.ReplaceAllString(processed, replacements[currEntry].(string))
	}

	processed = evaluateFunctions(processed, cfg)

	return processed
}

func evaluateFunctions(str string, cfg *ParserConfig) string {
	processed := str

	// Second loop: function evaluation
	fcRegex, err := makeFuncRe(cfg)
	if err != nil {
		return ""
	}
	funcCalls := fcRegex.FindAllStringSubmatch(processed, -1)

	for _, currMatch := range funcCalls {
		funcRef := currMatch[0]
		funcName := currMatch[1]
		funcParam := currMatch[2]

		funcResult := cfg.ExecFunc(funcName, funcParam)
		funcReplRe, err := regexp.Compile(regexp.QuoteMeta(funcRef))
		if err != nil {
			return ""
		}
		processed = funcReplRe.ReplaceAllString(processed, funcResult)
	}

	return processed

}

func makeReplRe(keyword string, cfg *ParserConfig) (*regexp.Regexp, error) {
	quoted := cfg.ToQuoted()
	return regexp.Compile(fmt.Sprintf(`(%s%s%s)`, quoted.KeywordPrefix, keyword, quoted.KeywordSuffix))
}

func makeFuncRe(cfg *ParserConfig) (*regexp.Regexp, error) {
	quoted := cfg.ToQuoted()

	return regexp.Compile(fmt.Sprintf(`%s(\w+)\(([\w\s]+)\)`, quoted.FunctionDelimiter))
}
