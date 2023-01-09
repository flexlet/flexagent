package restapi

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"gitee.com/yaohuiwang/utils"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"flexagent/pkg/config"
	"flexagent/pkg/handlers"
	"flexagent/pkg/services"
	"flexagent/restapi/operations"
)

// agent command line options
var options = struct {
	Config  string `short:"c" long:"config" env:"NODE_AGENT_CONFIG" description:"Configuration file" default:"/etc/flexagent.yaml" group:"agent"`
	Version bool   `short:"v" long:"version" description:"Show version" group:"agent"`
}{}

func configureFlags(api *operations.AgentAPI) {
	optionsGroup := swag.CommandLineOptionsGroup{
		ShortDescription: "Agent options",
		LongDescription:  "Agent options",
		Options:          &options,
	}
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		optionsGroup,
	}
}

func configureAPI(api *operations.AgentAPI) http.Handler {
	// show agent version and exit
	if options.Version {
		fmt.Printf("%s v%s\n", config.PROJECT, config.VERSION)
		os.Exit(0)
	}

	// load config file
	config.LoadAgentConfig(options.Config)

	// set log level
	utils.LogLevel = int(config.AgentConfig.LogLevel)

	// agent service
	services.AgentService = services.NewAgentService()

	// agent handlers
	api.AgentReadyzHandler = &handlers.AgentReadyzHandlerImpl{}
	api.AgentHealthzHandler = &handlers.AgentHealthzHandlerImpl{}

	// crypto service
	services.CryptoService = services.NewCryptoService()

	// crypto handlers
	api.CryptoEncryptHandler = &handlers.CryptoEncryptHandlerImpl{}
	api.CryptoDecryptHandler = &handlers.CryptoDecryptHandlerImpl{}
	api.CryptoSecretEncryptHandler = &handlers.CryptoSecretEncryptHandlerImpl{}
	api.CryptoSecretDecryptHandler = &handlers.CryptoSecretDecryptHandlerImpl{}
	api.CryptoListVaultsHandler = &handlers.CryptoListVaultsHandlerImpl{}
	api.CryptoCreateVaultHandler = &handlers.CryptoCreateVaultHandlerImpl{}
	api.CryptoUpdateVaultHandler = &handlers.CryptoUpdateVaultHandlerImpl{}
	api.CryptoQueryVaultHandler = &handlers.CryptoQueryVaultHandlerImpl{}
	api.CryptoDeleteVaultHandler = &handlers.CryptoDeleteVaultHandlerImpl{}

	// job services
	services.JobService = services.NewJobService()

	// job handlers
	api.JobSubmitHandler = &handlers.JobSubmitHandlerImpl{}
	api.JobListHandler = &handlers.JobListHandlerImpl{}
	api.JobQueryHandler = &handlers.JobQueryHandlerImpl{}
	api.JobInputHandler = &handlers.JobInputHandlerImpl{}
	api.JobKillHandler = &handlers.JobKillHandlerImpl{}
	api.JobDeleteHandler = &handlers.JobDeleteHandlerImpl{}

	// cronjob services
	services.CronJobService = services.NewCronJobService()

	// cronjob handlers
	api.CronjobSubmitCronJobsHandler = &handlers.CronJobSubmitHandlerImpl{}
	api.CronjobListCronJobsHandler = &handlers.CronJobListHandlerImpl{}
	api.CronjobQueryCronJobHandler = &handlers.CronJobQueryHandlerImpl{}
	api.CronjobUpdateCronJobHandler = &handlers.CronJobUpdateHandlerImpl{}
	api.CronjobDeleteCronJobHandler = &handlers.CronJobDeleteHandlerImpl{}
	api.CronjobStartCronJobHandler = &handlers.CronJobStartHandlerImpl{}
	api.CronjobStopCronJobHandler = &handlers.CronJobStopHandlerImpl{}

	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	// api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.PreServerShutdown = func() {
		services.JobService.GracefulShutdown()
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
