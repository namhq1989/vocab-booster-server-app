package application

type (
	Commands interface {
	}
	Queries interface {
	}
	App interface {
		Commands
		Queries
	}

	appCommandHandlers struct {
	}
	appQueryHandler struct {
	}
	Application struct {
		appCommandHandlers
		appQueryHandler
	}
)

var _ App = (*Application)(nil)

func New() *Application {
	return &Application{
		appCommandHandlers: appCommandHandlers{},
		appQueryHandler:    appQueryHandler{},
	}
}
