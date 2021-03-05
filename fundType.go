package main


type fund struct {
	Fundcode string `json:"fundcode"`
	Name string `json:"name"`
	Jzrq string `json:"jzrq"`
	Dwjz string `json:"dwjz"`
	Gsz string `json:"gsz"`
	Gszzl string `json:"gszzl"`
	Gztime string `json:"gztime"`
}
// [{"name": "mpd", "instance": "now playing", "full_text": " '${status}' '$1'", "color": "'${color}'"}]
type i3status struct {
	Name string `json:"name"`
	Instance string `json:"instance"`
	FullText string `json:"full_text"`
	Color string `json:"color"`
}
