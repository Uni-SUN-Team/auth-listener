package models

type Configs struct {
	App struct {
		ENV         string `mapstructure:"env"`
		Port        string `mapstructure:"port"`
		ContextPath string `mapstructure:"context_path"`
	} `mapstructure:"app"`
	Endpoint struct {
		ClassRoomListener struct {
			Host   string `mapstructure:"host"`
			Method string `mapstructure:"method"`
			Path   struct {
				List string `mapstructure:"list"`
				ID   string `mapstructure:"id"`
				Slug string `mapstructure:"slug"`
			} `mapstructure:"path"`
		} `mapstructure:"classroom_listener"`
	} `mapstructure:"endpoint"`
	Gin struct {
		Mode     string `mapstructure:"mode"`
		RootPath string `mapstructure:"root_path"`
		Version  string `mapstructure:"version"`
		Configs  struct {
			TrustedProxies string `mapstructure:"trusted_proxies"`
		} `mapstructure:"configs"`
	} `mapstructure:"gin"`
}
