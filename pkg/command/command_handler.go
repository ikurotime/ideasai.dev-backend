package pkg

import "reflect"

type CommandHandler interface {
	Handle(command Command) (QueryReponse error)
	CommandType() reflect.Type
}
