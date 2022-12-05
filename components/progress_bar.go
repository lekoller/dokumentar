package components

import (
	"time"

	"fyne.io/fyne/v2/widget"
)

type Progress struct {
	Bar *widget.ProgressBar
}

func NewProgressBar() *Progress {
	bar := widget.NewProgressBar()
	bar.Hidden = true

	return &Progress{Bar: bar}
}

func (p *Progress) Start() {
	p.Bar.Hidden = false
	go func() {
		for i := 0.0; i <= 1.0; i += 0.1 {
			time.Sleep(time.Millisecond * 100)
			p.Bar.SetValue(i)
		}
	}()
}

func (p *Progress) Finish() {
	p.Bar.SetValue(1)
	p.Bar.Hidden = true
}
