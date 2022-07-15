// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cubefs/cubefs/blobstore/scheduler/db (interfaces: IKafkaOffsetTable,IOrphanShardTable,IArchiveTable,IMigrateTaskTable,IRepairTaskTable,IInspectCheckPointTable)

// Package scheduler is a generated GoMock package.
package scheduler

import (
	context "context"
	reflect "reflect"

	proto "github.com/cubefs/cubefs/blobstore/common/proto"
	db "github.com/cubefs/cubefs/blobstore/scheduler/db"
	gomock "github.com/golang/mock/gomock"
)

// MockKafkaOffsetTable is a mock of IKafkaOffsetTable interface.
type MockKafkaOffsetTable struct {
	ctrl     *gomock.Controller
	recorder *MockKafkaOffsetTableMockRecorder
}

// MockKafkaOffsetTableMockRecorder is the mock recorder for MockKafkaOffsetTable.
type MockKafkaOffsetTableMockRecorder struct {
	mock *MockKafkaOffsetTable
}

// NewMockKafkaOffsetTable creates a new mock instance.
func NewMockKafkaOffsetTable(ctrl *gomock.Controller) *MockKafkaOffsetTable {
	mock := &MockKafkaOffsetTable{ctrl: ctrl}
	mock.recorder = &MockKafkaOffsetTableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKafkaOffsetTable) EXPECT() *MockKafkaOffsetTableMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockKafkaOffsetTable) Get(arg0 string, arg1 int32) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockKafkaOffsetTableMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockKafkaOffsetTable)(nil).Get), arg0, arg1)
}

// Set mocks base method.
func (m *MockKafkaOffsetTable) Set(arg0 string, arg1 int32, arg2 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockKafkaOffsetTableMockRecorder) Set(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockKafkaOffsetTable)(nil).Set), arg0, arg1, arg2)
}

// MockOrphanShardTable is a mock of IOrphanShardTable interface.
type MockOrphanShardTable struct {
	ctrl     *gomock.Controller
	recorder *MockOrphanShardTableMockRecorder
}

// MockOrphanShardTableMockRecorder is the mock recorder for MockOrphanShardTable.
type MockOrphanShardTableMockRecorder struct {
	mock *MockOrphanShardTable
}

// NewMockOrphanShardTable creates a new mock instance.
func NewMockOrphanShardTable(ctrl *gomock.Controller) *MockOrphanShardTable {
	mock := &MockOrphanShardTable{ctrl: ctrl}
	mock.recorder = &MockOrphanShardTableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrphanShardTable) EXPECT() *MockOrphanShardTableMockRecorder {
	return m.recorder
}

// Save mocks base method.
func (m *MockOrphanShardTable) Save(arg0 db.OrphanShard) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockOrphanShardTableMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockOrphanShardTable)(nil).Save), arg0)
}

// MockArchiveTable is a mock of IArchiveTable interface.
type MockArchiveTable struct {
	ctrl     *gomock.Controller
	recorder *MockArchiveTableMockRecorder
}

// MockArchiveTableMockRecorder is the mock recorder for MockArchiveTable.
type MockArchiveTableMockRecorder struct {
	mock *MockArchiveTable
}

// NewMockArchiveTable creates a new mock instance.
func NewMockArchiveTable(ctrl *gomock.Controller) *MockArchiveTable {
	mock := &MockArchiveTable{ctrl: ctrl}
	mock.recorder = &MockArchiveTableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArchiveTable) EXPECT() *MockArchiveTableMockRecorder {
	return m.recorder
}

// FindTask mocks base method.
func (m *MockArchiveTable) FindTask(arg0 context.Context, arg1 string) (*proto.ArchiveRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindTask", arg0, arg1)
	ret0, _ := ret[0].(*proto.ArchiveRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindTask indicates an expected call of FindTask.
func (mr *MockArchiveTableMockRecorder) FindTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindTask", reflect.TypeOf((*MockArchiveTable)(nil).FindTask), arg0, arg1)
}

// Insert mocks base method.
func (m *MockArchiveTable) Insert(arg0 context.Context, arg1 *proto.ArchiveRecord) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockArchiveTableMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockArchiveTable)(nil).Insert), arg0, arg1)
}

// MockMigrateTaskTable is a mock of IMigrateTaskTable interface.
type MockMigrateTaskTable struct {
	ctrl     *gomock.Controller
	recorder *MockMigrateTaskTableMockRecorder
}

// MockMigrateTaskTableMockRecorder is the mock recorder for MockMigrateTaskTable.
type MockMigrateTaskTableMockRecorder struct {
	mock *MockMigrateTaskTable
}

// NewMockMigrateTaskTable creates a new mock instance.
func NewMockMigrateTaskTable(ctrl *gomock.Controller) *MockMigrateTaskTable {
	mock := &MockMigrateTaskTable{ctrl: ctrl}
	mock.recorder = &MockMigrateTaskTableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMigrateTaskTable) EXPECT() *MockMigrateTaskTableMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockMigrateTaskTable) Delete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMigrateTaskTableMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMigrateTaskTable)(nil).Delete), arg0, arg1)
}

