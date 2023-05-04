// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/domain/boundaries/repository.go

// Package mock_boundaries is a generated GoMock package.
package mock_boundaries

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pagination "spells.tvblackman1.ru/lib/pagination"
	dto "spells.tvblackman1.ru/pkg/domain/dto"
)

// MockTagsRepository is a mock of TagsRepository interface.
type MockTagsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTagsRepositoryMockRecorder
}

// MockTagsRepositoryMockRecorder is the mock recorder for MockTagsRepository.
type MockTagsRepositoryMockRecorder struct {
	mock *MockTagsRepository
}

// NewMockTagsRepository creates a new mock instance.
func NewMockTagsRepository(ctrl *gomock.Controller) *MockTagsRepository {
	mock := &MockTagsRepository{ctrl: ctrl}
	mock.recorder = &MockTagsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTagsRepository) EXPECT() *MockTagsRepositoryMockRecorder {
	return m.recorder
}

// MockUsersRepository is a mock of UsersRepository interface.
type MockUsersRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUsersRepositoryMockRecorder
}

// MockUsersRepositoryMockRecorder is the mock recorder for MockUsersRepository.
type MockUsersRepositoryMockRecorder struct {
	mock *MockUsersRepository
}

// NewMockUsersRepository creates a new mock instance.
func NewMockUsersRepository(ctrl *gomock.Controller) *MockUsersRepository {
	mock := &MockUsersRepository{ctrl: ctrl}
	mock.recorder = &MockUsersRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsersRepository) EXPECT() *MockUsersRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUsersRepository) CreateUser(dto dto.UserToRepositoryDto) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", dto)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUsersRepositoryMockRecorder) CreateUser(dto interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUsersRepository)(nil).CreateUser), dto)
}

// GetById mocks base method.
func (m *MockUsersRepository) GetById(id dto.UserId) (dto.UserDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(dto.UserDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockUsersRepositoryMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockUsersRepository)(nil).GetById), id)
}

// GetUsers mocks base method.
func (m *MockUsersRepository) GetUsers(params dto.SearchUserDto) ([]dto.UserDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", params)
	ret0, _ := ret[0].([]dto.UserDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockUsersRepositoryMockRecorder) GetUsers(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUsersRepository)(nil).GetUsers), params)
}

// MockSourcesRepository is a mock of SourcesRepository interface.
type MockSourcesRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSourcesRepositoryMockRecorder
}

// MockSourcesRepositoryMockRecorder is the mock recorder for MockSourcesRepository.
type MockSourcesRepositoryMockRecorder struct {
	mock *MockSourcesRepository
}

// NewMockSourcesRepository creates a new mock instance.
func NewMockSourcesRepository(ctrl *gomock.Controller) *MockSourcesRepository {
	mock := &MockSourcesRepository{ctrl: ctrl}
	mock.recorder = &MockSourcesRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSourcesRepository) EXPECT() *MockSourcesRepositoryMockRecorder {
	return m.recorder
}

// CreateSource mocks base method.
func (m *MockSourcesRepository) CreateSource(sourceDto dto.SourceToRepositoryDto) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSource", sourceDto)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSource indicates an expected call of CreateSource.
func (mr *MockSourcesRepositoryMockRecorder) CreateSource(sourceDto interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSource", reflect.TypeOf((*MockSourcesRepository)(nil).CreateSource), sourceDto)
}

// GetById mocks base method.
func (m *MockSourcesRepository) GetById(id dto.SourceId) (dto.SourceDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(dto.SourceDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockSourcesRepositoryMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockSourcesRepository)(nil).GetById), id)
}

// GetSources mocks base method.
func (m *MockSourcesRepository) GetSources(userId dto.UserId, params dto.SearchSourceDto) ([]dto.SourceDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSources", userId, params)
	ret0, _ := ret[0].([]dto.SourceDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSources indicates an expected call of GetSources.
func (mr *MockSourcesRepositoryMockRecorder) GetSources(userId, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSources", reflect.TypeOf((*MockSourcesRepository)(nil).GetSources), userId, params)
}

// MockSetsRepository is a mock of SetsRepository interface.
type MockSetsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSetsRepositoryMockRecorder
}

// MockSetsRepositoryMockRecorder is the mock recorder for MockSetsRepository.
type MockSetsRepositoryMockRecorder struct {
	mock *MockSetsRepository
}

// NewMockSetsRepository creates a new mock instance.
func NewMockSetsRepository(ctrl *gomock.Controller) *MockSetsRepository {
	mock := &MockSetsRepository{ctrl: ctrl}
	mock.recorder = &MockSetsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSetsRepository) EXPECT() *MockSetsRepositoryMockRecorder {
	return m.recorder
}

// CreateSet mocks base method.
func (m *MockSetsRepository) CreateSet(setDto dto.SetToRepositoryDto) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateSet", setDto)
}

