package helpers

import "testing"

//func TestSum(t *testing.T){
//	result := Sum(10, 10)

//	if result != 20 {
//		panic("Result should be 20")
//	}
//}

//func TestFailSum(t *testing.T){
//	result := Sum (10, 10)
//	if result != 40 {
//		t.Fail()
//	}

//	fmt.Println("TestFailSum Eksekusi Terhenti")
//}

func TestFailSum(t *testing.T){
	result := Sum (10, 10)
	if result != 40 {
		t.Error("Result has to be 40")
	}

	fmt.Println("TestFailSum Eksekusi Terhenti")
}
cd
func TestSum(t *testing.T){
	result := Sum(10, 10)
	if result != 20{
		t.FailNow()
	}

	fmt.Println("TestSum Eksekusi Terhenti")
}