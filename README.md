# go-spannulls

go-spannulls is a Go package to generate a minimum(NULL) [`structpb.Value`](https://pkg.go.dev/google.golang.org/protobuf/types/known/structpb#Value)
or [`spanner.GenericColumnValue`](https://pkg.go.dev/cloud.google.com/go/spanner#GenericColumnValue)
for a input [`spannerpb.Type`](https://pkg.go.dev/cloud.google.com/go/spanner@v1.72.0/apiv1/spannerpb#Type).