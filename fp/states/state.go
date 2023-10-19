package main

import "fmt"

type App[DataT, ReturnT any] struct {
	ret ReturnT
}

type Ptr[DataT, ValueT any] struct {
	val ValueT
}

func Malloc[DataT, ValueT any](val ValueT) *App[DataT, *Ptr[DataT, ValueT]] {
	return &App[DataT, *Ptr[DataT, ValueT]]{ret: &Ptr[DataT, ValueT]{val: val}}
}

func Deref[DataT, ValueT any](p *Ptr[DataT, ValueT]) *App[DataT, ValueT] {
	return &App[DataT, ValueT]{ret: p.val}
}

func Assign[DataT, ValueT any](p *Ptr[DataT, ValueT], val ValueT) *App[DataT, error] {
	p.val = val
	return new(App[DataT, error])
}

func Return[DataT, ReturnT any](ret ReturnT) *App[DataT, ReturnT] {
	return &App[DataT, ReturnT]{ret: ret}
}

func Then[DataT, ReturnA, ReturnB any](app *App[DataT, ReturnA], f func(a ReturnA) *App[DataT, ReturnB]) *App[DataT, ReturnB] {
	return f(app.ret)
}

func AndThen[DataT, ReturnB any](a *App[DataT, error], b *App[DataT, ReturnB]) *App[DataT, ReturnB] {
	return Then(a, func(error) *App[DataT, ReturnB] { return b })
}

func Run[DataT, ReturnT any](app *App[DataT, ReturnT]) ReturnT { return app.ret }

func swap0[DataT, ReturnT any](pa *Ptr[DataT, ReturnT], pb *Ptr[DataT, ReturnT]) *App[DataT, error] {
	return Then(Deref(pa), func(a ReturnT) *App[DataT, error] {
		return Then(Deref(pb), func(b ReturnT) *App[DataT, error] {
			return Then(Assign(pa, b), func(error) *App[DataT, error] {
				return Assign(pb, a)
			})
		})
	})
}

func swap1[DataT, ReturnT any](pa *Ptr[DataT, ReturnT], pb *Ptr[DataT, ReturnT]) *App[DataT, error] {
	return Then(Deref(pa), func(a ReturnT) *App[DataT, error] {
		return Then(Deref(pb), func(b ReturnT) *App[DataT, error] {
			return AndThen(Assign(pa, b), Assign(pb, a))
		})
	})
}

type RealWorld struct{}

func Print[T any](a T) *App[RealWorld, error] {
	fmt.Println(a)
	return new(App[RealWorld, error])
}

func main() {
	if err := Run(
		Then(
			Malloc[RealWorld, int](42),
			func(p *Ptr[RealWorld, int]) *App[RealWorld, error] {
				return Then(
					Deref(p),
					func(n int) *App[RealWorld, error] {
						return AndThen(
							Print(n),
							Assign(p, n*2),
						)
					},
				)
			},
		),
	); err != nil {
		panic(err)
	}
}
