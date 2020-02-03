package protokit

import (
	"bytes"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	Name     string
	WireType WireType
	FieldId  float64
	TypeByte byte
	Data     []byte
	Result   uint64
}

func (tc *testCase) getTypeByte() byte {
	return byte(math.Pow(2, lenWireType) + tc.FieldId)
}


// TODO this cases cause panics - WireStartGroup, WireEndGroup
var testCases = []testCase{
	{
		Name:     "WireVarint",
		WireType: WireVarint,
		FieldId:  0,
		Data:     []byte{123},
		Result:   uint64(123),
	},
	{
		Name:     "WireFixed64",
		WireType: WireFixed64,
		FieldId:  1,
		Data:     []byte{1, 1, 1, 1, 1, 1, 1, 1, 1},
		Result:   uint64(0x101010101010101),
	},
	{
		Name:     "WireBytes",
		WireType: WireBytes,
		FieldId:  2,
		Data:     []byte{123},
		Result:   uint64(123),
	},
	{
		Name:     "WireFixed32",
		WireType: WireFixed32,
		FieldId:  5,
		Data:     []byte{1, 1, 1, 1, 1,},
		Result:   uint64(0x1010101),
	},
}

func TestSafeWireTag(t *testing.T) {
	var test uint64
	test = math.MaxUint32 + 1
	wireTag, err := SafeWireTag(test)
	require.Error(t, err)
	require.Equal(t, WireTag(0), wireTag)

	test = uint64(math.Pow(2, lenWireType-1))
	wireTag, err = SafeWireTag(test)
	require.Error(t, err)
	require.Equal(t, WireTag(0), wireTag)

	test = uint64(math.Pow(2, lenWireType))
	wireTag, err = SafeWireTag(test)
	require.NoError(t, err)
	require.Equal(t, WireTag(test), wireTag)
}

func TestWireTypeDecodeFromAnotherType(t *testing.T) {
	buf := bytes.Buffer{}
	buf.Write([]byte{byte(math.Pow(2, lenWireType)), 123})

	// try to decode WireVariant to WireBytes
	_, _, err := WireBytes.DecodeFrom(&buf)
	require.Error(t, err)

	// try to decode WireVariant to WireFixed64
	_, _, err = WireFixed64.DecodeFrom(&buf)
	require.Error(t, err)

	// try to decode WireVariant to WireFixed32
	_, _, err = WireFixed32.DecodeFrom(&buf)
	require.Error(t, err)
}

func TestWireTypeDecodeFrom(t *testing.T) {
	buf := bytes.Buffer{}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			buf.Reset()
			buf.Write(append([]byte{test.getTypeByte()}, test.Data...))

			wireTag, data, err := test.WireType.DecodeFrom(&buf)
			require.NoError(t, err)
			require.Equal(t, test.Result, data)
			require.Equal(t, test.WireType, wireTag.Type())
		})
	}
}

func TestWireTagDecodeFromAnotherType(t *testing.T) {
	buf := bytes.Buffer{}
	buf.Write([]byte{byte(math.Pow(2, lenWireType)), 123})

	// try to decode WireVariant to WireBytes
	_, err := WireBytes.Tag(1).DecodeFrom(&buf)
	require.Error(t, err)

	// try to decode WireVariant to WireFixed64
	_, err = WireFixed64.Tag(1).DecodeFrom(&buf)
	require.Error(t, err)

	// try to decode WireVariant to WireFixed32
	_, err = WireFixed32.Tag(1).DecodeFrom(&buf)
	require.Error(t, err)
}

func TestWireTagDecodeFrom(t *testing.T) {
	buf := bytes.Buffer{}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			buf.Reset()
			buf.Write(append([]byte{test.getTypeByte()}, test.Data...))

			tag := test.WireType.Tag(1)
			data, err := tag.DecodeFrom(&buf)
			require.NoError(t, err)
			require.Equal(t, test.Result, data)
		})
	}
}

func TestEnsureFixedFieldSize(t *testing.T) {
	tag := WireFixed64.Tag(1)
	newTag := tag.EnsureFixedFieldSize(9)
	require.Equal(t, tag, newTag)

	tag = WireFixed32.Tag(1)
	newTag = tag.EnsureFixedFieldSize(5)
	require.Equal(t, tag, newTag)
}

func TestEnsureFixedFieldSizeFailedForWireVarint(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("EnsureFixedFieldSize did not panic for WireVarint")
		}
	}()

	tag := WireVarint.Tag(1)
	tag.EnsureFixedFieldSize(12)
}

func TestEnsureFixedFieldSizeFailedForWireBytes(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("EnsureFixedFieldSize did not panic for WireBytes")
		}
	}()

	tag := WireBytes.Tag(1)
	tag.EnsureFixedFieldSize(12)
}

func TestWireTagMustEncodeTo(t *testing.T) {
	buf := bytes.Buffer{}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			buf.Reset()
			buf.Write(append([]byte{test.getTypeByte()}, test.Data...))

			tag := test.WireType.Tag(1)
			tag.MustEncodeTo(&buf, 1234)
		})
	}

}
