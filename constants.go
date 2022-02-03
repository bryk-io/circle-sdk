package circlesdk

// SupportedCurrency in Circle.
type SupportedCurrency string

// SupportedChain in Circle.
type SupportedChain string

const (
	// USD = USDC stablecoin.
	USD SupportedCurrency = "USD"

	// BTC = Bitcoin.
	BTC SupportedCurrency = "BTC"

	// ETH = Ethereum.
	ETH SupportedCurrency = "ETH"
)

const (
	// ChainALGO = Algorand blockchain.
	ChainALGO SupportedChain = "ALGO"

	// ChainAVAX = Avalanche blockchain.
	ChainAVAX SupportedChain = "AVAX"

	// ChainBTC = Bitcoin blockchain.
	ChainBTC SupportedChain = "BTC"

	// ChainETH = Ethereum blockchain.
	ChainETH SupportedChain = "ETH"

	// ChainFLOW = Flow blockchain.
	ChainFLOW SupportedChain = "FLOW"

	// ChainHBAR = Hedera Hash graph.
	ChainHBAR SupportedChain = "HBAR"

	// ChainSOL = Solana blockchain.
	ChainSOL SupportedChain = "SOL"

	// ChainTRX = TRON blockchain.
	ChainTRX SupportedChain = "TRX"

	// ChainXLM = Stellar blockchain.
	ChainXLM SupportedChain = "XLM"
)
