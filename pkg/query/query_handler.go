package pkg

import (
	"reflect"
)

type QueryHandler interface {
	Handle(query Query) (QueryResponse, error)
	QueryType() reflect.Type
}
