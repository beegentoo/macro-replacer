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
## Available functions
* `upper`: Transform string to upper case
* `lower`: Transform string to lower case
* `firstletter`: Deliver the first letter of the string
