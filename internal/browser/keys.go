package browser

type Key struct {
	Value string
	Label string
}

func (k Key) Render() string {
	inline := string(k.Label[0:1]) == k.Value

	key := KeyTextStyle.Render("(") + KeyValueStyle.Render(k.Value) + KeyTextStyle.Render(")")
	var label string
	if inline {
		label = KeyTextStyle.Render(k.Label[1:])
	} else {
		label = KeyTextStyle.Render(" " + k.Label)
	}

	return KeysStyle.Render(key + label)
}

var FilterKey Key = Key{
	Value: "f",
	Label: "filter",
}

var QuitKey Key = Key{
	Value: "q",
	Label: "quit",
}
