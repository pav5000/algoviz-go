package collector

import (
	"encoding/json"
	"sync"

	"github.com/pav5000/algoviz-go/step"
)

type Collector struct {
	lock sync.Mutex

	width  uint32
	height uint32
	steps  []*step.Step
}

func New(width uint32, height uint32) *Collector {
	return &Collector{
		width:  width,
		height: height,
	}
}

func (c *Collector) NewStep() *step.Step {
	step := step.New(c.width, c.height)
	c.lock.Lock()
	c.steps = append(c.steps, step)
	c.lock.Unlock()
	return step
}

type Contract struct {
	Width  uint32   `json:"width"`
	Height uint32   `json:"height"`
	Steps  []string `json:"steps"`
}

func (c *Collector) Serialize() []byte {
	c.lock.Lock()
	defer c.lock.Unlock()

	contract := Contract{
		Steps: []string{},
	}

	for _, step := range c.steps {
		contract.Steps = append(contract.Steps, step.Serialize())
	}

	rawJSON, err := json.Marshal(&contract)
	if err != nil {
		panic("cannot marshal JSON:" + err.Error())
	}

	return rawJSON
}
