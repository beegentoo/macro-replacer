# macro-replacer
Replace placeholders in Texts based on a set of replacements

## Intro
Mare is a very simple utility to evaluate text based on replacement-maps and simple functions

## Usage
```go
// Taken from MacroReplacer_test.go

func Test_Complicated(t *testing.T) {
	var defReplacements map[string]any = make(map[string]any)
	defReplacements["VoucherType"] = "OrderConfirmation"
	defReplacements["SupplierName"] = "ACME Incorporated"

	targetDirectoryTemplate := "/home/user/docs/[VoucherType]/§firstletter([SupplierName])/[SupplierName]/"
	targetDirectoryExpected := "/home/user/docs/OrderConfirmation/A/ACME Incorporated/"

	actResult := Process(targetDirectoryTemplate, defReplacements)

	if targetDirectoryExpected != actResult {
		t.Fatalf("Expected %s but got %s", targetDirectoryExpected, actResult)
	}
}
```
## Available predefined functions
* `upper`: Transform string to upper case
* `lower`: Transform string to lower case
* `firstletter`: Deliver the first letter of the string

## Register custom functions
A predefined alias type `ParserFunction` helps registering custom functions using the `RegisterFunc(funcName string, fn ParserFunction)` method:

```go
// Create a simple parser config
var cut *ParserConfig = &ParserConfig{
	KeywordPrefix:     "[",
	KeywordSuffix:     "]",
	FunctionDelimiter: "§",
	functions:         make(map[string]ParserFunction),
}

// Register a function
cut.RegisterFunc("reverse", func(s string) string {
	var retVal strings.Builder = strings.Builder{}

	for i := len(s) - 1; i >= 0; i-- {
		retVal.WriteByte(s[i])
	}

	return retVal.String()
})

// Usage in the template would then be used §Reverse[This will be reversed]

```


