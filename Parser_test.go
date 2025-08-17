package mare

import (
	"strings"
	"testing"
)

func Test_Functions(t *testing.T) {
	var cut *ParserConfig = &ParserConfig{
		KeywordPrefix:     "[",
		KeywordSuffix:     "]",
		FunctionDelimiter: "§",
		functions:         make(map[string]ParserFunction),
	}

	cut.RegisterFunc("reverse", func(s string) string {
		var retVal strings.Builder = strings.Builder{}

		for i := len(s) - 1; i >= 0; i-- {
			retVal.WriteByte(s[i])
		}

		return retVal.String()
	})

	defInput := "WillBeReversed"
	expOutput := "desreveReBlliW"

	actOutput := cut.ExecFunc("rEvErSE", defInput)

	if expOutput != actOutput {
		t.Fatalf("expected %s but got %s", expOutput, actOutput)
	}
}

func Test_RegisterFuncFail(t *testing.T) {
	cut := DefaultConfig()
	err1 := cut.RegisterFunc("Doublette", func(s string) string { return s })
	err2 := cut.RegisterFunc("douBlette", func(s string) string { return s })

	if err1 != nil {
		t.Fatal("first registration should have succeeded")
	}

	if err2 == nil {
		t.Fatal("second registration should have failed")
	}
}

func Test_NonEvalFunction(t *testing.T) {
	cut := DefaultConfig()

	defInput := "Will not be touched"
	expOutput := "Will not be touched"

	actResult := cut.ExecFunc("UnregisteredFunction", defInput)

	if expOutput != actResult {
		t.Fatalf("expected %s but got %s", expOutput, actResult)
	}
}
