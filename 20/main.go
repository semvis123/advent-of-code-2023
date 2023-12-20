package main

import (
	"log"
	"semvis123/aoc"
	"strings"

	"github.com/oleiade/lane/v2"
	"golang.org/x/exp/maps"
)

func main() {
	lines := aoc.GetInput()
	log.Printf("part 1: %d", part_1(lines[:len(lines)-1]))
	log.Printf("part 2: %d", part_2(lines[:len(lines)-1]))
}

type Pulse struct {
	val        bool
	from, dest string
}

type Controller struct {
	modules      map[string]Module
	toSend       *lane.Queue[Pulse]
	highCount    int
	lowCount     int
	buttonPress  int
	shouldStop   bool
	rxBtnPresses int
	rxBits       map[string]int
}

func (c *Controller) Send(from string, val bool, dest []string) {
	for _, d := range dest {
		c.toSend.Enqueue(Pulse{val, from, d})
	}
}

func (c *Controller) Tick() bool {
	x, _ := c.toSend.Dequeue()
	if _, exists := c.modules[x.dest]; exists {
		c.modules[x.dest].activate(x.from, x.val)
	}
	if x.val {
		c.highCount++
	} else {
		c.lowCount++
	}
	// log.Printf("%v -%v-> %v", x.from, aoc.Iff(x.val, "high", "low"), x.dest)

	if x.dest == "zh" && x.val {
		if _, ok := c.rxBits[x.from]; ok {
			if len(c.modules["zh"].(*Conjuctor).vals) == len(c.rxBits) {
				c.rxBtnPresses = aoc.LCM(1, 1, maps.Values(c.rxBits)...)
				c.shouldStop = true
			}
		} else {
			c.rxBits[x.from] = c.buttonPress
		}
	}

	return c.toSend.Size() > 0 && !c.shouldStop
}

func NewController(input []string) Controller {
	controller := Controller{}
	controller.modules = make(map[string]Module)
	controller.toSend = lane.NewQueue[Pulse]()
	controller.rxBits = make(map[string]int)

	for _, l := range input {
		nameStr, outputs, _ := strings.Cut(l, " -> ")
		if name, ok := strings.CutPrefix(nameStr, "%"); ok {
			mod := FlipFlop{name: name, outputs: strings.Split(outputs, ", ")}
			mod.c = controller
			controller.modules[name] = &mod
		}
		if name, ok := strings.CutPrefix(nameStr, "&"); ok {
			mod := Conjuctor{name: name, outputs: strings.Split(outputs, ", ")}
			mod.c = controller
			mod.vals = make(map[string]bool)
			controller.modules[name] = &mod
		}
		if nameStr == "broadcaster" {
			mod := Broadcaster{name: nameStr, outputs: strings.Split(outputs, ", ")}
			mod.c = controller
			controller.modules[nameStr] = &mod
		}
	}

	for name, mod := range controller.modules {
		for _, o := range mod.getOutputs() {
			oMod := controller.modules[o]
			if conj, ok := oMod.(*Conjuctor); ok {
				conj.vals[name] = false
			}
		}
	}

	return controller
}

type Module interface {
	activate(string, bool)
	getOutputs() []string
}

type FlipFlop struct {
	name    string
	c       Controller
	outputs []string
	val     bool
}

func (f *FlipFlop) activate(from string, pulse bool) {
	if pulse {
		return
	}
	f.val = !f.val
	f.c.Send(f.name, f.val, f.outputs)
}

func (f *FlipFlop) getOutputs() []string {
	return f.outputs
}

type Broadcaster struct {
	name    string
	c       Controller
	outputs []string
}

func (b *Broadcaster) activate(from string, pulse bool) {
	b.c.Send(b.name, pulse, b.outputs)
}

func (f *Broadcaster) getOutputs() []string {
	return f.outputs
}

type Conjuctor struct {
	name    string
	c       Controller
	outputs []string
	vals    map[string]bool
}

func (c *Conjuctor) activate(from string, pulse bool) {
	c.vals[from] = pulse
	c.c.Send(c.name, !aoc.All(maps.Values(c.vals)), c.outputs)
}

func (f *Conjuctor) getOutputs() []string {
	return f.outputs
}

func part_1(input []string) (s int) {
	controller := NewController(input)

	for i := 0; i < 1000; i++ {
		controller.Send("button", false, []string{"broadcaster"})
		for controller.Tick() {
		}
	}

	return controller.lowCount * controller.highCount
}

func part_2(input []string) (s int) {
	controller := NewController(input)

	for !controller.shouldStop {
		controller.buttonPress++
		controller.Send("button", false, []string{"broadcaster"})
		for controller.toSend.Size() > 0 {
			controller.Tick()
		}
	}

	return controller.rxBtnPresses
}
