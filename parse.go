package csscolor

import (
	"image/color"
	"strconv"
	"strings"
)

func Parse(col string) (color.Color, error) {
	// col can be one of many things
	// not supported yet: rgb() rgba() hls() etc...
	col = strings.TrimSpace(strings.ToLower(col))

	if len(col) == 0 {
		return nil, ErrInvalidColor
	}

	if col[0] == '#' {
		// hex color, either #rgb or #rrggbb
		switch len(col) {
		case 4: // #rgb
			v, err := strconv.ParseUint(col[1:], 16, 16)
			if err != nil {
				return nil, err
			}
			r := uint8((v >> 8) & 0xf)
			g := uint8((v >> 4) & 0xf)
			b := uint8(v & 0xf)
			r, g, b = r|r<<4, g|g<<4, b|b<<4
			return color.RGBA{R: r, G: g, B: b, A: 0xff}, nil
		case 7: // #rrggbb
			v, err := strconv.ParseUint(col[1:], 16, 32)
			if err != nil {
				return nil, err
			}
			r := uint8((v >> 16) & 0xff)
			g := uint8((v >> 8) & 0xff)
			b := uint8(v & 0xff)
			return color.RGBA{R: r, G: g, B: b, A: 0xff}, nil
		case 9: // #rrggbbaa
			v, err := strconv.ParseUint(col[1:], 16, 32)
			if err != nil {
				return nil, err
			}
			r := uint8((v >> 24) & 0xff)
			g := uint8((v >> 16) & 0xff)
			b := uint8((v >> 8) & 0xff)
			a := uint8(v & 0xff)
			return color.RGBA{R: r, G: g, B: b, A: a}, nil
		}
		return nil, ErrInvalidColor
	}

	if v, ok := namedColors[col]; ok {
		return Parse(v)
	}

	return nil, ErrInvalidColor
}
