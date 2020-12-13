package endpoints

import (
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"

	"github.com/MicahParks/terse-URL/configure"
	"github.com/MicahParks/terse-URL/models"
	"github.com/MicahParks/terse-URL/restapi/operations/api"
	"github.com/MicahParks/terse-URL/storage"
)

func HandleExport(logger *zap.SugaredLogger, terseStore storage.TerseStore) api.TerseExportHandlerFunc {
	return func(params api.TerseExportParams) middleware.Responder {

		// Debug info.
		logger.Debug("Performing data dump.")

		// TODO Non-debug level log?

		// Create a request context.
		//
		// Maybe make a longer context if timing out.
		ctx, cancel := configure.DefaultCtx()
		defer cancel()

		// Get the data dump.
		dump, err := terseStore.ExportAll(ctx)
		if err != nil {

			// Log at the appropriate level.
			logger.Errorw("Failed to perform data dump.",
				"error", err.Error(),
			)

			// Report the error to the client.
			code := int64(500)
			message := "Failed to perform data dump."
			return &api.TerseExportDefault{Payload: &models.Error{
				Code:    &code,
				Message: &message,
			}}
		}

		return &api.TerseExportOK{
			Payload: dump,
		}
	}
}