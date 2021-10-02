package flag

var opts options

type options map[string]string

func (t *options) Map() map[string]string {
	return *t
}

func Options() *options {
	return &opts
}
