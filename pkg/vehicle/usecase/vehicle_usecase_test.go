package usecase

//func NewActual() model.Vehicle {
//	var vehicle model.Vehicle
//	testdata.ToStruct("vehicle/actual.1.golden", &vehicle)
//	return vehicle
//}
//
//func NewActual2() model.Vehicle {
//	var vehicle model.Vehicle
//	testdata.ToStruct("vehicle/actual.2.golden", &vehicle)
//	return vehicle
//}
//
//func TestVehicleUseCase_Store(t *testing.T) {
//	repo := new(mocks.VehiclePostgresRepository)
//	uc := NewVehicleUseCase(repo)
//	testCases := []struct {
//		name             string
//		vehicle model.Vehicle
//		returnError      error
//	}{
//		{"default", NewActual(), nil},
//		{"default2", NewActual2(), nil},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//			repo.On("Store", &testCase.vehicle).Return(testCase.returnError)
//
//			err := uc.Store(&testCase.vehicle)
//
//			assert.NoError(t, err)
//			repo.AssertExpectations(t)
//		})
//	}
//}
//
//func TestVehicleUseCase_Fetch(t *testing.T) {
//	repo := new(mocks.VehiclePostgresRepository)
//	uc := NewVehicleUseCase(repo)
//	testCases := []struct {
//		name                   string
//		offset                 int
//		limit                  int
//		returnVehicle []model.Vehicle
//		returnError            error
//	}{
//		{"default", 0, 10, []model.Vehicle{NewActual(), NewActual2()}, nil},
//		{"default2", 0, 5, []model.Vehicle{NewActual(), NewActual2()}, nil},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//
//			param := rest.NewParam()
//			param.Offset = 1
//			param.Limit = 10
//
//			repo.On("Fetch", param).Return(testCase.returnVehicle, testCase.returnError)
//
//			expected, err := uc.Fetch(param)
//
//			assert.NotEmpty(t, expected)
//			assert.NoError(t, err)
//			assert.Len(t, expected, len(testCase.returnVehicle))
//			repo.AssertExpectations(t)
//		})
//	}
//}
//
//func TestVehicleUseCase_GetByID(t *testing.T) {
//	repo := new(mocks.VehiclePostgresRepository)
//	uc := NewVehicleUseCase(repo)
//	testCases := []struct {
//		name                   string
//		id                     string
//		returnVehicle model.Vehicle
//		returnError            error
//	}{
//		{"default", "1", NewActual(), nil},
//		{"default2", "2", NewActual(), nil},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//			repo.On("GetByID", testCase.id).Return(testCase.returnVehicle, testCase.returnError)
//
//			expected, err := uc.GetByID(testCase.id)
//
//			assert.NoError(t, err)
//			assert.NotNil(t, expected)
//			repo.AssertExpectations(t)
//		})
//	}
//}
//
//func TestVehicleUseCase_Update(t *testing.T) {
//	repo := new(mocks.VehiclePostgresRepository)
//	uc := NewVehicleUseCase(repo)
//	actual := NewActual()
//	//actual.Id = "100"
//	testCases := []struct {
//		name             string
//		vehicle model.Vehicle
//		returnError      error
//	}{
//		{"default", NewActual(), nil},
//		{"error", actual, errors.New("not found")},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//			repo.On("GetByID", testCase.vehicle.Id).Return(testCase.vehicle, testCase.returnError)
//			repo.On("Update", &testCase.vehicle).Return(testCase.returnError)
//
//			err := uc.Update(&testCase.vehicle)
//
//			assert.Equal(t, err, testCase.returnError)
//			repo.AssertExpectations(t)
//		})
//	}
//}
//
//func TestVehicleUseCase_Delete(t *testing.T) {
//	repo := new(mocks.VehiclePostgresRepository)
//	uc := NewVehicleUseCase(repo)
//	actual := NewActual()
//	//actual.Id = "100"
//	testCases := []struct {
//		name             string
//		vehicle model.Vehicle
//		returnError      error
//	}{
//		{"default", NewActual(), nil},
//		{"error", actual, errors.New("not found")},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//
//			wayback := time.Date(1974, time.May, 19, 1, 2, 3, 4, time.UTC)
//			patch := monkey.Patch(time.Now, func() time.Time { return wayback })
//			defer patch.Unpatch()
//
//			repo.On("GetByID", testCase.vehicle.Id).Return(testCase.vehicle, testCase.returnError)
//
//			testCase.vehicle.DeletedAt = time.Now()
//			repo.On("Update", &testCase.vehicle).Return(testCase.returnError)
//
//			err := uc.Delete(testCase.vehicle.Id.String())
//
//			assert.Equal(t, err, testCase.returnError)
//			//repo.AssertExpectations(t)
//		})
//	}
//}
