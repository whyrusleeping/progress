package progress

import (
	"fmt"
)

type Progress struct {
	Total    int
	Complete int
	Children map[string]*Progress
	Parent   *Progress

	Finished bool
}

func (p *Progress) Child(name string, numTasks int) *Progress {
	_, ok := p.Children[name]
	if ok {
		panic("child by that name already exists!")
	}

	np := NewProgress(numTasks)
	np.Parent = p
	p.Children[name] = np
	p.Total++

	return np
}

func (p *Progress) clearBar() {
}

// write to the screen, but don't mess up the progress bar
func (p *Progress) Write(b []byte) (int, error) {
	panic("NYI")
}

func (p *Progress) Finish() {
	if p.Total != p.Complete {
		panic("Finish called before all tasks completed")
	}

	if p.Finished {
		panic("finished called twice!")
	}

}

func (p *Progress) Ratio() (int, int) {
	return p.Complete, p.Total
}

func NewProgress(num int) *Progress {
	return &Progress{
		Total:    num,
		Children: make(map[string]*Progress),
	}
}
