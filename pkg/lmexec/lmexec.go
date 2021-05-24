package lmexec

import (
	"github.com/logicmonitor/k8s-argus/pkg/types"
	"github.com/logicmonitor/lm-sdk-go/client/lm"
	"github.com/logicmonitor/lm-sdk-go/models"
)

// LMExec Provides utility function for SDK calls using Base object
// LMExec is holding device related api calls at the moment to mitigate rate limit handling
type LMExec struct {
	*types.Base
}

// AddDevice Add new device
func (lmexec *LMExec) AddDevice(params *lm.AddDeviceParams) types.ExecRequest {
	return func() (interface{}, error) {
		return lmexec.LMClient.LM.AddDevice(params)
	}
}

// UpdateDevice Add new device
func (lmexec *LMExec) UpdateDevice(params *lm.UpdateDeviceParams) types.ExecRequest {
	return func() (interface{}, error) {
		return lmexec.LMClient.LM.UpdateDevice(params)
	}
}

// GetDeviceByID get device by id
func (lmexec *LMExec) GetDeviceByID(params *lm.GetDeviceByIDParams) types.ExecRequest {
	return func() (interface{}, error) {
		return lmexec.LMClient.LM.GetDeviceByID(params)
	}
}

// GetDeviceList Add new device
func (lmexec *LMExec) GetDeviceList(params *lm.GetDeviceListParams) types.ExecRequest {
	return func() (interface{}, error) {
		return lmexec.LMClient.LM.GetDeviceList(params)
	}
}

// PatchDevice Add new device
func (lmexec *LMExec) PatchDevice(params *lm.PatchDeviceParams) types.ExecRequest {
	return func() (interface{}, error) {
		return lmexec.LMClient.LM.PatchDevice(params)
	}
}

// DeleteDeviceByID Add new device
func (lmexec *LMExec) DeleteDeviceByID(params *lm.DeleteDeviceByIDParams) types.ExecRequest {
	return func() (interface{}, error) {
		return lmexec.LMClient.LM.DeleteDeviceByID(params)
	}
}

// GetImmediateDeviceListByDeviceGroupID Add new device
func (lmexec *LMExec) GetImmediateDeviceListByDeviceGroupID(params *lm.GetImmediateDeviceListByDeviceGroupIDParams) types.ExecRequest {
	return func() (interface{}, error) {
		return lmexec.LMClient.LM.GetImmediateDeviceListByDeviceGroupID(params)
	}
}

// AddDeviceErrResp parse error object and returns models.ErrorResponse
func (lmexec *LMExec) AddDeviceErrResp(err error) *models.ErrorResponse {
	return err.(*lm.AddDeviceDefault).Payload // nolint: errorlint
}

// UpdateDeviceErrResp parse error object and returns models.ErrorResponse
func (lmexec *LMExec) UpdateDeviceErrResp(err error) *models.ErrorResponse {
	return err.(*lm.UpdateDeviceDefault).Payload // nolint: errorlint
}

// GetDeviceByIDErrResp parse error object and returns models.ErrorResponse
func (lmexec *LMExec) GetDeviceByIDErrResp(err error) *models.ErrorResponse {
	return err.(*lm.GetDeviceByIDDefault).Payload // nolint: errorlint
}

// UpdateDevicePropertyErrResp parse error object and returns models.ErrorResponse
func (lmexec *LMExec) UpdateDevicePropertyErrResp(err error) *models.ErrorResponse {
	return err.(*lm.UpdateDevicePropertyByNameDefault).Payload // nolint: errorlint
}

// UpdateDevicePropertyByName updates specified device property.
func (lmexec *LMExec) UpdateDevicePropertyByName(params *lm.UpdateDevicePropertyByNameParams) types.ExecRequest {
	return func() (interface{}, error) {
		return lmexec.LMClient.LM.UpdateDevicePropertyByName(params)
	}
}

// GetDeviceListErrResp parse error object and returns models.ErrorResponse
func (lmexec *LMExec) GetDeviceListErrResp(err error) *models.ErrorResponse {
	return err.(*lm.GetDeviceListDefault).Payload // nolint: errorlint
}

// PatchDeviceErrResp parse error object and returns models.ErrorResponse
func (lmexec *LMExec) PatchDeviceErrResp(err error) *models.ErrorResponse {
	return err.(*lm.PatchDeviceDefault).Payload // nolint: errorlint
}

// DeleteDeviceByIDErrResp parse error object and returns models.ErrorResponse
func (lmexec *LMExec) DeleteDeviceByIDErrResp(err error) *models.ErrorResponse {
	return err.(*lm.DeleteDeviceByIDDefault).Payload // nolint: errorlint
}

// GetImmediateDeviceListByDeviceGroupIDErrResp parse error object and returns models.ErrorResponse
func (lmexec *LMExec) GetImmediateDeviceListByDeviceGroupIDErrResp(err error) *models.ErrorResponse {
	return err.(*lm.GetImmediateDeviceListByDeviceGroupIDDefault).Payload // nolint: errorlint
}
