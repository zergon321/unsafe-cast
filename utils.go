package unsafecast

func swap[T any](a, b *T) {
	tmp := *a
	*a = *b
	*b = tmp
}
