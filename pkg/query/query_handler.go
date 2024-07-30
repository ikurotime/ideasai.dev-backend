package pkg

import "reflect"

type QueryHandler interface {
	Handle(query Query) (QueryReponse error)
	QueryType() reflect.Type
}
