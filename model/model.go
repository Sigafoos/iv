package model

type Stats struct {
	Attack  float64 `json:"atk"`
	Defense float64 `json:"def"`
	HP      float64 `json:"hp"`
}

type Spread struct {
	Ranks      Ranks   `json:"ranks"`
	IVs        string  `json:"ivs"`
	Level      float64 `json:"level"`
	CP         int     `json:"cp"`
	Product    float64 `json:"statProduct"`
	Percentage float64 `json:"percent"`
	Stats      Stats   `json:"stats"`
}

type Ranks struct {
	All     *rank `json:"all,omitempty"`            // min 0
	Good    *rank `json:"goodFriend,omitempty"`     // min 1
	Great   *rank `json:"greatFriend,omitempty"`    // min 2
	Ultra   *rank `json:"ultraFriend,omitempty"`    // min 3
	Weather *rank `json:"weatherBoosted,omitempty"` // min 4
	Best    *rank `json:"bestFriend,omitempty"`     // min 5
	Hatched *rank `json:"hatched,omitempty"`        // min 10
	Lucky   *rank `json:"luckyFriend,omitempty"`    // min 12
}

type rank int

func Rank(i int) *rank {
	r := rank(i)
	return &r
}
