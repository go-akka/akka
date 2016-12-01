package actor

type Settings struct {
	config Config

	ConfigVersion string
	Loggers       []string
}

func NewSettings(config Config, name string) (settings Settings, err error) {
	return
}
