package temporal

import (
	"github.com/gagliardetto/solana-go"
)

var NOZOMI_TIP = solana.MustPublicKeyFromBase58("TEMPaMeCRFAS9EKF53Jd6KpHxgL47uWLcpFArU1Fanq")
var MIN_TIP_AMOUNT uint64 = 1000000

type Region string
type SendTransactionEndpoint string

var (
	USEast_STE    SendTransactionEndpoint = "http://nozomi-preview-pit.temporal.xyz/?c="
	Frankfurt_STE SendTransactionEndpoint = "http://fra1.nozomi.temporal.xyz/?c="
	Amsterdam_STE SendTransactionEndpoint = "http://nozomi-preview-ams.temporal.xyz/?c="

	USEast    Region = "https://nozomi-preview-pit.temporal.xyz/?c="
	Frankfurt Region = "https://fra1.nozomi.temporal.xyz/?c="
	Amsterdam Region = "https://nozomi-preview-ams.temporal.xyz/?c="
)
