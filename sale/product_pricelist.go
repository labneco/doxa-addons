// Copyright 2017 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package sale

import (
	"github.com/labneco/doxa/doxa/models"
	"github.com/labneco/doxa/doxa/models/types"
	"github.com/labneco/doxa/pool/h"
)

func init() {

	h.ProductPricelist().AddFields(map[string]models.FieldDefinition{
		"DiscountPolicy": models.SelectionField{Selection: types.Selection{
			"with_discount":    "Discount included in the price",
			"without_discount": "Show public price & discount to the customer",
		}, Default: models.DefaultValue("with_discount")},
	})

}
