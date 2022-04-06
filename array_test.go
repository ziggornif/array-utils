package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	firstname string
	lastname  string
}

var strArr = NewArray([]string{"a", "b", "c"})
var intArr = NewArray([]int{1, 2, 3})
var structArr = NewArray([]Person{
	{"Murphy", "Dunne"},
	{"Alexa", "Wilcox"},
	{"Esha", "Parker"},
})

func TestNewArray(t *testing.T) {
	nativeArr := []int{1, 2, 3}
	arr := NewArray(nativeArr)
	assert.Equal(t, nativeArr, arr.ToNative())
}

func TestArray_FindString(t *testing.T) {
	found := strArr.Find(func(val string) bool {
		return val == "b"
	})
	assert.Equal(t, "b", found)
}

func TestArray_FindStringNotFound(t *testing.T) {
	found := strArr.Find(func(val string) bool {
		return val == "d"
	})
	assert.Equal(t, "", found)
}

func TestArray_FindInt(t *testing.T) {
	found := intArr.Find(func(val int) bool {
		return val == 2
	})
	assert.Equal(t, 2, found)
}

func TestArray_FindIntNotFound(t *testing.T) {
	found := intArr.Find(func(val int) bool {
		return val == 4
	})
	assert.Equal(t, 0, found)
}

func TestArray_FindStruct(t *testing.T) {
	found := structArr.Find(func(val Person) bool {
		return val.lastname == "Wilcox"
	})
	assert.Equal(t, Person{"Alexa", "Wilcox"}, found)
}

func TestArray_FindStructNotFound(t *testing.T) {
	found := structArr.Find(func(val Person) bool {
		return val.lastname == "Doe"
	})
	assert.Equal(t, Person{}, found)
}

func TestArray_FindIndexString(t *testing.T) {
	found := strArr.FindIndex(func(val string) bool {
		return val == "b"
	})
	assert.Equal(t, 1, found)
}

func TestArray_FindIndexStringNotFound(t *testing.T) {
	found := strArr.FindIndex(func(val string) bool {
		return val == "d"
	})
	assert.Equal(t, -1, found)
}

func TestArray_FindIndexInt(t *testing.T) {
	found := intArr.FindIndex(func(val int) bool {
		return val == 2
	})
	assert.Equal(t, 1, found)
}

func TestArray_FindIndexNotFound(t *testing.T) {
	found := intArr.FindIndex(func(val int) bool {
		return val == 4
	})
	assert.Equal(t, -1, found)
}

func TestArray_FindIndexStruct(t *testing.T) {
	found := structArr.FindIndex(func(val Person) bool {
		return val.lastname == "Wilcox"
	})
	assert.Equal(t, 1, found)
}

func TestArray_FindIndexStructNotFound(t *testing.T) {
	found := structArr.FindIndex(func(val Person) bool {
		return val.lastname == "Doe"
	})
	assert.Equal(t, -1, found)
}

func TestArray_FilterString(t *testing.T) {
	filtered := strArr.Filter(func(val string) bool {
		return val != "b"
	})
	assert.Equal(t, []string{"a", "c"}, filtered.ToNative())
}

func TestArray_FilterInt(t *testing.T) {
	filtered := intArr.Filter(func(val int) bool {
		return val != 2
	})
	assert.Equal(t, []int{1, 3}, filtered.ToNative())
}

func TestArray_FilterStruct(t *testing.T) {
	filtered := structArr.Filter(func(val Person) bool {
		return val.lastname != "Wilcox"
	})
	assert.Equal(t, []Person{
		{"Murphy", "Dunne"},
		{"Esha", "Parker"},
	}, filtered.ToNative())
}

func TestArray_MapString(t *testing.T) {
	newarr := strArr.Map(func(val string) string {
		return val + val
	})
	assert.Equal(t, []string{"aa", "bb", "cc"}, newarr.ToNative())
}

func TestArray_MapInt(t *testing.T) {
	filtered := intArr.Map(func(val int) int {
		return val * 2
	})
	assert.Equal(t, []int{2, 4, 6}, filtered.ToNative())
}

func TestArray_MapStruct(t *testing.T) {
	filtered := structArr.Map(func(val Person) Person {
		tempLast := val.lastname
		val.lastname = val.firstname
		val.firstname = tempLast
		return val
	})
	assert.Equal(t, []Person{
		{"Dunne", "Murphy"},
		{"Wilcox", "Alexa"},
		{"Parker", "Esha"},
	}, filtered.ToNative())
}

func TestArray_ForEachString(t *testing.T) {
	var counter = 0
	strArr.ForEach(func(val string) {
		assert.Equal(t, strArr.ToNative()[counter], val)
		counter += 1
	})
}

func TestArray_ForEachInt(t *testing.T) {
	var counter = 0
	intArr.ForEach(func(val int) {
		assert.Equal(t, intArr.ToNative()[counter], val)
		counter += 1
	})
}

func TestArray_ForEachStruct(t *testing.T) {
	var counter = 0
	structArr.ForEach(func(val Person) {
		assert.Equal(t, structArr.ToNative()[counter].lastname, val.lastname)
		counter += 1
	})
}

func TestArray_MapChaining(t *testing.T) {
	result := intArr.Map(func(val int) int {
		return val * 3
	}).Filter(func(val int) bool {
		return val%2 == 0
	})
	assert.Equal(t, []int{6}, result.ToNative())
}
