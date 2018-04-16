// Copyright 2017 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package sale

import (
	"github.com/labneco/doxa/doxa/models"
	"github.com/labneco/doxa/pool/h"
)

func init() {

	h.Company().AddFields(map[string]models.FieldDefinition{
		"SaleNote": models.TextField{String: "Default Terms and Conditions", Translate: true},
	})

}
