package analyze

import (
	"testing"

	"github.com/michaeldelali/contractfault/internal/contract"
)

// build constructs a minimal contract from endpoints and types for tests.
func build(version string, types map[string]contract.Type, eps ...contract.Endpoint) *contract.Contract {
	c := &contract.Contract{Service: "svc", Version: version, Types: types, Endpoints: eps}
	return c
}

func find(d *Diff, code string) *Change {
	for i := range d.Changes {
		if d.Changes[i].Code == code {
			return &d.Changes[i]
		}
	}
	return nil
