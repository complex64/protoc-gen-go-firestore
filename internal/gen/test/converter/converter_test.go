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
		StringField:            "",
		BoolField:              false,
		BytesField:             []byte(""),
		Int32Field:             0,
		Int64Field:             0,
		Uint32Field:            0,
		Sint32Field:            0,
		Sint64Field:            0,
		FloatField:             0,
		DoubleField:            0,
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
	fs, err := city.ToFirestore()
	require.NoError(t, err)

	require.Equal(t, city.StringField, fs.StringField)
	require.Equal(t, city.BoolField, fs.BoolField)
	require.Equal(t, city.BytesField, fs.BytesField)
	require.Equal(t, city.Int32Field, fs.Int32Field)
	require.Equal(t, city.Int64Field, fs.Int64Field)
	require.Equal(t, city.Uint32Field, fs.Uint32Field)
	require.Equal(t, city.Sint32Field, fs.Sint32Field)
	require.Equal(t, city.Sint64Field, fs.Sint64Field)
	require.Equal(t, city.FloatField, fs.FloatField)
	require.Equal(t, city.DoubleField, fs.DoubleField)
	require.Equal(t, city.TimestampField.AsTime(), fs.TimestampField)
	require.Equal(t, city.RepeatedStringField, fs.RepeatedStringField)
	require.Equal(t, city.RepeatedBoolField, fs.RepeatedBoolField)
	require.Equal(t, city.RepeatedBytesField, fs.RepeatedBytesField)
	require.Equal(t, city.RepeatedInt32Field, fs.RepeatedInt32Field)
	require.Equal(t, city.RepeatedInt64Field, fs.RepeatedInt64Field)
	require.Equal(t, city.RepeatedUint32Field, fs.RepeatedUint32Field)
	require.Equal(t, city.RepeatedSint32Field, fs.RepeatedSint32Field)
	require.Equal(t, city.RepeatedSint64Field, fs.RepeatedSint64Field)
	require.Equal(t, city.RepeatedFloatField, fs.RepeatedFloatField)
	require.Equal(t, city.RepeatedDoubleField, fs.RepeatedDoubleField)

	require.Len(t, fs.RepeatedTimestampField, 1)
	require.Equal(t, now, fs.RepeatedTimestampField[0])
}

func TestNestedConverters_ToFirestore(t *testing.T) {
	city := &converter.City{
		Mayor: &converter.Mayor{
			Name:    "Mayor",
			Address: &converter.Address{Value: "Address"},
		},
		MyNestedField: &converter.City_MyNestedMessage{
			Name: "Nested",
		},
	}
	fs, err := city.ToFirestore()
	require.NoError(t, err)

	require.NotNil(t, fs.Mayor)
	require.Equal(t, city.Mayor.Name, fs.Mayor.Name)
	require.Equal(t, city.Mayor.Address.Value, fs.Mayor.Address.Value)
	require.Equal(t, city.MyNestedField.Name, fs.MyNestedField.Name)
}
