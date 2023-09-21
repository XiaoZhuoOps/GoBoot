package goExample

import (
	"fmt"
	"strings"
	"testing"
)

func TestGenSomething(t *testing.T) {
	template := `
	if $Parsed := signalDiagnosticLabelsParsed.Get("*"); $Parsed.Exists() {
    $Formatted := &entity.DiagnosticLabel{}
    $Formatted.Label = $Parsed.Get("label").String()
    if $AbnormalTypesParsed := $Parsed.Get("abnormal_types"); $AbnormalTypesParsed.Exists() {
        $AbnormalTypesFormatted := make([]string, 0, 5)
        for _, $AbnormalTypeParsed := range $AbnormalTypesParsed.Array() {
            $AbnormalTypesFormatted = append($AbnormalTypesFormatted, $AbnormalTypeParsed.String())
        }
        $Formatted.AbnormalTypes = $AbnormalTypesFormatted
    }
    if $SuggestedValuesParsed := $Parsed.Get("suggested_values"); $SuggestedValuesParsed.Exists() {
        $SuggestedValuesFormatted := make([]string, 0, 5)
        for _, $SuggestedValue := range $SuggestedValuesParsed.Array() {
            $SuggestedValuesFormatted = append($SuggestedValuesFormatted, $SuggestedValue.String())
        }
        $Formatted.SuggestedValues = $SuggestedValuesFormatted
    }
    signalDiagnosticLabelsFormatted.% = $Formatted
	}`
	bs := []string{"hashed_email", "raw_phone", "raw_auto_phone", "hashed_phone"}
	as := []string{"hashedEmail", "rawPhone", "rawAutoPhone", "hashedPhone"}
	cs := []string{"HashedEmail", "RawPhone", "RawAutoPhone", "HashedPhone"}
	for i, _ := range as {
		templateA := strings.ReplaceAll(template, "$", as[i])
		templateB := strings.ReplaceAll(templateA, "*", bs[i])
		templateC := strings.ReplaceAll(templateB, "%", cs[i])
		fmt.Println(templateC)
	}
}
