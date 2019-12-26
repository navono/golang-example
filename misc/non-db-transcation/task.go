package non_db_transcation

type (
	Task struct {
		Name, Group string
		Config      map[string]string
	}
)

func (t *Task) UpdateConfig(c map[string]string) map[string]string {
	for k, v := range t.Config {
		c[k] = v
	}

	return c
}
