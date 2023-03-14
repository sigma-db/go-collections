package util

import "sync"

type Function0[R any] func() R
type Function1[T1 comparable, R any] func(T1) R
type Function2[T1, T2 comparable, R any] func(T1, T2) R
type Function3[T1, T2, T3 comparable, R any] func(T1, T2, T3) R
type Function4[T1, T2, T3, T4 comparable, R any] func(T1, T2, T3, T4) R

type FunctionAny[R any] func(...any) R

func Memoize0[R any](f Function0[R]) Function0[R] {
	// since f is required to be pure and thus returns the same value on every invocation,
	// we use a slighly easier implementation that doesn't require a sync.Map
	var result R
	var once sync.Once
	return func() R {
		once.Do(func() { result = f() })
		return result
	}
}

func Memoize1[T1 comparable, R any](f Function1[T1, R]) Function1[T1, R] {
	var cache sync.Map
	return func(t1 T1) R {
		if v, ok := cache.Load(t1); ok {
			return v.(R)
		}
		v := f(t1)
		cache.Store(t1, v)
		return v
	}
}

func Memoize2[T1, T2 comparable, R any](f Function2[T1, T2, R]) Function2[T1, T2, R] {
	var cache sync.Map
	return func(t1 T1, t2 T2) R {
		key := [...]any{t1, t2}
		if v, ok := cache.Load(key); ok {
			return v.(R)
		}
		v := f(t1, t2)
		cache.Store(key, v)
		return v
	}
}

func Memoize3[T1, T2, T3 comparable, R any](f Function3[T1, T2, T3, R]) Function3[T1, T2, T3, R] {
	var cache sync.Map
	return func(t1 T1, t2 T2, t3 T3) R {
		key := [...]any{t1, t2, t3}
		if v, ok := cache.Load(key); ok {
			return v.(R)
		}
		v := f(t1, t2, t3)
		cache.Store(key, v)
		return v
	}
}

func Memoize4[T1, T2, T3, T4 comparable, R any](f Function4[T1, T2, T3, T4, R]) Function4[T1, T2, T3, T4, R] {
	var cache sync.Map
	return func(t1 T1, t2 T2, t3 T3, t4 T4) R {
		key := [...]any{t1, t2, t3, t4}
		if v, ok := cache.Load(key); ok {
			return v.(R)
		}
		v := f(t1, t2, t3, t4)
		cache.Store(key, v)
		return v
	}
}
