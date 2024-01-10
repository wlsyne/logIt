package constants

const ConfigFileName = "LogItConfig.json"
const ChangelogFileName = "CHANGELOG.md"

type Mode int

const (
	WriteMode Mode = iota
	PublishMode
)

const (
	RegexpStringUrl = `^https?://[\w-]+(\.[\w-]+)+(:\d{1,5})?(/[\w-./?%&=]*)?$`
)
