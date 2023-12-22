package write

type ChangelogType string

const (
	Feat        ChangelogType = "✨ Feat"
	Doc         ChangelogType = "📝 Doc"
	Fix         ChangelogType = "🐛 Fix"
	Style       ChangelogType = "🎨 Style"
	SpeedUp     ChangelogType = "⚡️ SpeedUp"
	Config      ChangelogType = "🔧 Config"
	Test        ChangelogType = "✅ Test"
	BreakChange ChangelogType = "💥 BreakChange"
	Finish      ChangelogType = "Finish"
	Cancel      ChangelogType = "Cancel"
)

var ChangelogTypeList = []ChangelogType{
	Feat,
	Doc,
	Fix,
	Style,
	SpeedUp,
	Config,
	Test,
	BreakChange,
	Finish,
	Cancel,
}
