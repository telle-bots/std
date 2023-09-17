package runner

// Runner runs all commands and stops them if any failed or stopped
type Runner interface {
	// Add command to runner
	Add(cmd Command)
	// Start all runner's commands
	Start() error
}

// Command represents async action
type Command interface {
	// Run command
	Run() error
	// Stop command
	Stop()
}

// Start runner
func Start(run Runner) error {
	return run.Start()
}

// AddCommand to run
func AddCommand[C Command]() func(Runner, C) {
	return func(run Runner, command C) {
		run.Add(command)
	}
}
