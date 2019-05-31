package model

type Stats struct {
	Attack  float64 `json:"atk"`
	Defense float64 `json:"def"`
	HP      float64 `json:"hp"`
}

type Spread struct {
	Rank       int     `json:"rank"`
	IVs        string  `json:"ivs"`
	Level      float64 `json:"level"`
	CP         int     `json:"cp"`
	Product    float64 `json:"statProduct"`
	Percentage float64 `json:"percent"`
	Stats      Stats   `json:"stats"`
}
