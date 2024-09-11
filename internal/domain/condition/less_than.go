package condition

type lessThan struct {
	condition
}

func (l lessThan) Check(map[string]any) bool {
	if l.X.(int) < l.Y.(int) {
		return true
	}
	return false
}

func LessThan(key string, x int) Condition {
	return &lessThan{
		condition{
			key: key,
			X:   x,
		},
	}
}
