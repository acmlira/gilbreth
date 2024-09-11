package condition

type lessThan struct {
	condition
}

func (l lessThan) Check(map[string]any) bool {
	return l.X.(int) < l.Y.(int)
}

func LessThan(key string, x int) Condition {
	return &lessThan{
		condition{
			key: key,
			X:   x,
		},
	}
}
