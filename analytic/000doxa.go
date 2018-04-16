// Copyright 2017 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package analytic

import (
	// Import dependencies
	_ "github.com/labneco/doxa-addons/product"
	"github.com/labneco/doxa/doxa/models/security"
	"github.com/labneco/doxa/doxa/server"
)

const MODULE_NAME string = "analytic"

// GroupAnalyticAccounting is the group of the people allowed manage analytic accounting
var GroupAnalyticAccounting *security.Group

func init() {
	server.RegisterModule(&server.Module{
		Name:     MODULE_NAME,
		PostInit: func() {},
	})

	GroupAnalyticAccounting = security.Registry.NewGroup("analytic_group_analytic_accounting", "Analytic Accounting")
}
