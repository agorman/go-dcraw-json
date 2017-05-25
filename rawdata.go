package dcraw

type RawData struct {
	Filename  string  `json:"filename"`
	Timestamp int64   `json:"timestamp"`
	ISO       int     `json:"iso"`
	Shutter   float64 `json:"shutter"`
	Aperture  float64 `json:"aperture"`
	Focal     float64 `json:"focal"`
	Camera    Camera  `json:"camera"`
	ThumbSize Size    `json:"thumbSize"`
	FullSize  Size    `json:"fullSize"`
	GPS       GPS     `json:"gps"`
}

type Camera struct {
	Make  string `json:"make"`
	Model string `json:"model"`
}

type Size struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type GPS struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
