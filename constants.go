package temporal

import (
	"math/rand"
	"time"
)

var (
	tipStreamURL = "wss://api.nozomi.temporal.xyz/tip_stream"
)

var NOZOMI_TIP_ADDRESSES = []string{
	"TEMPaMeCRFAS9EKF53Jd6KpHxgL47uWLcpFArU1Fanq",
	"noz3jAjPiHuBPqiSPkkugaJDkJscPuRhYnSpbi8UvC4",
	"noz3str9KXfpKknefHji8L1mPgimezaiUyCHYMDv1GE",
	"noz6uoYCDijhu1V7cutCpwxNiSovEwLdRHPwmgCGDNo",
	"noz9EPNcT7WH6Sou3sr3GGjHQYVkN3DNirpbvDkv9YJ",
	"nozc5yT15LazbLTFVZzoNZCwjh3yUtW86LoUyqsBu4L",
	"nozFrhfnNGoyqwVuwPAW4aaGqempx4PU6g6D9CJMv7Z",
	"nozievPk7HyK1Rqy1MPJwVQ7qQg2QoJGyP71oeDwbsu",
	"noznbgwYnBLDHu8wcQVCEw6kDrXkPdKkydGJGNXGvL7",
	"nozNVWs5N8mgzuD3qigrCG2UoKxZttxzZ85pvAQVrbP",
	"nozpEGbwx4BcGp6pvEdAh1JoC2CQGZdU6HbNP1v2p6P",
	"nozrhjhkCr3zXT3BiT4WCodYCUFeQvcdUkM7MqhKqge",
	"nozrwQtWhEdrA6W8dkbt9gnUaMs52PdAv5byipnadq3",
	"nozUacTVWub3cL4mJmGCYjKZTnE9RbdY5AP46iQgbPJ",
	"nozWCyTPppJjRuw2fpzDhhWbW355fzosWSzrrMYB1Qk",
	"nozWNju6dY353eMkMqURqwQEoM3SFgEKC6psLCSfUne",
	"nozxNBgWohjR75vdspfxR5H9ceC7XXH99xpxhVGt3Bb",
}

func PickRandomNozomiTipAddress() string {
	rand.NewSource(time.Now().UnixNano())
	return NOZOMI_TIP_ADDRESSES[rand.Intn(len(NOZOMI_TIP_ADDRESSES))]
}

var MIN_TIP_AMOUNT uint64 = 1000000

type Region string

var (
	PITT_HTTP Region = "http://pit1.nozomi.temporal.xyz/?c="
	FRA_HTTP  Region = "http://fra2.nozomi.temporal.xyz/?c="
	EWR_HTTP  Region = "http://ewr1.nozomi.temporal.xyz/?c="
	AMS_HTTP  Region = "http://ams1.nozomi.temporal.xyz/?c="

	PITT_HTTPS Region = "https://pit1.secure.nozomi.temporal.xyz/?c="
	FRA_HTTPS  Region = "https://fra2.secure.nozomi.temporal.xyz/?c="
	EWR_HTTPS  Region = "https://ewr1.secure.nozomi.temporal.xyz/?c="
	AMS_HTTPS  Region = "https://ams1.secure.nozomi.temporal.xyz/?c="
)
