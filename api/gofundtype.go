package api


type GoFund struct {
	Code string `json:"code"`
	Name string `json:"name"`
	UpdateTime string `json:"update_time"`
	Value string `json:"value"`
	NowValue string `json:"now_value"`
	Valuation string `json:"valuation"`
	Trend string `json:"trend"`
}

const (
	UP="↑"
	DOWN="↓"
)

type Fund interface {
	Api(code string) (GoFund,error)
}
