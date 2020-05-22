package syncs

type Counter struct {
	count int
}

func (c *Counter) Inc() {
	c.count += 1
}
func (c *Counter) Value() int {
	return c.count
}
