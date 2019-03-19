package mybiginteger_test

import (
	. "number/mybiginteger"
	"reflect"
	"testing"
)

func TestConvertStringToBytes(t *testing.T) {
	type TestData struct {
		expectedResult []byte
		inputData      string
	}
	testData := []TestData{
		{[]byte{1, 2, 3, 5}, "1235"},
		{[]byte{253, 1, 2, 3, 5}, "-1235"},
		{[]byte{1, 2, 3, 254, 5}, "123.5"},
		{[]byte{253, 1, 2, 3, 254, 5}, "-123.5"},
		{[]byte(nil), "a1235"},
		{[]byte(nil), "12$35"},
		{[]byte(nil), "1235@"},
	}

	for _, testCase := range testData {
		actualBytes := ConvertStringToBytes(testCase.inputData)
		if !reflect.DeepEqual(testCase.expectedResult, actualBytes) {
			t.Errorf("expect %v but it got %v", testCase.expectedResult, actualBytes)
		}
	}
}