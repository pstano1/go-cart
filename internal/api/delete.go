// Package api provides a logic for the application
// This file contains definitions of delete method
package api

import "go.uber.org/zap"

// Delete deletes a specified model instance with given conditions
func (a *InstanceAPI) Delete(model interface{}, conditions ...interface{}) error {
	a.log.Debug("deleting record",
		zap.Any("from", model),
		zap.Any("conditions", conditions),
	)
	return a.dbController.Delete(model, conditions)
}
