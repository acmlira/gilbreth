package condition

type isPresent struct {
	condition
}

func (l isPresent) Check(variables map[string]any) bool {
	if _, ok := variables[l.key]; ok {
		return true
	}
	return false
}

func IsPresent(key string) Condition {
	return &isPresent{
		condition{
			key: key,
		},
	}
}
