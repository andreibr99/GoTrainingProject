package handlers

type Permission struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type UserPermissions map[string][]Permission

var Permissions = map[string]UserPermissions{
	"alice": UserPermissions{
		"payments": []Permission{
			Permission{Key: "payments.accounts.user.create", Value: true},
			Permission{Key: "payments.accounts.user.modify", Value: false},
		},
		"inventory": []Permission{
			Permission{Key: "inventory.stock.read", Value: true},
			Permission{Key: "inventory.stock.modify", Value: false},
		},
	},
	"bob": UserPermissions{
		"payments": []Permission{
			Permission{Key: "payments.accounts.user.create", Value: false},
			Permission{Key: "payments.accounts.user.modify", Value: true},
		},
		"shipping": []Permission{
			Permission{Key: "shipping.orders.create", Value: true},
			Permission{Key: "shipping.orders.modify", Value: false},
		},
	},
}
