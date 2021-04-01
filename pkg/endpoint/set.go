package endpoint

import (
	"github.com/brumhard/geckgo/pkg/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-playground/validator/v10"
)

// Set collects all of the endpoints that compose an add service. It's meant to
// be used as a helper struct, to collect all of the endpoints into a single
// parameter.
type Set struct {
	AddList    endpoint.Endpoint
	GetLists   endpoint.Endpoint
	GetList    endpoint.Endpoint
	UpdateList endpoint.Endpoint
	DeleteList endpoint.Endpoint
	AddDay     endpoint.Endpoint
	GetDays    endpoint.Endpoint
	GetDay     endpoint.Endpoint
	UpdateDay  endpoint.Endpoint
	DeleteDay  endpoint.Endpoint
}

// New returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func New(svc service.Service) Set {
	valMW := validationMW(validator.New())

	return Set{
		AddList:    valMW(makeAddListEndpoint(svc)),
		GetLists:   valMW(makeGetListsEndpoint(svc)),
		GetList:    valMW(makeGetListEndpoint(svc)),
		UpdateList: valMW(makeUpdateListEndpoint(svc)),
		DeleteList: valMW(makeDeleteListEndpoint(svc)),
		AddDay:     valMW(makeAddDayEndpoint(svc)),
		GetDays:    valMW(makeGetDaysEndpoint(svc)),
		GetDay:     valMW(makeGetDayEndpoint(svc)),
		UpdateDay:  valMW(makeUpdateDayEndpoint(svc)),
		DeleteDay:  valMW(makeDeleteDayEndpoint(svc)),
	}
}
