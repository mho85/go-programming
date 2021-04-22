package dog

import (
	"fmt"
	"testing"
)

var ageTest = 5
var expectedResult = 35

// TestYears tests the function Years
func TestYears(t *testing.T) {

	if Years(ageTest) != expectedResult {
		t.Error("Expected", expectedResult, "Got", Years(ageTest))
	}
}

func TestYearsTwo(t *testing.T) {
	if YearsTwo(ageTest) != expectedResult {
		t.Error("Expected", expectedResult, "Got", Years(ageTest))
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(ageTest)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(ageTest)
	}
}

func ExampleYears() {
	fmt.Println(Years(ageTest))
	// Output:
	// 35
}

func ExampleYearsTwo() {
	fmt.Println(Years(ageTest))
	// Output:
	// 35
}
