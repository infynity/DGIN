package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomething(t *testing.T) {

	//str :=123
	//assert.NotEqual(t, str, "")
	var err error
	a := make([]int,0)

	fmt.Println(a)
	//err = errors.New("sd")
	assert.True(t, true, "True is true!")
	//if err!=nil{
		assert.Nil(t, err)
	//}

}