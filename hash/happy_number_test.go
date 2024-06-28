package hash

import (
	"testing"
)

func TestIsHappy(t *testing.T){
    if (!IsHappy(19)){
        t.Error("19 is happy number")
        return
    }
    if (IsHappy(2)){
        t.Error("2 is not happy number")
        return
    }
   
 
}

