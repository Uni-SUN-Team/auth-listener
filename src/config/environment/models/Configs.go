package models

type Configs struct {
	App struct {
		ENV         string `mapstructure:"env"`
		Port        string `mapstructure:"port"`
		ContextPath string `mapstructure:"context_path"`
	} `mapstructure:"app"`
	Gin struct {
		Mode     string `mapstructure:"mode"`
		RootPath string `mapstructure:"root_path"`
		Version  string `mapstructure:"version"`
		Configs  struct {
			TrustedProxies string `mapstructure:"trusted_proxies"`
		} `mapstructure:"configs"`
	} `mapstructure:"gin"`
	Swag struct {
		Title       string `mapstructure:"title"`
		Description string `mapstructure:"description"`
		Version     string `mapstructure:"version"`
		Host        string `mapstructure:"host"`
		Schemes     string `mapstructure:"schemes"`
	} `mapstructure:"swag"`
	Secure struct {
		Jwt struct {
			Secret string `mapstructure:"secret"`
		} `mapstructure:"jwt"`
	} `mapstructure:"secure"`
}
