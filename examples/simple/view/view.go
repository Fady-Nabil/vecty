package view

import (
	"github.com/neelance/dom"
	"github.com/neelance/dom/bind"
	"github.com/neelance/dom/elem"
	"github.com/neelance/dom/event"
	"github.com/neelance/dom/examples/simple/model"
	"github.com/neelance/dom/prop"
	"github.com/neelance/dom/style"
)

func Page(model *model.Model) dom.Aspect {
	return elem.Div(
		greeting(model),
		items(model),
	)
}

func greeting(model *model.Model) dom.Aspect {
	return elem.Div(
		dom.Text("Your name: "),
		elem.Input(
			prop.Type("text"),
			bind.Value(&model.Name, model.Scope),
		),
		elem.H1(
			style.Color("blue"),
			dom.Text("Hello "),
			bind.TextPtr(&model.Name, model.Scope),
			dom.Text("!"),
		),
	)
}

func items(model *model.Model) dom.Aspect {
	return elem.Div(
		elem.UL(
			bind.Dynamic(model.Scope, func(aspects *bind.Aspects) {
				for _, item := range model.Items {
					if !aspects.Reuse(item) {
						aspects.Add(item, elem.LI(
							bind.TextPtr(&item.Label, model.Scope),
							dom.Text(" "),
							elem.Button(
								dom.Text("Delete"),
								event.Click(model.DeleteItem(item)),
							),
						))
					}
				}
			}),
		),

		elem.Button(
			dom.Text("Append item"),
			event.Click(model.AppendItem),
		),
		elem.Button(
			dom.Text("Prepend item"),
			event.Click(model.PrependItem),
		),
	)
}