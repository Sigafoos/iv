package main

import (
	"fmt"
	"math"
	"sort"
)

type Pokemon struct {
	Dex   int    `json:"dex"`
	Name  string `json:"speciesName"`
	ID    string `json:"speciesId"`
	Stats Stats  `json:"baseStats"`
}

type Stats struct {
	Attack  int `json:"atk"`
	Defense int `json:"def"`
	HP      int `json:"hp"`
}

type Spread struct {
	Rank       int     `json:"rank"`
	IVs        string  `json:"ivs"`
	Level      float64 `json:"level"`
	CP         int     `json:"cp"`
	Product    float64 `json:"statProduct"`
	Percentage float64 `json:"percent"`
}

var cpm = []float64{
	0.094,
	0.1351374318,
	0.16639787,
	0.192650919,
	0.21573247,
	0.2365726613,
	0.25572005,
	0.2735303812,
	0.29024988,
	0.3060573775,
	0.3210876,
	0.3354450362,
	0.34921268,
	0.3624577511,
	0.3752356,
	0.387592416,
	0.39956728,
	0.4111935514,
	0.4225,
	0.4329264091,
	0.44310755,
	0.4530599591,
	0.4627984,
	0.472336093,
	0.48168495,
	0.4908558003,
	0.49985844,
	0.508701765,
	0.51739395,
	0.5259425113,
	0.5343543,
	0.5426357375,
	0.5507927,
	0.5588305862,
	0.5667545,
	0.5745691333,
	0.5822789,
	0.5898879072,
	0.5974,
	0.6048236651,
	0.6121573,
	0.6194041216,
	0.6265671,
	0.6336491432,
	0.64065295,
	0.6475809666,
	0.65443563,
	0.6612192524,
	0.667934,
	0.6745818959,
	0.6811649,
	0.6876849038,
	0.69414365,
	0.70054287,
	0.7068842,
	0.7131691091,
	0.7193991,
	0.7255756136,
	0.7317,
	0.7347410093,
	0.7377695,
	0.7407855938,
	0.74378943,
	0.7467812109,
	0.74976104,
	0.7527290867,
	0.7556855,
	0.7586303683,
	0.76156384,
	0.7644860647,
	0.76739717,
	0.7702972656,
	0.7731865,
	0.7760649616,
	0.77893275,
	0.7817900548,
	0.784637,
	0.7874736075,
	0.7903,
}

func (p *Pokemon) Spreads() map[string]Spread {
	spreads := make(map[string]Spread)
	var working []Spread

	for atk := 0; atk < 16; atk++ {
		for def := 0; def < 16; def++ {
			for hp := 0; hp < 16; hp++ {
				for level := len(cpm) - 1; level >= 0; level-- {
					cp, product := p.CP(level, atk, def, hp)
					if cp <= 1500 {
						working = append(working, Spread{
							IVs:     fmt.Sprintf("%v/%v/%v", atk, def, hp),
							Level:   float64(level/2) + 1,
							CP:      cp,
							Product: product,
						})
						break
					}
				}
			}
		}
	}

	sort.Slice(working, func(i, j int) bool { return working[i].Product > working[j].Product })

	ideal := working[0].Product
	for k, v := range working {
		v.Rank = k
		v.Percentage = v.Product / ideal * 100
		spreads[v.IVs] = v
	}

	return spreads
}

func (p *Pokemon) CP(level, atk, def, hp int) (int, float64) {
	attack := float64(p.Stats.Attack+atk) * cpm[level]
	defense := float64(p.Stats.Defense+def) * cpm[level]
	stamina := float64(p.Stats.HP+hp) * cpm[level]

	cp := math.Floor((math.Pow(stamina, 0.5) * attack * math.Pow(defense, 0.5)) / 10)
	product := attack * defense * stamina
	if cp < 10 {
		return 10, product
	}
	return int(cp), product
}
