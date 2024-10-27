package types

const (
	ModuleName   = "intertx"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
	PortID       = ModuleName
	Version      = "ics27-1"
)

var (
	KeyPrefixButlers     = []byte{0x01}
	KeyPrefixQueue       = []byte{0x02}
	KeyPrefixNextQueueId = []byte{0x03}
)

const (
	MEMO_TRANSFER_TOKEN = "memo_transfer_token"
	MEMO_TRANSFER_PAIR  = "memo_transfer_pair"
)