// Find mocks base method.
func (m *MockMigrateTaskTable) Find(arg0 context.Context, arg1 string) (*proto.MigrateTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*proto.MigrateTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockMigrateTaskTableMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMigrateTaskTable)(nil).Find), arg0, arg1)
}

// FindAll mocks base method.
func (m *MockMigrateTaskTable) FindAll(arg0 context.Context) ([]*proto.MigrateTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", arg0)
	ret0, _ := ret[0].([]*proto.MigrateTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockMigrateTaskTableMockRecorder) FindAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockMigrateTaskTable)(nil).FindAll), arg0)
}

// FindByDiskID mocks base method.
func (m *MockMigrateTaskTable) FindByDiskID(arg0 context.Context, arg1 proto.DiskID) ([]*proto.MigrateTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByDiskID", arg0, arg1)
	ret0, _ := ret[0].([]*proto.MigrateTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByDiskID indicates an expected call of FindByDiskID.
func (mr *MockMigrateTaskTableMockRecorder) FindByDiskID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByDiskID", reflect.TypeOf((*MockMigrateTaskTable)(nil).FindByDiskID), arg0, arg1)
}

// Insert mocks base method.
func (m *MockMigrateTaskTable) Insert(arg0 context.Context, arg1 *proto.MigrateTask) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockMigrateTaskTableMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockMigrateTaskTable)(nil).Insert), arg0, arg1)
}

// MarkDeleteByDiskID mocks base method.
func (m *MockMigrateTaskTable) MarkDeleteByDiskID(arg0 context.Context, arg1 proto.DiskID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkDeleteByDiskID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkDeleteByDiskID indicates an expected call of MarkDeleteByDiskID.
func (mr *MockMigrateTaskTableMockRecorder) MarkDeleteByDiskID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkDeleteByDiskID", reflect.TypeOf((*MockMigrateTaskTable)(nil).MarkDeleteByDiskID), arg0, arg1)
}

// MarkDeleteByStates mocks base method.
func (m *MockMigrateTaskTable) MarkDeleteByStates(arg0 context.Context, arg1 []proto.MigrateState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkDeleteByStates", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkDeleteByStates indicates an expected call of MarkDeleteByStates.
func (mr *MockMigrateTaskTableMockRecorder) MarkDeleteByStates(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkDeleteByStates", reflect.TypeOf((*MockMigrateTaskTable)(nil).MarkDeleteByStates), arg0, arg1)
}

// Name mocks base method.
func (m *MockMigrateTaskTable) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockMigrateTaskTableMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockMigrateTaskTable)(nil).Name))
}

// QueryMarkDeleteTasks mocks base method.
func (m *MockMigrateTaskTable) QueryMarkDeleteTasks(arg0 context.Context, arg1 int) ([]*proto.ArchiveRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryMarkDeleteTasks", arg0, arg1)
	ret0, _ := ret[0].([]*proto.ArchiveRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryMarkDeleteTasks indicates an expected call of QueryMarkDeleteTasks.
func (mr *MockMigrateTaskTableMockRecorder) QueryMarkDeleteTasks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryMarkDeleteTasks", reflect.TypeOf((*MockMigrateTaskTable)(nil).QueryMarkDeleteTasks), arg0, arg1)
}

// RemoveMarkDelete mocks base method.
func (m *MockMigrateTaskTable) RemoveMarkDelete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveMarkDelete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveMarkDelete indicates an expected call of RemoveMarkDelete.
func (mr *MockMigrateTaskTableMockRecorder) RemoveMarkDelete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMarkDelete", reflect.TypeOf((*MockMigrateTaskTable)(nil).RemoveMarkDelete), arg0, arg1)
}

// Update mocks base method.
func (m *MockMigrateTaskTable) Update(arg0 context.Context, arg1 proto.MigrateState, arg2 *proto.MigrateTask) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockMigrateTaskTableMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMigrateTaskTable)(nil).Update), arg0, arg1, arg2)
}

// MockRepairTaskTable is a mock of IRepairTaskTable interface.
type MockRepairTaskTable struct {
	ctrl     *gomock.Controller
	recorder *MockRepairTaskTableMockRecorder
}

// MockRepairTaskTableMockRecorder is the mock recorder for MockRepairTaskTable.
type MockRepairTaskTableMockRecorder struct {
	mock *MockRepairTaskTable
}

// NewMockRepairTaskTable creates a new mock instance.
func NewMockRepairTaskTable(ctrl *gomock.Controller) *MockRepairTaskTable {
	mock := &MockRepairTaskTable{ctrl: ctrl}
	mock.recorder = &MockRepairTaskTableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepairTaskTable) EXPECT() *MockRepairTaskTableMockRecorder {
	return m.recorder
}

// Find mocks base method.
func (m *MockRepairTaskTable) Find(arg0 context.Context, arg1 string) (*proto.MigrateTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*proto.MigrateTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockRepairTaskTableMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockRepairTaskTable)(nil).Find), arg0, arg1)
}

