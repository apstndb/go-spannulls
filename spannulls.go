package spannulls

import (
	"slices"

	"cloud.google.com/go/spanner"
	sppb "cloud.google.com/go/spanner/apiv1/spannerpb"
	"google.golang.org/protobuf/types/known/structpb"
)

// NullGenericColumnValueFromType generates a minimum valid spanner.GenericColumnValue for the input type.
func NullGenericColumnValueFromType(typ *sppb.Type) spanner.GenericColumnValue {
	return spanner.GenericColumnValue{Type: typ, Value: NullRawValueFromType(typ)}
}

// NullRawValueFromType generates a minimum valid value for the input type.
func NullRawValueFromType(typ *sppb.Type) *structpb.Value {
	switch typ.GetCode() {
	// Only STRUCT needs a non-null value.
	case sppb.TypeCode_STRUCT:
		return structpb.NewListValue(&structpb.ListValue{Values: slices.Repeat(
			[]*structpb.Value{structpb.NewNullValue()},
			len(typ.StructType.GetFields()))})
	default:
		return structpb.NewNullValue()
	}
}
