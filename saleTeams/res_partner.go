// Copyright 2017 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package saleTeams

import (
	"github.com/labneco/doxa/doxa/models"
	"github.com/labneco/doxa/pool/h"
)

func init() {

	h.Partner().AddFields(map[string]models.FieldDefinition{
		"Team": models.Many2OneField{String: "Sales Team", RelationModel: h.CRMTeam(),
			Help: "If set, sale team used notably for sales and assignations related to this partner"},
	})

}
