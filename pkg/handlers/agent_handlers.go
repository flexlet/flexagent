package handlers

import (
	"flexagent/pkg/services"
	"flexagent/restapi/operations/agent"

	"github.com/go-openapi/runtime/middleware"
)

type AgentReadyzHandlerImpl struct {
}

func (h *AgentReadyzHandlerImpl) Handle(params agent.ReadyzParams) middleware.Responder {
	return agent.NewReadyzOK().WithPayload(services.AgentService.Readyz())
}

type AgentHealthzHandlerImpl struct {
}

func (h *AgentHealthzHandlerImpl) Handle(params agent.HealthzParams) middleware.Responder {
	return agent.NewHealthzOK().WithPayload(services.AgentService.Healthz())
}
