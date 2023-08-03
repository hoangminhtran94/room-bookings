package form

type errors map[string][]string

func (e errors) Add(field string, message string) {
	e[field] = append(e[field], message)
}
func (e errors) Get(field string) string {
	errors := e[field]
	if len(errors) == 0 {
		return ""
	}
	return errors[0]
}