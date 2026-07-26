package entrylist

import (
	"maps"
	"strconv"
)


func (c *Component) SetSize(width, height int) {
	c.list.SetSize(width, height)

	limit := 20

	limit = max(limit, height)

	c.loader.SetLimit(limit)

	c.loader.Load()
}

func (c *Component) Render() string {
	var items []string

	for _, entry := range c.loader.GetEntries() {
		values := map[string]string{}

		maps.Copy(values, entry.Metas)

		values["id"] = strconv.FormatInt(entry.ID, 10)
		values["path"] = entry.Path

		items = append(items, entry.Path)
	}

	c.list.SetItems(items)

	return c.list.Render()
}
