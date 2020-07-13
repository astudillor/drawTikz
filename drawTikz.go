package drawTikz

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Transformer interface {
	Scale(alpha float64)
	ScaleX(alpha float64)
	ScaleY(alpha float64)
	Translate(d float64)
	TranslateX(dx float64)
	TranslateY(dy float64)
}

type PointXY struct {
	x, y float64
}

func (p *PointXY) ParseStr(data string) error {
	if p == nil {
		p = new(PointXY)
	}
	re := regexp.MustCompile(`\(.*?\)`)
	result := string(re.Find([]byte(data)))
	if result == "" {
		return fmt.Errorf("Error parsing PointXY %q\n", string(result))
	}
	v := strings.Split(string(result), ",")
	if len(v) > 0 {
		xstr := string(v[0][1:])
		x, err := strconv.ParseFloat(xstr, 64)
		if err != nil {
			return fmt.Errorf("Error parsing x coordinate: %s", xstr)
		}
		p.x = x
	}
	if len(v) > 1 {
		ystr := string(v[1][:len(v[1])-1])
		y, err := strconv.ParseFloat(ystr, 64)
		if err != nil {
			return fmt.Errorf("Error parsing y coordinate: %s", ystr)
		}
		p.y = y
	}
	return nil
}

func (p *PointXY) Scale(alpha float64) {
	p.x *= alpha
	p.y *= alpha
}

func (p *PointXY) ScaleX(alpha float64) {
	p.x *= alpha
}

func (p *PointXY) ScaleY(alpha float64) {
	p.y *= alpha
}

func (p *PointXY) Translate(d float64) {
	p.x += d
	p.y += d
}

func (p *PointXY) TranslateX(dx float64) {
	p.x += dx
}

func (p *PointXY) TranslateY(dy float64) {
	p.y += dy
}

func (p *PointXY) String() string {
	if p == nil {
		return ""
	}
	return fmt.Sprintf("(%6.4f,%6.4f)", p.x, p.y)
}

type Square struct {
	points [4]*PointXY
}

func (s *Square) Scale(alpha float64) {
	for _, p := range s.points {
		if p != nil {
			p.Scale(alpha)
		}
	}
}

func (s *Square) ScaleX(alpha float64) {
	for _, p := range s.points {
		if p != nil {
			p.ScaleX(alpha)
		}
	}
}

func (s *Square) ScaleY(alpha float64) {
	for _, p := range s.points {
		if p != nil {
			p.ScaleY(alpha)
		}
	}
}

func (s *Square) Translate(d float64) {
	for _, p := range s.points {
		if p != nil {
			p.Translate(d)
		}
	}
}

func (s *Square) TranslateX(dx float64) {
	for _, p := range s.points {
		if p != nil {
			p.Translate(dx)
		}
	}
}

func (s *Square) TranslateY(dy float64) {
	for _, p := range s.points {
		if p != nil {
			p.Translate(dy)
		}
	}
}

func (s *Square) String() string {
	return fmt.Sprintf("\\draw %s -- %s -- %s -- %s;", s.points[0],
		s.points[1], s.points[2], s.points[3])
}

type Scene []Transformer
