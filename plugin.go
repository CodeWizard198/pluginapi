package pluginapi

type Plugin interface {
	Name() string          // 插件名
	Execute() (any, error) // 具体的执行函数
}
