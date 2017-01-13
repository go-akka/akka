package akka

type Deployer struct {
	settings Settings
}

func NewDeployer(settings Settings) Deployer {
	return Deployer{
		settings: settings,
	}
}
