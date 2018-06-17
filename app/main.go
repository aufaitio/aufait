package app

// CLI manages the configuration of Au Fait
type CLI struct {
	repositoryList []string
}

// NewCLI creates CLI for configuring Au Fait
func NewCLI(repositoryList []string) CLI {
	return CLI{repositoryList}
}

// ConfigureRepositories sets up provided repositories for Au Fait
func (cli CLI) ConfigureRepositories() error {
	return nil
}
