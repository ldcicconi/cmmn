package translate

type mapTranslator struct {
	m        map[string]string
	mReverse map[string]string
}

func newMapTranslator(mapping map[string]string) PairTranslator {
	t := &mapTranslator{
		m:        mapping,
		mReverse: make(map[string]string),
	}
	for k, v := range mapping {
		t.mReverse[v] = k
	}
	return t
}

func (t mapTranslator) FromOurs(s string) string {
	return t.m[s]
}

func (t mapTranslator) ToOurs(s string) string {
	return t.mReverse[s]
}