// CreateSet indicates an expected call of CreateSet.
func (mr *MockSetsRepositoryMockRecorder) CreateSet(setDto interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSet", reflect.TypeOf((*MockSetsRepository)(nil).CreateSet), setDto)
}

// EditSpellComments mocks base method.
func (m *MockSetsRepository) EditSpellComments(id dto.SetSpellId, setDto dto.EditSpellInSetDto) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EditSpellComments", id, setDto)
}

// EditSpellComments indicates an expected call of EditSpellComments.
func (mr *MockSetsRepositoryMockRecorder) EditSpellComments(id, setDto interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditSpellComments", reflect.TypeOf((*MockSetsRepository)(nil).EditSpellComments), id, setDto)
}

// GetById mocks base method.
func (m *MockSetsRepository) GetById(id dto.SetId) dto.SetDto {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(dto.SetDto)
	return ret0
}

// GetById indicates an expected call of GetById.
func (mr *MockSetsRepositoryMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockSetsRepository)(nil).GetById), id)
}

// GetSetsByName mocks base method.
func (m *MockSetsRepository) GetSetsByName(name string) []dto.SetDto {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSetsByName", name)
	ret0, _ := ret[0].([]dto.SetDto)
	return ret0
}

// GetSetsByName indicates an expected call of GetSetsByName.
func (mr *MockSetsRepositoryMockRecorder) GetSetsByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSetsByName", reflect.TypeOf((*MockSetsRepository)(nil).GetSetsByName), name)
}

// GetSpell mocks base method.
func (m *MockSetsRepository) GetSpell(params dto.SetSpellId) dto.SetSpellDto {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSpell", params)
	ret0, _ := ret[0].(dto.SetSpellDto)
	return ret0
}

// GetSpell indicates an expected call of GetSpell.
func (mr *MockSetsRepositoryMockRecorder) GetSpell(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSpell", reflect.TypeOf((*MockSetsRepository)(nil).GetSpell), params)
}

// GetSpells mocks base method.
func (m *MockSetsRepository) GetSpells(id dto.SetId, params dto.SearchSpellDto) []dto.SetSpellDto {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSpells", id, params)
	ret0, _ := ret[0].([]dto.SetSpellDto)
	return ret0
}

// GetSpells indicates an expected call of GetSpells.
func (mr *MockSetsRepositoryMockRecorder) GetSpells(id, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSpells", reflect.TypeOf((*MockSetsRepository)(nil).GetSpells), id, params)
}

// UpdateSpellList mocks base method.
func (m *MockSetsRepository) UpdateSpellList(id dto.SetId, dto dto.UpdateSetSpellListDto) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateSpellList", id, dto)
}

// UpdateSpellList indicates an expected call of UpdateSpellList.
func (mr *MockSetsRepositoryMockRecorder) UpdateSpellList(id, dto interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSpellList", reflect.TypeOf((*MockSetsRepository)(nil).UpdateSpellList), id, dto)
}

// MockSpellsRepository is a mock of SpellsRepository interface.
type MockSpellsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSpellsRepositoryMockRecorder
}

// MockSpellsRepositoryMockRecorder is the mock recorder for MockSpellsRepository.
type MockSpellsRepositoryMockRecorder struct {
	mock *MockSpellsRepository
}

// NewMockSpellsRepository creates a new mock instance.
func NewMockSpellsRepository(ctrl *gomock.Controller) *MockSpellsRepository {
	mock := &MockSpellsRepository{ctrl: ctrl}
	mock.recorder = &MockSpellsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpellsRepository) EXPECT() *MockSpellsRepositoryMockRecorder {
	return m.recorder
}

// CreateSpell mocks base method.
func (m *MockSpellsRepository) CreateSpell(spellDto dto.SpellToRepositoryDto) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSpell", spellDto)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSpell indicates an expected call of CreateSpell.
func (mr *MockSpellsRepositoryMockRecorder) CreateSpell(spellDto interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSpell", reflect.TypeOf((*MockSpellsRepository)(nil).CreateSpell), spellDto)
}

// GetById mocks base method.
func (m *MockSpellsRepository) GetById(id dto.SpellId) dto.SpellDto {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(dto.SpellDto)
	return ret0
}

// GetById indicates an expected call of GetById.
func (mr *MockSpellsRepositoryMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockSpellsRepository)(nil).GetById), id)
}

// GetSpells mocks base method.
func (m *MockSpellsRepository) GetSpells(params dto.SearchSpellDto, pagination pagination.Pagination) ([]dto.SpellDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSpells", params, pagination)
	ret0, _ := ret[0].([]dto.SpellDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSpells indicates an expected call of GetSpells.
func (mr *MockSpellsRepositoryMockRecorder) GetSpells(params, pagination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSpells", reflect.TypeOf((*MockSpellsRepository)(nil).GetSpells), params, pagination)
}
