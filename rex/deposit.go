package rex

import (
	eos "github.com/eoscanada/eos-go"
)

func NewDeposit(owner eos.AccountName, amount eos.Asset) *eos.Action {
	return &eos.Action{
		Account: REXAN,
		Name:    ActN("deposit"),
		Authorization: []eos.PermissionLevel{
			{Actor: owner, Permission: eos.PermissionName("active")},
		},
		ActionData: eos.NewActionData(Deposit{
			Owner:  owner,
			Amount: amount,
		}),
	}
}

type Deposit struct {func NewDeposit(owner eos.AccountName, amount eos.Asset) (*eos.Action, error) {
    if owner == "" {
        return nil, fmt.Errorf("owner cannot be empty")
    }
    if amount.Amount <= 0 {
	Owner  eos.AccountName
	Amount eos.Asset
}
