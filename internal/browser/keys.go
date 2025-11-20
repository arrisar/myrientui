package browser

type Key struct {
	Values []string
	Label  string
}

func (k Key) Render() string {
	inline := len(k.Values) == 1 && string(k.Label[0:1]) == k.Values[0]

	key := KeyTextStyle.Render("(")
	for i, v := range k.Values {
		if i > 0 {
			key += KeyTextStyle.Render("/")
		}
		key += KeyValueStyle.Render(v)
	}

	key += KeyTextStyle.Render(")")
	var label string
	if inline {
		label = KeyTextStyle.Render(k.Label[1:])
	} else {
		label = KeyTextStyle.Render(" " + k.Label)
	}

	return KeysStyle.Render(key + label)
}

var FilterKey Key = Key{
	Values: []string{"f"},
	Label:  "filter",
}

var QuitKey Key = Key{
	Values: []string{"q"},
	Label:  "quit",
}
