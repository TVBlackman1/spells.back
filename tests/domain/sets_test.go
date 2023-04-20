package domain

import (
	"github.com/golang/mock/gomock"
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
