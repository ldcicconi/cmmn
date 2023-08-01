package translate

// PairTranslator is an interface for translating between the pair names used by the exchange
// and the pair names we use.
type PairTranslator interface {
	FromOurs(string) string
	ToOurs(string) string
}

func NewPairTranslator(exchange string) PairTranslator {
	constructors := map[string]func() PairTranslator{
		"bitflyer": newBitflyerTranslator,
		"kraken":   newKrakenTranslator,
		"sfox":     newSFOXTranslator,
	}

	constructor, ok := constructors[exchange]
	if !ok {
		panic("unknown exchange " + exchange)
	}

	return constructor()
}

func newBitflyerTranslator() PairTranslator {
	return newMapTranslator(map[string]string{
		"btcusd": "BTC_USD",
		"ethusd": "ETH_USD",
	})
}

func newKrakenTranslator() PairTranslator {
	return newMapTranslator(map[string]string{
		"btcusd": "XBT/USD",
		"ethusd": "ETH/USD",
	})
}

func newSFOXTranslator() PairTranslator {
	return NewNoopTranslator()
}
