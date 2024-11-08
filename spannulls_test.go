package spannulls_test

import (
	sppb "cloud.google.com/go/spanner/apiv1/spannerpb"
	"github.com/apstndb/go-spannulls"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/structpb"
	"testing"
)

func TestNullGenericColumnValueFromType(t *testing.T) {
}

func codeToSimpleType(code sppb.TypeCode) *sppb.Type {
	return &sppb.Type{Code: code}
}

func codeToArrayType(code sppb.TypeCode) *sppb.Type {
	return &sppb.Type{Code: sppb.TypeCode_ARRAY,
		ArrayElementType: &sppb.Type{Code: code},
	}
}

func TestNullRawValueFromType_EnumerateSimpleTypes(t *testing.T) {
	for rawcode, typename := range sppb.TypeCode_name {
		if typename == "STRUCT" || typename == "ARRAY" {
			continue
		}
		got := spannulls.NullRawValueFromType(codeToSimpleType(sppb.TypeCode(rawcode)))
		t.Run(typename, func(t *testing.T) {
			if diff := cmp.Diff(structpb.NewNullValue(), got, protocmp.Transform()); diff != "" {
				t.Errorf("diff (-want, +got) = %v", diff)
			}
		})
	}
}

func TestNullRawValueFromType_ARRAY(t *testing.T) {
	input := codeToArrayType(sppb.TypeCode_STRING)
	want := structpb.NewNullValue()

	got := spannulls.NullRawValueFromType(input)

	if diff := cmp.Diff(want, got, protocmp.Transform()); diff != "" {
		t.Errorf("diff (-want, +got) = %v", diff)
	}
}

func TestNullRawValueFromType_STRUCT(t *testing.T) {
	input := &sppb.Type{Code: sppb.TypeCode_STRUCT, StructType: &sppb.StructType{Fields: []*sppb.StructType_Field{
		{Name: "int64", Type: codeToSimpleType(sppb.TypeCode_INT64)},
		{Name: "string", Type: codeToSimpleType(sppb.TypeCode_STRING)},
	}}}
	want := structpb.NewListValue(&structpb.ListValue{Values: []*structpb.Value{structpb.NewNullValue(), structpb.NewNullValue()}})

	got := spannulls.NullRawValueFromType(input)

	if diff := cmp.Diff(want, got, protocmp.Transform()); diff != "" {
		t.Errorf("diff (-want, +got) = %v", diff)
	}
}
