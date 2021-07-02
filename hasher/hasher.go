package hasher

type API interface {
	MakeSHA256_FromString(string) (string, error)
}

type Hasher struct{}

var _ API = (*Hasher)(nil)

func New() *Hasher {
	return &Hasher{}
}
