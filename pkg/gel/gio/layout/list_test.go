// SPDX-License-Identifier: Unlicense OR MIT

package layout

import (
	"image"
	"testing"

	"github.com/p9c/p9/pkg/gel/gio/f32"
	"github.com/p9c/p9/pkg/gel/gio/io/event"
	"github.com/p9c/p9/pkg/gel/gio/io/pointer"
	"github.com/p9c/p9/pkg/gel/gio/io/router"
	"github.com/p9c/p9/pkg/gel/gio/op"
)

func TestListPosition(t *testing.T) {
	_s := func(e ...event.Event) []event.Event { return e }
	r := new(router.Router)
	gtx := Context{
		Ops: new(op.Ops),
		Constraints: Constraints{
			Max: image.Pt(20, 10),
		},
		Queue: r,
	}
	el := func(gtx Context, idx int) Dimensions {
		return Dimensions{Size: image.Pt(10, 10)}
	}
	for _, tc := range []struct {
		label  string
		num    int
		scroll []event.Event
		first  int
		count  int
		offset int
		last   int
	}{
		{label: "no item", last: 20},
		{label: "1 visible 0 hidden", num: 1, count: 1, last: 10},
		{label: "2 visible 0 hidden", num: 2, count: 2},
		{label: "2 visible 1 hidden", num: 3, count: 2},
		{label: "3 visible 0 hidden small scroll", num: 3, count: 3, offset: 5, last: -5,
			scroll: _s(
				pointer.Event{
					Source:   pointer.Mouse,
					Buttons:  pointer.ButtonPrimary,
					Type:     pointer.Press,
					Position: f32.Pt(0, 0),
				},
				pointer.Event{
					Source: pointer.Mouse,
					Type:   pointer.Scroll,
					Scroll: f32.Pt(5, 0),
				},
				pointer.Event{
					Source:   pointer.Mouse,
					Buttons:  pointer.ButtonPrimary,
					Type:     pointer.Release,
					Position: f32.Pt(5, 0),
				},
			)},
		{label: "3 visible 0 hidden small scroll 2", num: 3, count: 3, offset: 3, last: -7,
			scroll: _s(
				pointer.Event{
					Source:   pointer.Mouse,
					Buttons:  pointer.ButtonPrimary,
					Type:     pointer.Press,
					Position: f32.Pt(0, 0),
				},
				pointer.Event{
					Source: pointer.Mouse,
					Type:   pointer.Scroll,
					Scroll: f32.Pt(3, 0),
				},
				pointer.Event{
					Source:   pointer.Mouse,
					Buttons:  pointer.ButtonPrimary,
					Type:     pointer.Release,
					Position: f32.Pt(5, 0),
				},
			)},
		{label: "2 visible 1 hidden large scroll", num: 3, count: 2, first: 1,
			scroll: _s(
				pointer.Event{
					Source:   pointer.Mouse,
					Buttons:  pointer.ButtonPrimary,
					Type:     pointer.Press,
					Position: f32.Pt(0, 0),
				},
				pointer.Event{
					Source: pointer.Mouse,
					Type:   pointer.Scroll,
					Scroll: f32.Pt(10, 0),
				},
				pointer.Event{
					Source:   pointer.Mouse,
					Buttons:  pointer.ButtonPrimary,
					Type:     pointer.Release,
					Position: f32.Pt(15, 0),
				},
			)},
	} {
		t.Run(tc.label, func(t *testing.T) {
			gtx.Ops.Reset()

			var list List
			// Initialize the list.
			list.Layout(gtx, tc.num, el)
			// Generate the scroll events.
			r.Frame(gtx.Ops)
			r.Queue(tc.scroll...)
			// Let the list process the events.
			list.Layout(gtx, tc.num, el)

			pos := list.Position
			if got, want := pos.First, tc.first; got != want {
				t.Errorf("List: invalid first position: got %v; want %v", got, want)
			}
			if got, want := pos.Count, tc.count; got != want {
				t.Errorf("List: invalid number of visible children: got %v; want %v", got, want)
			}
			if got, want := pos.Offset, tc.offset; got != want {
				t.Errorf("List: invalid first visible offset: got %v; want %v", got, want)
			}
			if got, want := pos.OffsetLast, tc.last; got != want {
				t.Errorf("List: invalid last visible offset: got %v; want %v", got, want)
			}
		})
	}
}
