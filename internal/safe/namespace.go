package safe

type Encrypted []byte

type Namespace struct {
	Name    string
	Content map[string]Encrypted
}
