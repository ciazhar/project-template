package rest

//func NewActual() model.File {
//	var file model.File
//	testdata.ToStruct("file/actual.1.golden", &file)
//	return file
//}
//
//func NewActual2() model.File {
//	var file model.File
//	testdata.ToStruct("file/actual.2.golden", &file)
//	return file
//}
//
//func NewActualReader(payload model.File) *strings.Reader {
//	return testdata.ToReader(payload)
//}
//
//func TestFileController_Fetch(t *testing.T) {
//	testCases := []struct {
//		name                   string
//		offset                 int
//		limit                  int
//		returnFile []model.File
//		returnError            error
//		httpStatus             int
//	}{
//		{"default", 0, 10, []model.File{NewActual(), NewActual2()}, nil, http.StatusOK},
//		{"error", 1, 5, []model.File{NewActual(), NewActual2()}, errors.New(""), http.StatusInternalServerError},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//			uc := new(mocks.FileUseCase)
//			ctrl := NewFileController(uc)
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
//			uc.On("Fetch", param).Return(testCase.returnFile, testCase.returnError)
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
//func TestFileController_GetByID(t *testing.T) {
//	testCases := []struct {
//		name                   string
//		id                     string
//		returnFile model.File
//		returnError            error
//		httpStatus             int
//	}{
//		{"default", "1", NewActual(), nil, http.StatusOK},
//		//{"bad-request","bukan-id",model.File{}, errors.New(""), http.StatusBadRequest},
//		{"internal-server-error", "10", model.File{}, errors.New(""), http.StatusInternalServerError},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//			uc := new(mocks.FileUseCase)
//			ctrl := NewFileController(uc)
//			w := httptest.NewRecorder()
//
//			router := gin.New()
//			router.GET("/file/:id", ctrl.GetByID)
//
//			r, err := http.NewRequest(http.MethodGet, "/file/"+testCase.id, nil)
//			assert.NoError(t, err)
//			uc.On("GetByID", testCase.id).Return(testCase.returnFile, testCase.returnError)
//
//			router.ServeHTTP(w, r)
//
//			assert.Equal(t, testCase.httpStatus, w.Code)
//			uc.AssertExpectations(t)
//		})
//	}
//}
//
//func TestFileController_Store(t *testing.T) {
//	testCases := []struct {
//		name        string
//		payload     model.File
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
//			uc := new(mocks.FileUseCase)
//			ctrl := NewFileController(uc)
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
//func TestFileController_Update(t *testing.T) {
//	testCases := []struct {
//		name        string
//		payload     model.File
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
//			uc := new(mocks.FileUseCase)
//			ctrl := NewFileController(uc)
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
//func TestFileController_Delete(t *testing.T) {
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
//			uc := new(mocks.FileUseCase)
//			ctrl := NewFileController(uc)
//			w := httptest.NewRecorder()
//
//			router := gin.New()
//			router.DELETE("/file/:id", ctrl.Delete)
//
//			r, err := http.NewRequest(http.MethodDelete, "/file/"+testCase.id, nil)
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
