package handlers

import (
	"flexagent/pkg/services"
	"flexagent/restapi/operations/cronjob"

	"github.com/go-openapi/runtime/middleware"
)

// submit cronjobs
type CronJobSubmitHandlerImpl struct {
}

func (h *CronJobSubmitHandlerImpl) Handle(params cronjob.SubmitCronJobsParams) middleware.Responder {
	cronJobList, err := services.CronJobService.BatchSubmit(params.Spec)
	if err != nil {
		errMsg := err.Error()
		body := &cronjob.SubmitCronJobsBadRequestBody{Message: &errMsg}
		return cronjob.NewSubmitCronJobsBadRequest().WithPayload(body)
	}
	return cronjob.NewSubmitCronJobsOK().WithPayload(cronJobList)
}

// list cronjobs
type CronJobListHandlerImpl struct {
}

func (h *CronJobListHandlerImpl) Handle(params cronjob.ListCronJobsParams) middleware.Responder {
	cronJobList, err := services.CronJobService.List(params.Name)
	if err != nil {
		errMsg := err.Error()
		body := &cronjob.ListCronJobsBadRequestBody{Message: &errMsg}
		return cronjob.NewListCronJobsBadRequest().WithPayload(body)
	}
	return cronjob.NewListCronJobsOK().WithPayload(cronJobList)
}

// query cronjob
type CronJobQueryHandlerImpl struct {
}

func (h *CronJobQueryHandlerImpl) Handle(params cronjob.QueryCronJobParams) middleware.Responder {
	cronJobResp, err := services.CronJobService.Query(params.ID)
	if err != nil {
		errMsg := err.Error()
		body := &cronjob.QueryCronJobBadRequestBody{Message: &errMsg}
		return cronjob.NewQueryCronJobBadRequest().WithPayload(body)
	}
	return cronjob.NewQueryCronJobOK().WithPayload(cronJobResp)
}

// update cronjob
type CronJobUpdateHandlerImpl struct {
}

func (h *CronJobUpdateHandlerImpl) Handle(params cronjob.UpdateCronJobParams) middleware.Responder {
	cronJobResp, err := services.CronJobService.Update(params.ID, params.Spec)
	if err != nil {
		errMsg := err.Error()
		body := &cronjob.UpdateCronJobBadRequestBody{Message: &errMsg}
		return cronjob.NewUpdateCronJobBadRequest().WithPayload(body)
	}
	return cronjob.NewUpdateCronJobOK().WithPayload(cronJobResp)
}

// delete cronjob
type CronJobDeleteHandlerImpl struct {
}

func (h *CronJobDeleteHandlerImpl) Handle(params cronjob.DeleteCronJobParams) middleware.Responder {
	err := services.CronJobService.Delete(params.ID)
	if err != nil {
		errMsg := err.Error()
		body := &cronjob.DeleteCronJobBadRequestBody{Message: &errMsg}
		return cronjob.NewDeleteCronJobBadRequest().WithPayload(body)
	}
	return cronjob.NewDeleteCronJobOK().WithPayload("ok")
}

// start cronjob
type CronJobStartHandlerImpl struct {
}

func (h *CronJobStartHandlerImpl) Handle(params cronjob.StartCronJobParams) middleware.Responder {
	cronJobResp, err := services.CronJobService.Start(params.ID)
	if err != nil {
		errMsg := err.Error()
		body := &cronjob.StartCronJobBadRequestBody{Message: &errMsg}
		return cronjob.NewStartCronJobBadRequest().WithPayload(body)
	}
	return cronjob.NewStartCronJobOK().WithPayload(cronJobResp)
}

// stop cronjob
type CronJobStopHandlerImpl struct {
}

func (h *CronJobStopHandlerImpl) Handle(params cronjob.StopCronJobParams) middleware.Responder {
	cronJobResp, err := services.CronJobService.Stop(params.ID)
	if err != nil {
		errMsg := err.Error()
		body := &cronjob.StopCronJobBadRequestBody{Message: &errMsg}
		return cronjob.NewStopCronJobBadRequest().WithPayload(body)
	}
	return cronjob.NewStopCronJobOK().WithPayload(cronJobResp)
}
