package customobject_test

import (
	"testing"
	"time"

	"github.com/complex64/protoc-gen-go-firestore/v2/internal/gen/test/customobject"
	"github.com/complex64/protoc-gen-go-firestore/v2/internal/require"
)

func TestCustomObjectStructTypeGeneration(t *testing.T) {
	city := new(customobject.FirestoreCity)

	city.StringField = ""
	city.RepeatedStringField = []string{}

	city.BoolField = true
	city.RepeatedBoolField = []bool{}

	city.BytesField = []byte("")
	city.RepeatedBytesField = [][]byte{}

	city.Int32Field = 0
	city.RepeatedInt32Field = []int32{}

	city.Int64Field = 0
	city.RepeatedInt64Field = []int64{}

	city.Uint32Field = 0
	city.RepeatedUint32Field = []uint32{}

	city.Sint32Field = 0
	city.RepeatedSint32Field = []int32{}

	city.Sint64Field = 0
	city.RepeatedSint64Field = []int64{}

	city.FloatField = 0.0
	city.RepeatedFloatField = []float32{}

	city.DoubleField = 0.0
	city.RepeatedDoubleField = []float64{}

	city.TimestampField = time.Now()
	city.RepeatedTimestampField = []time.Time{}
}

func TestCustomObjectTags(t *testing.T) {
	city := new(customobject.FirestoreCity)

	require.StructFieldTags(t, city, "StringField", map[string]string{"firestore": "stringField,omitempty"})
	require.StructFieldTags(t, city, "RepeatedStringField", map[string]string{"firestore": "repeatedStringField,omitempty"})

	require.StructFieldTags(t, city, "BoolField", map[string]string{"firestore": "boolField,omitempty"})
	require.StructFieldTags(t, city, "RepeatedBoolField", map[string]string{"firestore": "repeatedBoolField,omitempty"})

	require.StructFieldTags(t, city, "BytesField", map[string]string{"firestore": "bytesField,omitempty"})
	require.StructFieldTags(t, city, "RepeatedBytesField", map[string]string{"firestore": "repeatedBytesField,omitempty"})

	require.StructFieldTags(t, city, "Int32Field", map[string]string{"firestore": "int32Field,omitempty"})
	require.StructFieldTags(t, city, "RepeatedInt32Field", map[string]string{"firestore": "repeatedInt32Field,omitempty"})

	require.StructFieldTags(t, city, "Int64Field", map[string]string{"firestore": "int64Field,omitempty"})
	require.StructFieldTags(t, city, "RepeatedInt64Field", map[string]string{"firestore": "repeatedInt64Field,omitempty"})

	require.StructFieldTags(t, city, "Uint32Field", map[string]string{"firestore": "uint32Field,omitempty"})
	require.StructFieldTags(t, city, "RepeatedUint32Field", map[string]string{"firestore": "repeatedUint32Field,omitempty"})

	require.StructFieldTags(t, city, "Sint32Field", map[string]string{"firestore": "sint32Field,omitempty"})
	require.StructFieldTags(t, city, "RepeatedSint32Field", map[string]string{"firestore": "repeatedSint32Field,omitempty"})

	require.StructFieldTags(t, city, "Sint64Field", map[string]string{"firestore": "sint64Field,omitempty"})
	require.StructFieldTags(t, city, "RepeatedSint64Field", map[string]string{"firestore": "repeatedSint64Field,omitempty"})

	require.StructFieldTags(t, city, "FloatField", map[string]string{"firestore": "floatField,omitempty"})
	require.StructFieldTags(t, city, "RepeatedFloatField", map[string]string{"firestore": "repeatedFloatField,omitempty"})

	require.StructFieldTags(t, city, "DoubleField", map[string]string{"firestore": "doubleField,omitempty"})
	require.StructFieldTags(t, city, "RepeatedDoubleField", map[string]string{"firestore": "repeatedDoubleField,omitempty"})

	require.StructFieldTags(t, city, "TimestampField", map[string]string{"firestore": "timestampField,omitempty"})
	require.StructFieldTags(t, city, "RepeatedTimestampField", map[string]string{"firestore": "repeatedTimestampField,omitempty"})

	require.StructFieldTags(t, city, "NamedStringField", map[string]string{"firestore": "custom_name,omitempty"})
	require.StructFieldTags(t, city, "IgnoredStringField", map[string]string{"firestore": "-"})
	require.StructFieldTags(t, city, "ServerTimestampField", map[string]string{"firestore": "serverTimestampField,serverTimestamp"})
}

func TestCustomObjectStructNestedTypeGeneration(t *testing.T) {
	_ = new(customobject.FirestoreMayor)
	_ = new(customobject.FirestoreAddress)
}
