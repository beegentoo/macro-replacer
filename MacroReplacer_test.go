package mare

import (
	"fmt"
	"testing"
)

func Test_ProcessReplace(t *testing.T) {
	keywords := make(map[string]any)

	keywords["Word_de"] = "Wort"
	keywords["Word_fr"] = "mot"

	origString := "Word is [Word_de] in german and [Word_fr] in french"
	expResult := "Word is Wort in german and mot in french"

	processedString := Process(origString, keywords)

	if processedString != expResult {
		t.Fatal(fmt.Errorf("Process() did not return the expected result %s but %s", expResult, processedString))
	}
}

func Test_ProcessFunction(t *testing.T) {
	origString := "The first letter of sausage is §firstletter(Sausage)"
	expResult := "The first letter of sausage is S"

	actResult := Process(origString, map[string]any{})

	if expResult != actResult {
		t.Fatal(fmt.Errorf("Process did not return the expected result for function eval"))
	}
}

func Test_Functions(t *testing.T) {
	origString := "§Upper(HansDampf)"
	expResult := "HANSDAMPF"

	actResult := Process(origString, map[string]any{})

	if expResult != actResult {
		t.Fatalf("Expected %s but got %s", expResult, actResult)
	}
}

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

func Test_execFunc(t *testing.T) {
	if execFunc("upper", "aAa") != "AAA" {
		t.Fail()
	}

	if execFunc("firstletter", "Zebra Flotilla") != "Z" {
		t.Fail()
	}

	if execFunc("lower", "Xxx") != "xxx" {
		t.Fail()
	}

	if execFunc("nonExistant", "Dummy") != "Dummy" {
		t.Fail()
	}
}
