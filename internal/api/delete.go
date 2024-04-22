package api

import "go.uber.org/zap"

func (a *InstanceAPI) DeleteUser(model interface{}, conditions ...interface{}) error {
	a.log.Debug("deleting record",
		zap.Any("from", model),
		zap.Any("conditions", conditions),
	)
	return a.dbController.Delete(model, conditions)
}
