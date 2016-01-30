package btcinterned

import (
	"time"

	"github.com/btcsuite/btcrpcclient"
	"github.com/btcsuite/btcutil"
)

const SixMonths = 180 * 24 * time.Hour

// LookupAddress returns the minimum balance of the specified address for the last dur amount of time.
func LookupAddress(c *btcrpcclient.Client, addr btcutil.Address, dur time.Duration) (float64, error) {
	more := true
	offset := 0
	for more {
		more = false
		txs, err := c.SearchRawTransactionsVerbose(addr, offset, 1000, false, false, nil)
		if err != nil {
			return 0, err
		}
		if len(txs) == 1000 {
			more = true
			offset += 1000
		}
		for _, tx := range txs {
			_ = tx
		}
	}
	return 0, nil
}
