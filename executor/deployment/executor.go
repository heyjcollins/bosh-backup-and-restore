package deployment

//go:generate counterfeiter -o fakes/fake_executor.go . Executor
type DeploymentExecutor interface {
	Run([]Executable) []DeploymentError
}

//go:generate counterfeiter -o fakes/fake_executable.go . Executable
type Executable interface {
	Execute() DeploymentError
}