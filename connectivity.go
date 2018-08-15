package storageos

import (
	"encoding/json"
	"net/http"
	"path"

	"github.com/storageos/go-api/types"
)

var (
	// NetworkDiagnosticsAPIPrefix is a partial path to the HTTP endpoint for
	// the node connectivity diagnostics report.
	NetworkDiagnosticsAPIPrefix = "diagnostics/network"
)

// Connectivity returns a node by its reference.
func (c *Client) Connectivity(ref string) ([]types.ConnectivityResult, error) {
	resp, err := c.do("GET", path.Join(NetworkDiagnosticsAPIPrefix, ref), doOptions{})
	if err != nil {
		if e, ok := err.(*Error); ok && e.Status == http.StatusNotFound {
			return nil, ErrNoSuchNode
		}
		return nil, err
	}
	defer resp.Body.Close()

	var nodecon []types.ConnectivityResult
	if err := json.NewDecoder(resp.Body).Decode(&nodecon); err != nil {
		return nil, err
	}
	return nodecon, nil
}
