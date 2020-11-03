package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	//Test for empty arguments
	emptyresult := hello("") //should return Hello K8's!

	if emptyresult != "Hello K8's!" {

		t.Errorf("hello(\"\") failed , expected %v, got %v", "Hello K8's!", emptyresult)
	} else {
		t.Logf("hello(\"\") success , expected %v, got %v", "Hello K8's", emptyresult)
	}

	//Test for filled arguments

	result := hello("Ritesh")

	if result != "Hello Ritesh!" {
		t.Errorf("hello(\"\") failed , expected %v, got %v", "Hello Ritesh!", result)
	} else {
		t.Logf("hello(\"\") success , expected %v, got %v", "Hello Ritesh!", result)
	}

}
