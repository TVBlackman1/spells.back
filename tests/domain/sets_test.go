package domain

import (
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"spells.tvblackman1.ru/pkg/domain/boundaries"
	"spells.tvblackman1.ru/pkg/domain/dto"
	"spells.tvblackman1.ru/pkg/domain/usecases"
	mock_boundaries "spells.tvblackman1.ru/tests/mocks/pkg/domain/boundaries"
	"testing"
)

func TestSetUseCase_CreateSet(t *testing.T) {
	tests := getCreateSetTests()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockSets := mock_boundaries.NewMockSetsRepository(ctrl)
			mockSets.EXPECT().
				CreateSet(gomock.Any()).
				AnyTimes()
			mockSets.EXPECT().
				GetSetsByName(gomock.Any()).
				AnyTimes().
				Return(tt.getSetsByNameReturns)
			repository := &boundaries.Repository{
				Sets: mockSets,
			}
			useCase := usecases.NewSetUseCase(repository)
			if err := useCase.CreateSet(tt.args.userId, tt.args.setDto); (err != nil) != tt.wantErr {
				t.Errorf("CreateSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

//func Test_CreateSet_Default(t *testing.T) {
//
//}

func TestSetUseCase_EditSpellList(t *testing.T) {
	tests := getEditSpellListTests()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockSets := mock_boundaries.NewMockSetsRepository(ctrl)
			mockSets.EXPECT().
				UpdateSpellList(gomock.Any(), gomock.Any()).
				AnyTimes()
			mockSets.EXPECT().
				GetById(gomock.Any()).
				AnyTimes().
				Return(tt.getGetByIdReturns)
			repository := &boundaries.Repository{
				Sets: mockSets,
			}
			useCase := usecases.NewSetUseCase(repository)
			if err := useCase.EditSpellList(tt.args.userId, tt.args.setId, tt.args.spells); (err != nil) != tt.wantErr {
				t.Errorf("EditSpellList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type createSetArgs struct {
	userId dto.UserId
	setDto dto.CreateSetDto
}

type createSetTestData struct {
	name                 string
	args                 createSetArgs
	wantErr              bool
	getSetsByNameReturns []dto.SetDto
}

func getCreateSetTests() []createSetTestData {
	return []createSetTestData{
		// TODO: Add test cases.
		{
			name: "default",
			args: createSetArgs{
				userId: dto.UserId{},
				setDto: dto.CreateSetDto{
					Name: "123123",
				},
			},
		},
		{
			name: "empty name",
			args: createSetArgs{
				userId: dto.UserId{},
				setDto: dto.CreateSetDto{
					Name: "",
				},
			},
			wantErr: true,
		},
		{
			name: "already exist name",
			args: createSetArgs{
				userId: dto.UserId{},
				setDto: dto.CreateSetDto{
					Name: "name-name",
				},
			},
			getSetsByNameReturns: []dto.SetDto{
				{
					Id:          dto.SetId{},
					Name:        "name-name",
					UserId:      dto.UserId{},
					Description: "",
					Sources:     nil,
				},
			},
			wantErr: true,
		},
	}
}

type editSpellListArgs struct {
	userId dto.UserId
	setId  dto.SetId
	spells dto.UpdateSetSpellListDto
}

type editSpellListTestData struct {
	name              string
	args              editSpellListArgs
	wantErr           bool
	getGetByIdReturns dto.SetDto
}

func getEditSpellListTests() []editSpellListTestData {
	id1, _ := uuid.Parse("3aa1287a-68ee-48e4-beb0-4b000989ed32")
	id2, _ := uuid.Parse("4867eed8-63ca-463d-8ed3-e7f78f323604")
	id3, _ := uuid.Parse("3fc0f213-f0ae-4703-abc9-8c4784e8988f")
	userId := dto.UserId(id1)
	otherUserId := dto.UserId(id2)
	setId := dto.SetId(id3)
	return []editSpellListTestData{
		// TODO: Add test cases.
		{
			name: "default",
			args: editSpellListArgs{
				userId: userId,
				setId:  setId,
				spells: dto.UpdateSetSpellListDto{},
			},
			getGetByIdReturns: dto.SetDto{
				Id:     setId,
				UserId: userId,
			},
		},
		{
			name: "other user",
			args: editSpellListArgs{
				userId: userId,
				setId:  setId,
				spells: dto.UpdateSetSpellListDto{},
			},
			getGetByIdReturns: dto.SetDto{
				Id:     setId,
				UserId: otherUserId,
			},
			wantErr: true,
		},
	}
}
