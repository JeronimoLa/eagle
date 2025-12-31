package main

func main() {
	appCfg := &appConfig{}
	buildAppConfig(appCfg)
	Start(appCfg)
}
