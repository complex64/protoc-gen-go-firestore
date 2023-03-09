package converter_test

import (
	"testing"
	"time"

	"github.com/complex64/protoc-gen-go-firestore/v2/internal/gen/test/converter"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestConverters_ToFirestore(t *testing.T) {
	now := time.Now().UTC()
	city := &converter.City{
		StringField:            "string",
		BoolField:              true,
		BytesField:             []byte("bytes"),
		Int32Field:             1,
		Int64Field:             2,
		Uint32Field:            3,
		Sint32Field:            4,
		Sint64Field:            5,
		FloatField:             6,
		DoubleField:            7,
		TimestampField:         timestamppb.New(now),
		RepeatedStringField:    []string{""},
		RepeatedBoolField:      []bool{true},
		RepeatedBytesField:     [][]byte{[]byte("")},
		RepeatedInt32Field:     []int32{0},
		RepeatedInt64Field:     []int64{0},
		RepeatedUint32Field:    []uint32{0},
		RepeatedSint32Field:    []int32{0},
		RepeatedSint64Field:    []int64{0},
		RepeatedFloatField:     []float32{0},
		RepeatedDoubleField:    []float64{0},
		RepeatedTimestampField: []*timestamppb.Timestamp{timestamppb.New(now)},
	}
	fs, err := city.ToFirestoreMap()
	require.NoError(t, err)

	require.Equal(t, "string", fs["stringField"])
	require.Equal(t, true, fs["boolField"])
	require.Equal(t, "Ynl0ZXM=", fs["bytesField"]) // base64 for bytes
	require.Equal(t, 1.0, fs["int32Field"])
	require.Equal(t, "2", fs["int64Field"])
	require.Equal(t, 3.0, fs["uint32Field"])
	require.Equal(t, 4.0, fs["sint32Field"])
	require.Equal(t, "5", fs["sint64Field"])
	require.Equal(t, 6.0, fs["floatField"])
	require.Equal(t, 7.0, fs["doubleField"])
}
