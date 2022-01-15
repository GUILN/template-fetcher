package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SliceContains_ReturnsExpectedValue_IndicatingIfTheContainingValueIsPresentOnSlice(t *testing.T) {
	cases := []struct {
		description    string
		slice          []interface{}
		conditionFunc  func(interface{}) bool
		expectedResult bool
	}{
		{
			description: "when slice contains element that satisfies condition function",
			slice: []interface{}{
				"banana",
				"apple",
				"pinaple",
			},
			conditionFunc: func(element interface{}) bool {
				return element == "banana"
			},
			expectedResult: true,
		},
		{
			description: "when slice DOES NOT contain element that satisfies condition function",
			slice: []interface{}{
				1, 2, 3, 4,
			},
			conditionFunc: func(element interface{}) bool {
				e, ok := element.(int)
				if !ok {
					panic(fmt.Errorf("cannot convert interface into int"))
				}
				return e > 5
			},
			expectedResult: false,
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			actualResult := SliceContains(c.slice, c.conditionFunc)
			assert.Equal(t, c.expectedResult, actualResult)
		})
	}
}

func Test_SliceReduce_ReturnsExpectedReducedResult_WhenPerformsSliceReduce(t *testing.T) {
	cases := []struct {
		description    string
		slice          []interface{}
		reducesFunc    func(interface{}, interface{}) interface{}
		expectedResult interface{}
	}{
		{
			description: "reduces fruit container into a single STRING",
			slice: []interface{}{
				"Banana",
				"Apple",
				"Pinaple",
			},
			reducesFunc: func(previousElement, currentElement interface{}) interface{} {
				curr, _ := currentElement.(string)
				prev, _ := previousElement.(string)
				return prev + curr
			},
			expectedResult: "BananaApplePinaple",
		},
		{
			description: "reduces fruit container into a single INT",
			slice: []interface{}{
				1,
				2,
				3,
				4,
			},
			reducesFunc: func(previousElement, currentElement interface{}) interface{} {
				curr, _ := currentElement.(int)
				prev, _ := previousElement.(int)
				return prev + curr
			},
			expectedResult: 10,
		},
		{
			description: "reduces correctly slice of one element STRING",
			slice: []interface{}{
				"Banana",
			},
			reducesFunc: func(previousElement, currentElement interface{}) interface{} {
				curr, _ := currentElement.(string)
				prev, _ := previousElement.(string)
				return prev + curr
			},
			expectedResult: "Banana",
		},
		{
			description: "reduces correctly slice of one element INT",
			slice: []interface{}{
				1,
			},
			reducesFunc: func(previousElement, currentElement interface{}) interface{} {
				curr, _ := currentElement.(int)
				prev, _ := previousElement.(int)
				return prev + curr
			},
			expectedResult: 1,
		},
		{
			description: "reduces COMPLEX TYPES",
			slice: []interface{}{
				&myDirType{Name: "nodeRepo", Path: "asdf/asdf"},
				&myDirType{Name: "nodeRepo", Path: "asdf/asdf"},
				&myDirType{Name: "nodeRepo", Path: "asdf/asdf"},
			},
			reducesFunc: func(previousElement, currentElement interface{}) interface{} {
				curr, _ := currentElement.(myDirType)
				prev, _ := previousElement.(myDirType)
				return prev.Name + curr.Name
			},
			expectedResult: "nodeReponodeReponodeRepo",
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			actualResult := SliceReduce(c.slice, c.reducesFunc)
			assert.Equal(t, c.expectedResult, actualResult)
		})
	}
}

type myDirType struct {
	Name string
	Path string
	Type string
}
