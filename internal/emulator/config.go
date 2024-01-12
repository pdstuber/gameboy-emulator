package emulator

type Config struct {
	Debug             bool
	PathToRomFile     string
	PathToBootRomFile string
}

func NewConfig(debug bool, pathToBootRomFile, pathToRomFile string) *Config {
	return &Config{
		Debug:             debug,
		PathToBootRomFile: pathToBootRomFile,
		PathToRomFile:     pathToRomFile,
	}
}
