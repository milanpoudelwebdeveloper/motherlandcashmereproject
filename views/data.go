package views

const (
	//AlertLvlError is
	AlertLvlError = "danger"
	//AlertLvlWarning is
	AlertLvlWarning = "warning"
	//AlertLvlInfo is
	AlertLvlInfo = "info"
	//AlertLevelSuccess is
	AlertLevelSuccess = "success"
)

//Alert is used to render Bootstrap Alert messages in templates
type Alert struct {
	Level   string
	Message string
}

//Data is the top level structure that views expect data to come in
type Data struct {
	Alert *Alert
	Yield interface{}
}
