package sinkplugin

type Plugin interface {
	Execute() error
}
