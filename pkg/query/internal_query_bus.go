package pkg

import (
	"context"
	"fmt"

	"reflect"
	"sync"

	"go.uber.org/zap"
)

type InternalQueryBus struct {
	handlers sync.Map
	log      *zap.Logger
}

func NewInternalQueryBus(handlers []QueryHandler, log *zap.Logger) QueryBus {
	bus := &InternalQueryBus{
		log: log,
	}

	for _, handler := range handlers {
		bus.Register(handler)
	}

	return bus
}

func (bus *InternalQueryBus) Register(handler QueryHandler) {
	queryType := handler.QueryType()
	handlerType := reflect.TypeOf(handler)

	bus.handlers.Store(queryType, handler)
	bus.log.Debug(fmt.Sprintf("Registered %s for %s", handlerType.String(), queryType.String()))
}

func (bus *InternalQueryBus) Execute(query Query) (interface{}, error) {
	queryType := reflect.TypeOf(query)
	bus.log.Debug(fmt.Sprintf("Executing query %s", queryType.String()))
	//print available handlers
	bus.handlers.Range(func(key, value interface{}) bool {
		bus.log.Debug(fmt.Sprintf("Available handler: %s", key))
		return true
	})
	handler, ok := bus.handlers.Load(queryType)
	if !ok {
		// Try non-pointer type
		nonPtrType := queryType.Elem()
		handler, ok = bus.handlers.Load(nonPtrType)
		bus.log.Debug(fmt.Sprintf("Non-pointer type: %s", nonPtrType.String()))

		if !ok {
			err := fmt.Errorf("no handler found for query type %s or %s", queryType.String(), nonPtrType.String())
			bus.log.Error(fmt.Sprintf("Query execution failed: %s", err.Error()))
			return nil, err
		}
	}

	queryHandler, ok := handler.(QueryHandler)
	if !ok {
		err := fmt.Errorf("invalid handler type for query %s", queryType.String())
		bus.log.Error(fmt.Sprintf("Query execution failed: %s", err.Error()))
		return nil, err
	}

	result, err := queryHandler.Handle(query)
	if err != nil {
		bus.log.Error(fmt.Sprintf("Error handling Query: %s", err.Error()))
	}

	return result, err
}

func (bus *InternalQueryBus) ExecuteWithContext(ctx context.Context, query Query) (interface{}, error) {
	resultChan := make(chan interface{}, 1)
	errChan := make(chan error, 1)

	go func() {
		result, err := bus.Execute(query)
		if err != nil {
			errChan <- err
		} else {
			resultChan <- result
		}
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case err := <-errChan:
		return nil, err
	case result := <-resultChan:
		return result, nil
	}
}
