package model

type Fact struct {
	PeriodStart         string
	PeriodEnd           string
	PeriodKey           string
	IndicatorToMoId     int
	IndicatorToMoFactId int
	Value               int
	FactTime            string
	IsPlan              int
	AuthUserId          int
	Comment             string
}
