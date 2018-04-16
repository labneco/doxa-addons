// Copyright 2017 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package product

import (
	"github.com/labneco/doxa-addons/decimalPrecision"
	"github.com/labneco/doxa/doxa/tools/nbutils"
)

func init() {
	decimalPrecision.Precisions["Product Price"] = nbutils.Digits{Precision: 16, Scale: 2}
	decimalPrecision.Precisions["Discount"] = nbutils.Digits{Precision: 16, Scale: 2}
	decimalPrecision.Precisions["Stock Weight"] = nbutils.Digits{Precision: 16, Scale: 2}
	decimalPrecision.Precisions["Product Unit of Measure"] = nbutils.Digits{Precision: 16, Scale: 3}
}
