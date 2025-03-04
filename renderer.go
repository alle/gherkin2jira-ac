package main

import (
	"strings"

	messages "github.com/cucumber/messages/go/v22"
	"github.com/willf/pad/utf8"
)

type renderer struct {
	*strings.Builder
}

func NewRenderer() *renderer {
	return &renderer{&strings.Builder{}}
}

func (r *renderer) Render(d *messages.GherkinDocument) string {
	r.renderFeature(d.Feature)

	return r.Builder.String()
}

func (r *renderer) renderFeature(f *messages.Feature) {
	s := "**" + f.Keyword + ":"

	if f.Name != "" {
		s += " " + f.Name
	}

	r.writeLine(s + "**")

	r.writeDescription(f.Description)

	for _, c := range f.Children {
		r.writeLine("")

		if c.Background != nil {
			r.renderBackground(c.Background)
		}

		if c.Scenario != nil {
			r.renderScenario(c.Scenario)
		}
	}
}

func (r *renderer) renderBackground(b *messages.Background) {
	s := "**" + b.Keyword + ":"

	if b.Name != "" {
		s += " " + b.Name
	}

	r.writeLine(s + "**")

	r.writeDescription(b.Description)
	r.renderSteps(b.Steps)
}

func (r *renderer) renderScenario(s *messages.Scenario) {
	st := "**" + s.Keyword + ":"

	if s.Name != "" {
		st += " " + s.Name
	}

	r.writeLine(st + "**")

	r.writeDescription(s.Description)
	r.renderSteps(s.Steps)

	if len(s.Examples) != 0 {
		r.writeLine("")
		r.renderExamples(s.Examples)
	}
}

func (r *renderer) renderSteps(ss []*messages.Step) {
	if len(ss) != 0 {
		r.writeLine("")
	}

	for _, s := range ss {
		r.renderStep(s)
	}
}

func (r *renderer) renderDocString(d *messages.DocString) {
	r.writeLine("```" + d.MediaType)
	r.writeLine(d.Content)
	r.writeLine("```")
}

func (r *renderer) renderStep(s *messages.Step) {
	r.writeLine("**" + strings.TrimSpace(s.Keyword) + "** " + s.Text + "")

	if s.DocString != nil {
		r.writeLine("")
		r.renderDocString(s.DocString)
	}

	if s.DataTable != nil {
		r.writeLine("")
		r.renderDataTable(s.DataTable)
	}
}

func (r *renderer) renderExamples(es []*messages.Examples) {
	r.writeLine("**Examples**:")

	for _, e := range es {
		if e.Name != "" {
			r.writeLine("")
			r.writeLine("*" + e.Name + "*")
		}

		r.writeDescription(e.Description)

		r.writeLine("")
		r.renderExampleTable(e.TableHeader, e.TableBody)
	}
}

func (r renderer) renderExampleTable(h *messages.TableRow, rs []*messages.TableRow) {
	ws := r.getCellWidths(append([]*messages.TableRow{h}, rs...))

	r.renderCells(h.Cells, ws)

	s := "|"

	for _, w := range ws {
		s += strings.Repeat("-", w+2) + "|"
	}

	r.writeLine(s)

	for _, t := range rs {
		r.renderCells(t.Cells, ws)
	}
}

func (r renderer) renderDataTable(t *messages.DataTable) {
	ws := r.getCellWidths(t.Rows)

	for _, t := range t.Rows {
		r.renderCells(t.Cells, ws)
	}
}

func (r renderer) renderCells(cs []*messages.TableCell, ws []int) {
	s := "|"

	for i, c := range cs {
		s += " " + utf8.Right(c.Value, ws[i], " ") + " |"
	}

	r.writeLine(s)
}

func (renderer) getCellWidths(rs []*messages.TableRow) []int {
	ws := make([]int, len(rs[0].Cells))

	for _, r := range rs {
		for i, c := range r.Cells {
			if w := len(c.Value); w > ws[i] {
				ws[i] = w
			}
		}
	}

	return ws
}

func (r renderer) writeDescription(s string) {
	trimmed := strings.TrimSpace(s)

	if trimmed != "" {
		r.writeLine("")

		lines := strings.Split(trimmed, "\n")
		for _, line := range lines {
			r.writeLine(strings.TrimSpace(line))
		}
	}
}

func (r renderer) writeLine(s string) {
	_, err := r.WriteString(s + "\n")

	if err != nil {
		panic(err)
	}
}