// FindAll mocks base method.
func (m *MockRepairTaskTable) FindAll(arg0 context.Context) ([]*proto.MigrateTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", arg0)
	ret0, _ := ret[0].([]*proto.MigrateTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockRepairTaskTableMockRecorder) FindAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockRepairTaskTable)(nil).FindAll), arg0)
}

// FindByDiskID mocks base method.
func (m *MockRepairTaskTable) FindByDiskID(arg0 context.Context, arg1 proto.DiskID) ([]*proto.MigrateTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByDiskID", arg0, arg1)
	ret0, _ := ret[0].([]*proto.MigrateTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByDiskID indicates an expected call of FindByDiskID.
func (mr *MockRepairTaskTableMockRecorder) FindByDiskID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByDiskID", reflect.TypeOf((*MockRepairTaskTable)(nil).FindByDiskID), arg0, arg1)
}

// Insert mocks base method.
func (m *MockRepairTaskTable) Insert(arg0 context.Context, arg1 *proto.MigrateTask) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockRepairTaskTableMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRepairTaskTable)(nil).Insert), arg0, arg1)
}

// MarkDeleteByDiskID mocks base method.
func (m *MockRepairTaskTable) MarkDeleteByDiskID(arg0 context.Context, arg1 proto.DiskID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkDeleteByDiskID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkDeleteByDiskID indicates an expected call of MarkDeleteByDiskID.
func (mr *MockRepairTaskTableMockRecorder) MarkDeleteByDiskID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkDeleteByDiskID", reflect.TypeOf((*MockRepairTaskTable)(nil).MarkDeleteByDiskID), arg0, arg1)
}

// Name mocks base method.
func (m *MockRepairTaskTable) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockRepairTaskTableMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockRepairTaskTable)(nil).Name))
}

// QueryMarkDeleteTasks mocks base method.
func (m *MockRepairTaskTable) QueryMarkDeleteTasks(arg0 context.Context, arg1 int) ([]*proto.ArchiveRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryMarkDeleteTasks", arg0, arg1)
	ret0, _ := ret[0].([]*proto.ArchiveRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryMarkDeleteTasks indicates an expected call of QueryMarkDeleteTasks.
func (mr *MockRepairTaskTableMockRecorder) QueryMarkDeleteTasks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryMarkDeleteTasks", reflect.TypeOf((*MockRepairTaskTable)(nil).QueryMarkDeleteTasks), arg0, arg1)
}

// RemoveMarkDelete mocks base method.
func (m *MockRepairTaskTable) RemoveMarkDelete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveMarkDelete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveMarkDelete indicates an expected call of RemoveMarkDelete.
func (mr *MockRepairTaskTableMockRecorder) RemoveMarkDelete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMarkDelete", reflect.TypeOf((*MockRepairTaskTable)(nil).RemoveMarkDelete), arg0, arg1)
}

// Update mocks base method.
func (m *MockRepairTaskTable) Update(arg0 context.Context, arg1 *proto.MigrateTask) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRepairTaskTableMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepairTaskTable)(nil).Update), arg0, arg1)
}

// MockInspectCheckPointTable is a mock of IInspectCheckPointTable interface.
type MockInspectCheckPointTable struct {
	ctrl     *gomock.Controller
	recorder *MockInspectCheckPointTableMockRecorder
}

// MockInspectCheckPointTableMockRecorder is the mock recorder for MockInspectCheckPointTable.
type MockInspectCheckPointTableMockRecorder struct {
	mock *MockInspectCheckPointTable
}

// NewMockInspectCheckPointTable creates a new mock instance.
func NewMockInspectCheckPointTable(ctrl *gomock.Controller) *MockInspectCheckPointTable {
	mock := &MockInspectCheckPointTable{ctrl: ctrl}
	mock.recorder = &MockInspectCheckPointTableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInspectCheckPointTable) EXPECT() *MockInspectCheckPointTableMockRecorder {
	return m.recorder
}

// GetCheckPoint mocks base method.
func (m *MockInspectCheckPointTable) GetCheckPoint(arg0 context.Context) (*proto.InspectCheckPoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCheckPoint", arg0)
	ret0, _ := ret[0].(*proto.InspectCheckPoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCheckPoint indicates an expected call of GetCheckPoint.
func (mr *MockInspectCheckPointTableMockRecorder) GetCheckPoint(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCheckPoint", reflect.TypeOf((*MockInspectCheckPointTable)(nil).GetCheckPoint), arg0)
}

// SaveCheckPoint mocks base method.
func (m *MockInspectCheckPointTable) SaveCheckPoint(arg0 context.Context, arg1 proto.Vid) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCheckPoint", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveCheckPoint indicates an expected call of SaveCheckPoint.
func (mr *MockInspectCheckPointTableMockRecorder) SaveCheckPoint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCheckPoint", reflect.TypeOf((*MockInspectCheckPointTable)(nil).SaveCheckPoint), arg0, arg1)
}
