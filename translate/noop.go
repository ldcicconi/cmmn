package translate

type NoopTranslator struct{}

func NewNoopTranslator() NoopTranslator {
	return NoopTranslator{}
}

func (n NoopTranslator) FromOurs(s string) string {
	return s
}

func (n NoopTranslator) ToOurs(s string) string {
	return s
}
