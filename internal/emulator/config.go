package emulator

type Config struct {
	Debug         bool
	PathToRomFile string
}

func NewConfig(debug bool, pathToRomFile string) *Config {
	return &Config{
		Debug:         debug,
		PathToRomFile: pathToRomFile,
	}
}
