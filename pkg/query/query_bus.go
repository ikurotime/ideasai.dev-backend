package pkg

type QueryBus interface {
	Register(handler QueryHandler)
	Execute(query Query) (interface{}, error)
}
