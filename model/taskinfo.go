package model

type TaskInfo struct{
	Name string
	Url  string
	Rules string
}

type RuleModle struct {
	Name      string `xml:"name,attr"`
	ParseFunc string `xml:"ParseFunc>Script"`
	AidFunc   string `xml:"AidFunc>Script"`
}
