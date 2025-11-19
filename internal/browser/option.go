package browser

/**
 * OPTION
 */

type Option string

func (o Option) IsFile() bool {
	return len(o) > 0 && !o.IsDir()
}

func (o Option) IsDir() bool {
	if len(o) < 1 {
		return false
	}

	last := string(o[len(o)-1])
	return last == "/"
}

func (o Option) Len() int {
	return len(o)
}

func (o Option) String() string {
	return string(o)
}

/**
 * OPTION
 */

type Options []Option

func (o Options) Len() int {
	return len(o)
}

func (o Options) String(i int) string {
	return string(o[i])
}
