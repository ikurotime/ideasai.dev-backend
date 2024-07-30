package pkg

type CommandBus interface {
	Register(handler CommandHandler)
	Execute(command Command) CommandResponse
}
