package avatar

type Config struct {
	Width          int     `json:"width"`
	Height         int     `json:"height"`
	FontSize       float64 `json:"font_size"`
	FontPath       string  `json:"font_path"`
	ColorsPath     string  `json:"colors_path"`
	OutputDir      string  `json:"output_dir"`
	TruncateLength int     `json:"truncate_length"`
}