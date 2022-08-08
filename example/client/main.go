package main

import (
	"fmt"

	"github.com/ymzuiku/w/w"
)

func removeLoading() {
	loading := w.Document.GetElementById("first-loading")
	if loading != nil {
		loading.Remove()
	}
}

type Anime struct {
	Name string
	Si   func() int
	Age  int
}

func App() *w.Ele {
	num, setNum := w.CreateSignal(0)
	show5, setShow5 := w.CreateSignal(false)
	_, setVal := w.CreateSignal("")
	animes, setAnimes := w.CreateSignal([]*Anime{
		{Name: "the dog", Si: num, Age: 5},
		{Name: "the cat", Si: num, Age: 10},
	})
	setNumHandle := func(l int) {
		setNum(l)
		setShow5(l > 5)
	}

	handleButtonClick := func() {
		list := animes()
		list = append(list, &Anime{Name: "the fish", Age: 15})
		setAnimes(list)
	}

	handleInput := func(e *w.Ele) {
		l := len(e.GetValue())
		setVal(e.GetValue())
		setNumHandle(l)
		list := []*Anime{animes()[0]}
		for i := 0; i < l*100; i++ {
			list = append(list, &Anime{Name: fmt.Sprintf("the fish %v", i), Si: num, Age: i})
		}
		setAnimes(list)
	}

	onRemove := func() {
		w.Console.Log("on mount")
	}

	w.Div().InnerHTML(`
		<div></div>
	`)

	return w.Div().OnRemove(onRemove).Append(
		w.Div().Text(num),
		w.Button().Text("add anime").OnClick(handleButtonClick),
		w.For(animes, func(a *Anime, i int) *w.Ele {
			return w.Div().Append(
				w.If(show5, func() *w.Ele {
					return w.Div().Text("big 5")
				}),
				w.Input().Placeholder(a.Name).OnInput(handleInput).Ref(func(e *w.Ele) {
					if i > 10 && i < 20 {
						e.Val(a.Si)
					}
				}))
		}),
	)
}

func main() {
	w.BrowserRun(func() {
		removeLoading()
		w.Body.Append(App())
	})
}
