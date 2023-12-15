package constants

const ConfigFileName = "LogItConfig.json"

type Mode int

const (
	WriteMode Mode = iota
	PublishMode
)
