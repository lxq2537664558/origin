package originnode

import (
	_ "net/http/pprof"
)

/*
type GlobalModule struct {
	service.BaseModule
}

// 全局模块定义
var g_module GlobalModule

func AddModule(module service.IModule) uint32 {
	return g_module.AddModule(module)
}

func GetModuleById(moduleId uint32) service.IModule {
	return g_module.GetModuleById(moduleId)
}

func GetLog(logmodule uint32) sysmodule.ILogger {
	module := g_module.GetModuleById(logmodule)
	if nil == module {
		return nil
	}

	return module.(sysmodule.ILogger)
}

func InitGlobalModule(exit chan bool, pwaitGroup *sync.WaitGroup) {
	g_module.InitModule(exit, pwaitGroup)
}

func RunGlobalModule() {
	go g_module.RunModule(&g_module)
}
*/
