package collections

type Yield[T any] func(*T)
type Generator[T any] func(yield Yield[T])

type generatorIterator[T any] struct {
	m    Monitor
	done bool
	vp   *T
}

func (it *generatorIterator[T]) Next() bool {
	if !it.done {
		it.m.Notify()
		it.done = !it.m.Wait()
	}
	return !it.done
}

func (it *generatorIterator[T]) Iterator() (Iterator[*T], *T) {
	return it, it.vp
}

func (g Generator[T]) Iterator() (IterableIterator[*T], *T) {
	it := generatorIterator[T]{m: NewMonitor()}
	go func() {
		it.m.Wait()
		g(func(vp *T) {
			it.vp = vp
			it.m.Notify()
			it.m.Wait()
		})
		it.m.Close()
	}()
	return &it, it.vp
}

func (g Generator[T]) Stream() Stream[T] {
	it, _ := g.Iterator()
	return FromValueIterable[T](it)
}
