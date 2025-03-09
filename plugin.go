package pluginapi

type Plugin interface {
	Execute() error
}
