package rest

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
//func NewActualReader(payload model.Vehicle) *strings.Reader {
//	return testdata.ToReader(payload)
//}
//
//func TestVehicleController_Fetch(t *testing.T) {
//	testCases := []struct {
//		name                   string
//		offset                 int
//		limit                  int
//		returnVehicle []model.Vehicle
//		returnError            error
//		httpStatus             int
//	}{
//		{"default", 0, 10, []model.Vehicle{NewActual(), NewActual2()}, nil, http.StatusOK},
//		{"error", 1, 5, []model.Vehicle{NewActual(), NewActual2()}, errors.New(""), http.StatusInternalServerError},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//			uc := new(mocks.VehicleUseCase)
//			ctrl := NewVehicleController(uc)
//
//			w := httptest.NewRecorder()
//
//			router := gin.New()
//			router.GET("/organization", ctrl.Fetch)
//
//			r, err := http.NewRequest(http.MethodGet, "/organization?offset="+strconv.Itoa(testCase.offset)+"&limit="+strconv.Itoa(testCase.limit), nil)
//			assert.NoError(t, err)
//			param := rest.NewParam()
//			param.Limit = testCase.limit
//			param.Offset = testCase.offset
//			uc.On("Fetch", param).Return(testCase.returnVehicle, testCase.returnError)
//
//			router.ServeHTTP(w, r)
//
//			assert.Equal(t, testCase.httpStatus, w.Code)
//			uc.AssertExpectations(t)
//		})
//	}
//
//}
//
//func TestVehicleController_GetByID(t *testing.T) {
//	testCases := []struct {
//		name                   string
//		id                     string
//		returnVehicle model.Vehicle
//		returnError            error
//		httpStatus             int
//	}{
//		{"default", "1", NewActual(), nil, http.StatusOK},
//		//{"bad-request","bukan-id",model.Vehicle{}, errors.New(""), http.StatusBadRequest},
//		{"internal-server-error", "10", model.Vehicle{}, errors.New(""), http.StatusInternalServerError},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//			uc := new(mocks.VehicleUseCase)
//			ctrl := NewVehicleController(uc)
//			w := httptest.NewRecorder()
//
//			router := gin.New()
//			router.GET("/vehicle/:id", ctrl.GetByID)
//
//			r, err := http.NewRequest(http.MethodGet, "/vehicle/"+testCase.id, nil)
//			assert.NoError(t, err)
//			uc.On("GetByID", testCase.id).Return(testCase.returnVehicle, testCase.returnError)
//
//			router.ServeHTTP(w, r)
//
//			assert.Equal(t, testCase.httpStatus, w.Code)
//			uc.AssertExpectations(t)
//		})
//	}
//}
//
//func TestVehicleController_Store(t *testing.T) {
//	testCases := []struct {
//		name        string
//		payload     model.Vehicle
//		reader      io.Reader
//		returnError error
//		httpStatus  int
//	}{
//		{"default", NewActual(), NewActualReader(NewActual()), nil, http.StatusOK},
//		{"bad-request", NewActual(), strings.NewReader("{"), errors.New(""), http.StatusBadRequest},
//		{"error", NewActual2(), NewActualReader(NewActual2()), errors.New(""), http.StatusInternalServerError},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//
//			uc := new(mocks.VehicleUseCase)
//			ctrl := NewVehicleController(uc)
//			w := httptest.NewRecorder()
//			router := gin.New()
//			router.POST("/organization-type", ctrl.Store)
//
//			r, err := http.NewRequest(http.MethodPost, "/organization-type", testCase.reader)
//			assert.NoError(t, err)
//			uc.On("Store", &testCase.payload).Return(testCase.returnError)
//
//			router.ServeHTTP(w, r)
//
//			assert.Equal(t, testCase.httpStatus, w.Code)
//		})
//	}
//}
//
//func TestVehicleController_Update(t *testing.T) {
//	testCases := []struct {
//		name        string
//		payload     model.Vehicle
//		reader      io.Reader
//		returnError error
//		httpStatus  int
//	}{
//		{"default", NewActual(), NewActualReader(NewActual()), nil, http.StatusOK},
//		{"bad-request", NewActual(), strings.NewReader("{"), errors.New(""), http.StatusBadRequest},
//		{"error", NewActual2(), NewActualReader(NewActual2()), errors.New(""), http.StatusInternalServerError},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//			uc := new(mocks.VehicleUseCase)
//			ctrl := NewVehicleController(uc)
//			w := httptest.NewRecorder()
//			router := gin.New()
//			router.PUT("/organization-type", ctrl.Update)
//
//			r, err := http.NewRequest(http.MethodPut, "/organization-type", testCase.reader)
//			assert.NoError(t, err)
//			uc.On("Update", &testCase.payload).Return(testCase.returnError)
//
//			router.ServeHTTP(w, r)
//
//			assert.Equal(t, testCase.httpStatus, w.Code)
//		})
//	}
//}
//
//func TestVehicleController_Delete(t *testing.T) {
//	testCases := []struct {
//		name        string
//		id          string
//		returnError error
//		httpStatus  int
//	}{
//		{"default", "1", nil, http.StatusOK},
//		{"error", "10", errors.New(""), http.StatusInternalServerError},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//			uc := new(mocks.VehicleUseCase)
//			ctrl := NewVehicleController(uc)
//			w := httptest.NewRecorder()
//
//			router := gin.New()
//			router.DELETE("/vehicle/:id", ctrl.Delete)
//
//			r, err := http.NewRequest(http.MethodDelete, "/vehicle/"+testCase.id, nil)
//			assert.NoError(t, err)
//			uc.On("Delete", testCase.id).Return(testCase.returnError)
//
//			router.ServeHTTP(w, r)
//
//			assert.Equal(t, testCase.httpStatus, w.Code)
//			uc.AssertExpectations(t)
//		})
//	}
//}
