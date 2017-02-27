// Copyright 2015, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package schemamanager

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/golang/glog"
	"golang.org/x/net/context"
)

// LiquibaseController handles schema events.
type LiquibaseController struct {
	sql      string
	keyspace string
	writer   http.ResponseWriter
}

// NewLiquibaseController creates a LiquibaseController instance
func NewLiquibaseController(
	sqlStr string, keyspace string, writer http.ResponseWriter) *LiquibaseController {
	controller := &LiquibaseController{
		keyspace: keyspace,
		writer:   writer,
		sql:      sqlStr,
	}

	return controller
}

// Open is a no-op.
func (controller *LiquibaseController) Open(ctx context.Context) error {
	return nil
}

// Read reads schema changes
func (controller *LiquibaseController) Read(ctx context.Context) ([]string, error) {
	return []string{controller.sql}, nil
}

// Close is a no-op.
func (controller *LiquibaseController) Close() {
}

// Keyspace returns keyspace to apply schema.
func (controller *LiquibaseController) Keyspace() string {
	return controller.keyspace
}

// OnReadSuccess is no-op
func (controller *LiquibaseController) OnReadSuccess(ctx context.Context) error {
	controller.writer.Write(
		[]byte(fmt.Sprintf("OnReadSuccess, sqls: %v\n", controller.sql)))
	return nil
}

// OnReadFail is no-op
func (controller *LiquibaseController) OnReadFail(ctx context.Context, err error) error {
	controller.writer.Write(
		[]byte(fmt.Sprintf("OnReadFail, error: %v\n", err)))
	return err
}

// ShouldValidate determines whether this controller expects its input to need to be validated
func (controller *LiquibaseController) ShouldValidate(ctx context.Context) bool {
	return false
}

// OnValidationSuccess is no-op
func (controller *LiquibaseController) OnValidationSuccess(ctx context.Context) error {
	controller.writer.Write(
		[]byte(fmt.Sprintf("OnValidationSuccess, sqls: %v\n", controller.sql)))
	return nil
}

// OnValidationFail is no-op
func (controller *LiquibaseController) OnValidationFail(ctx context.Context, err error) error {
	controller.writer.Write(
		[]byte(fmt.Sprintf("OnValidationFail, error: %v\n", err)))
	return err
}

// OnExecutorComplete is no-op
func (controller *LiquibaseController) OnExecutorComplete(ctx context.Context, result *ExecuteResult) error {
	data, err := json.Marshal(result)
	if err != nil {
		log.Errorf("Failed to serialize ExecuteResult: %v", err)
		return err
	}
	controller.writer.Write([]byte(fmt.Sprintf("Executor succeeds: %s", string(data))))
	return nil
}

var _ Controller = (*LiquibaseController)(nil)
