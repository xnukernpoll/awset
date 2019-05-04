package awset

import "testing"


func TestAWSET(t *testing.T) {

	a, b := New("a"), New("b")

	a.Insert("world")
	b.Insert("world")

	b.Remove("world")


	if !a.Contains("world") {
		t.Error("insert operation failed")
	}

	if b.Contains("world") {
		t.Error("Remove Operation failed")
	}
	

	b.Merge(&a)

	if !b.Contains("world") {
		t.Error("expected element to exist")
	}



	a.Remove("world")
	b.Merge(&a)

	
	if b.Contains("world") {
		t.Error("expected element to not exists")
	}
	


	

	
}
