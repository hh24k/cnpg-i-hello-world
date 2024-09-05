package operator

import (
	"context"

	"encoding/json"

	"github.com/cloudnative-pg/cnpg-i/pkg/operator"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type pluginStatus struct {
	TS      timestamppb.Timestamp
	Greeted bool
}

// SetStatusInCluster is invoked at the end of the reconciliation loop and it is used to set the plugin status
// inside the .status.pluginStatus[pluginName] map key
func (Implementation) SetStatusInCluster(
	_ context.Context,
	request *operator.SetStatusInClusterRequest,
) (*operator.SetStatusInClusterResponse, error) {

	pluginStatus := &pluginStatus{
		TS:      *timestamppb.Now(),
		Greeted: true,
	}

	pluginStatusBytes, err := json.Marshal(pluginStatus)
	if err != nil {
		return nil, err
	}

	return &operator.SetStatusInClusterResponse{JsonStatus: pluginStatusBytes}, nil
}
