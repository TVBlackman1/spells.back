package domain

import (
	"github.com/golang/mock/gomock"
	"spells.tvblackman1.ru/pkg/domain/boundaries"
	"spells.tvblackman1.ru/pkg/domain/dto"
	"spells.tvblackman1.ru/pkg/domain/usecases"
	mock_boundaries "spells.tvblackman1.ru/tests/mocks/pkg/domain/boundaries"
	"testing"
)

func TestUserUseCase_Register(t *testing.T) {
	tests := getRegistrationTests()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockUsers := mock_boundaries.NewMockUsersRepository(ctrl)
			mockUsers.EXPECT().GetUsers(gomock.Any()).AnyTimes()
			mockUsers.EXPECT().CreateUser(gomock.Any()).AnyTimes()
			repository := &boundaries.Repository{
				Users: mockUsers,
			}
			useCase := usecases.NewUserUseCase(repository)
			if _, err := useCase.Register(tt.args.innerDto); (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type registrationArgs struct {
	innerDto dto.UserCreateDto
}

type RegistrationTestData struct {
	name    string
	args    registrationArgs
	wantErr bool
}

func getRegistrationTests() []RegistrationTestData {
	return []RegistrationTestData{
		// TODO: Add test cases.
		{
			name: "default",
			args: registrationArgs{innerDto: dto.UserCreateDto{
				Login:    "123",
				Password: "123",
			}},
		},
		{
			name: "no password error",
			args: registrationArgs{innerDto: dto.UserCreateDto{
				Login:    "123",
				Password: "",
			}},
			wantErr: true,
		},
		{
			name: "admin login",
			args: registrationArgs{innerDto: dto.UserCreateDto{
				Login:    "tvblackman1",
				Password: "123",
			}},
		},
		{
			name: "admin like login error",
			args: registrationArgs{innerDto: dto.UserCreateDto{
				Login:    "tvblackman",
				Password: "123",
			}},
			wantErr: true,
		},
	}
}
