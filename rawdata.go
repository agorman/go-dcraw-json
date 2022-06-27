package dcraw

// RawData defines the output from the dcraw-json command
type RawData struct {
	Filename  string  `json:"filename"`
	Timestamp int64   `json:"timestamp"`
	ISO       int     `json:"iso"`
	Shutter   float64 `json:"shutter"`
	Aperture  float64 `json:"aperture"`
	Focal     float64 `json:"focal"`
	Camera    *Camera `json:"camera"`
	ThumbSize *Size   `json:"thumbSize"`
	FullSize  *Size   `json:"fullSize"`
	GPS       *GPS    `json:"gps,omitempty"`
}

// Camera defines an image's camera
type Camera struct {
	Make  string `json:"make"`
	Model string `json:"model"`
}

// Size defines an image's size
type Size struct {
	Width  uint `json:"width"`
	Height uint `json:"height"`
}

// Size defines an image's GPS coordinates
type GPS struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
