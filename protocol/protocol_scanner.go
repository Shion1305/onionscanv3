package protocol

import (
	"github.com/415ALS/onionscanv3/config"
	"github.com/415ALS/onionscanv3/report"
)

type Scanner interface {
	ScanProtocol(hiddenService string, onionscanConfig *config.OnionScanConfig, report *report.OnionScanReport)
}
