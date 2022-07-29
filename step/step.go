package step

import (
	"strings"
	"sync"

	"github.com/pav5000/algoviz-go/element"
)

type Step struct {
	lock sync.Mutex

	width    uint32
	height   uint32
	elements []element.Element
}

func New(width uint32, height uint32) *Step {
	return &Step{
		width:  width,
		height: height,
	}
}

// Serialize возвращает полный SVG данного шага
func (s *Step) Serialize() string {
	s.lock.Lock()
	defer s.lock.Unlock()

	str := strings.Builder{}
	for _, elem := range s.elements {
		str.WriteString(elem.SVG())
	}
	return str.String()
}
