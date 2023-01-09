package handlers

import (
	"flexagent/pkg/services"
	"flexagent/restapi/operations/job"

	"github.com/go-openapi/runtime/middleware"
)

// submit jobs
type JobSubmitHandlerImpl struct {
}

func (h *JobSubmitHandlerImpl) Handle(params job.SubmitParams) middleware.Responder {
	jobList, err := services.JobService.BatchSubmit(params.Spec, params.Wait)
	if err != nil {
		errMsg := err.Error()
		body := &job.SubmitBadRequestBody{Message: &errMsg}
		return job.NewSubmitBadRequest().WithPayload(body)
	}
	return job.NewSubmitOK().WithPayload(jobList)
}

// list jobs
type JobListHandlerImpl struct {
}

func (h *JobListHandlerImpl) Handle(params job.ListParams) middleware.Responder {
	jobList, err := services.JobService.List(params.Plugin, params.Operation, params.StartTimeBegin, params.StartTimeEnd)
	if err != nil {
		errMsg := err.Error()
		body := &job.ListBadRequestBody{Message: &errMsg}
		return job.NewListBadRequest().WithPayload(body)
	}
	return job.NewListOK().WithPayload(jobList)
}

// query job
type JobQueryHandlerImpl struct {
}

func (h *JobQueryHandlerImpl) Handle(params job.QueryParams) middleware.Responder {
	jobResp, err := services.JobService.Query(params.Urn, params.OutputLineStart, params.OutputLineLimit)
	if err != nil {
		errMsg := err.Error()
		body := &job.QueryBadRequestBody{Message: &errMsg}
		return job.NewQueryBadRequest().WithPayload(body)
	}
	return job.NewQueryOK().WithPayload(jobResp)
}

// delete job
type JobDeleteHandlerImpl struct {
}

func (h *JobDeleteHandlerImpl) Handle(params job.DeleteParams) middleware.Responder {
	err := services.JobService.Delete(params.Urn)
	if err != nil {
		errMsg := err.Error()
		body := &job.DeleteBadRequestBody{Message: &errMsg}
		return job.NewDeleteBadRequest().WithPayload(body)
	}
	return job.NewDeleteOK().WithPayload("ok")
}

// input job
type JobInputHandlerImpl struct {
}

func (h *JobInputHandlerImpl) Handle(params job.InputParams) middleware.Responder {
	jobResp, err := services.JobService.Input(params.Urn, params.Input)
	if err != nil {
		errMsg := err.Error()
		body := &job.InputBadRequestBody{Message: &errMsg}
		return job.NewInputBadRequest().WithPayload(body)
	}
	return job.NewInputOK().WithPayload(jobResp)
}

// kill job
type JobKillHandlerImpl struct {
}

func (h *JobKillHandlerImpl) Handle(params job.KillParams) middleware.Responder {
	jobResp, err := services.JobService.Kill(params.Urn, params.Force)
	if err != nil {
		errMsg := err.Error()
		body := &job.KillBadRequestBody{Message: &errMsg}
		return job.NewKillBadRequest().WithPayload(body)
	}
	return job.NewKillOK().WithPayload(jobResp)
}
