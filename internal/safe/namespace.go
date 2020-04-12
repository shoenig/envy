package safe

type Encrypted string

type Namespace struct {
	Name    string
	Content map[string]Encrypted
}
