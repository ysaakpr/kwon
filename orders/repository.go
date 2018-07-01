package orders

import (
	"context"
	"encoding/json"
	"time"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
)

//Repository holds all the dgraph dealings
type Repository struct {
	dg *dgo.Dgraph
}

type root struct {
	Orders []Order `json:"orders"`
}

// CreateOrder which creates order in the store
func (repo *Repository) CreateOrder(o Order) (*api.Assigned, error) {
	o.CreatedAt = time.Now()
	out, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}

	assigned, err := repo.dg.NewTxn().Mutate(context.Background(), &api.Mutation{SetJson: out, CommitNow: true})
	return assigned, err
}

//GetOrder get order deatils for uid
func (repo *Repository) GetOrder(uid string) (*Order, error) {
	q := `query Orders($id: string) {
		orders(func: uid($id)) {
			uid
			buyer
			seller
			ship_to
			bill_to
			items {
				uid
			}
		}
	}`
	// Assigned uids for nodes which were created would be returned in the resp.AssignedUids map.
	variables := map[string]string{"$id": uid}
	resp, err := repo.dg.NewTxn().QueryWithVars(context.Background(), q, variables)
	if err != nil {
		return nil, err
	}

	var r root
	err = json.Unmarshal(resp.Json, &r)
	if err != nil {
		return nil, err
	}

	return &r.Orders[0], nil
}
