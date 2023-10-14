package generic

type Label[L, A any, S any] struct {
}

type Labeler[L, A any, S any] interface {
	Label(payload []byte, initial Label[L, A, S]) (Label[L, A, S], error)
}

type PIILabeler[L, A any, S any] interface {
	Labeler[L, A, S]
	Skip(payload []byte, initial Label[L, A, S]) (Label[L, A, S], error)
}

func wrap[L, A any, S any](detector Labeler[L, A, S]) {
	return
}

type xxLabeler struct {
}

func (x *xxLabeler) Label(payload []byte, initial Label[string, string, string]) (Label[string, string, string], error) {
	//TODO implement me
	panic("implement me")
}

func (x *xxLabeler) Skip(payload []byte, initial Label[string, string, string]) (Label[string, string, string], error) {
	//TODO implement me
	panic("implement me")
}

func NewLabeler() PIILabeler[string, string, string] {
	return &xxLabeler{}
}
