package unsafecast_test

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
	unsafecast "github.com/zergon321/unsafe-cast"
)

type TestStruct8 struct {
	X float32
	Y float32
}

type TestStruct12 struct {
	X float32
	Y float32
	Z float32
}

type TestStruct12Int struct {
	X int32
	Y int32
	Z int32
}

func TestConvertSliceFromGreater(t *testing.T) {
	slice := []float32{16.34, 29.51}
	result, err := unsafecast.ConvertSlice[float32, byte](slice)

	assert.Nil(t, err)
	assert.Len(t, result, 8)
}

func TestConvertSliceFromLesser(t *testing.T) {
	slice := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	result, err := unsafecast.ConvertSlice[byte, int32](slice)

	assert.Nil(t, err)
	assert.Len(t, result, 2)
}

func TestConvertSliceFromGreater2to3(t *testing.T) {
	slice := []TestStruct12{
		{X: 12.3, Y: 10.2, Z: 7.8},
		{X: 7.8, Y: 12.3, Z: 10.2},
		{X: 10.2, Y: 7.8, Z: 12.3},
		{X: 33.14, Y: 47.12, Z: 48.32},
	}
	result, err := unsafecast.ConvertSlice[TestStruct12, TestStruct8](slice)

	assert.Nil(t, err)
	assert.Len(t, result, 6)
}

func TestConvertSliceFromLesser2to3(t *testing.T) {
	slice := []TestStruct8{
		{X: 12.3, Y: 10.2},
		{X: 7.8, Y: 12.3},
		{X: 10.2, Y: 7.8},
	}
	result, err := unsafecast.ConvertSlice[TestStruct8, TestStruct12](slice)

	assert.Nil(t, err)
	assert.Len(t, result, 2)
}

func TestConvertOne(t *testing.T) {
	value, err := unsafecast.ConvertOne[TestStruct12Int, TestStruct12](
		TestStruct12Int{X: 1, Y: 2, Z: 3})
	_ = value

	assert.Nil(t, err)
}

func TestConvertOnePointer(t *testing.T) {
	value, err := unsafecast.ConvertOnePointer[TestStruct12Int, TestStruct12](
		&TestStruct12Int{X: 1, Y: 2, Z: 3})
	_ = value

	assert.Nil(t, err)
}

func BenchmarkConvertSlice(b *testing.B) {
	slice := []byte{1, 2, 3, 4, 5, 6, 7, 8}

	for i := 0; i < b.N; i++ {
		unsafecast.ConvertSlice[byte, int32](slice)
	}
}

func BenchmarkConvertOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unsafecast.ConvertOne[TestStruct12Int, TestStruct12](
			TestStruct12Int{X: 1, Y: 2, Z: 3})
	}
}

func BenchmarkConvertOnePointer(b *testing.B) {
	value := &TestStruct12Int{X: 1, Y: 2, Z: 3}

	for i := 0; i < b.N; i++ {
		unsafecast.ConvertOnePointer[TestStruct12Int, TestStruct12](
			value)
	}
}

func BenchmarkConvertStructToByteSlice(b *testing.B) {
	value := &TestStruct12Int{X: 1, Y: 2, Z: 3}

	for i := 0; i < b.N; i++ {
		pointer, _ := unsafecast.ConvertOnePointer[TestStruct12Int,
			[unsafe.Sizeof(TestStruct12Int{})]byte](value)
		data := pointer[:]
		_ = data
	}
}

func BenchmarkConvertStructFromSlice(b *testing.B) {
	slice := []float32{13.64, 72.53, 12.95}

	for i := 0; i < b.N; i++ {
		unsafecast.ConvertOneFromSlice[float32, TestStruct12Int](slice)
	}
}

func BenchmarkConvertStructFromSlicePointer(b *testing.B) {
	slice := []float32{13.64, 72.53, 12.95}

	for i := 0; i < b.N; i++ {
		unsafecast.ConvertOneFromSlicePointer[float32, TestStruct12Int](slice)
	}
}
