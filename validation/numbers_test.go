package validation

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsIntStr(t *testing.T) {
	assert.Equal(t, true, IsIntStr("+3"))
	assert.Equal(t, true, IsIntStr("-3"))
	assert.Equal(t, false, IsIntStr("3."))
	assert.Equal(t, false, IsIntStr("abc"))
}

func TestIsCreditCard(t *testing.T) {
	assert.Equal(t, true, IsCreditCard("4111111111111111"))
	assert.Equal(t, false, IsCreditCard("123456"))
}

func TestIsZeroValue(t *testing.T) {
	var (
		zeroPtr       *string
		zeroSlice     []int
		zeroFunc      func() string
		zeroMap       map[string]string
		nilInterface  interface{}
		zeroInterface fmt.Formatter
	)
	zeroValues := []interface{}{
		nil,
		false,
		0,
		int8(0),
		int16(0),
		int32(0),
		int64(0),
		uint(0),
		uint8(0),
		uint16(0),
		uint32(0),
		uint64(0),

		0.0,
		float32(0.0),

		"",

		// func
		zeroFunc,

		// array / slice
		[0]int{},
		zeroSlice,

		// map
		zeroMap,

		// interface
		nilInterface,
		zeroInterface,

		// pointer
		zeroPtr,

		// struct
		time.Time{},
	}

	for _, value := range zeroValues {
		assert.Equal(t, true, IsZeroValue(value))
	}

	var nonZeroInterface fmt.Stringer = time.Now()

	nonZeroValues := []interface{}{
		// bool
		true,

		// int
		1,
		int8(1),
		int16(1),
		int32(1),
		int64(1),
		uint8(1),
		uint16(1),
		uint32(1),
		uint64(1),

		// float
		1.0,
		float32(1.0),

		// string
		"test",

		// func
		time.Now,

		// array / slice
		[]int{},
		[]int{42},
		[1]int{42},

		// map
		make(map[string]string, 1),

		// interface
		nonZeroInterface,

		// pointer
		&nonZeroInterface,

		// struct
		time.Now(),
	}

	for _, value := range nonZeroValues {
		assert.Equal(t, false, IsZeroValue(value))
	}
}
