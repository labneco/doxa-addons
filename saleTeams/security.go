// Copyright 2017 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package saleTeams

import (
	"github.com/labneco/doxa-ui/base"
	"github.com/labneco/doxa/doxa/models/security"
	"github.com/labneco/doxa/pool/h"
)

var (
	GroupSaleSalesman         *security.Group
	GroupSaleManager          *security.Group
	GroupSaleSalesmanAllLeads *security.Group
)

func init() {

	h.CRMTeam().Methods().Load().AllowGroup(base.GroupUser)
	h.CRMTeam().Methods().Load().AllowGroup(GroupSaleSalesman)
	h.CRMTeam().Methods().AllowAllToGroup(GroupSaleManager)
}
