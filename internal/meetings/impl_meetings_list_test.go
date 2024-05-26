package meetings

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/revilO602/meetings-webapi/internal/db_service"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MeetingsSuite struct {
	suite.Suite
	dbServiceMock *DbServiceMock[MeetingsListEntry]
}

func TestMeetingsSuite(t *testing.T) {
	suite.Run(t, new(MeetingsSuite))
}

type DbServiceMock[DocType interface{}] struct {
	mock.Mock
}

func (this *DbServiceMock[DocType]) CreateDocument(ctx context.Context, id string, document *DocType) error {
	args := this.Called(ctx, id, document)
	return args.Error(0)
}

func (this *DbServiceMock[DocType]) FindDocument(ctx context.Context, id string) (*DocType, error) {
	args := this.Called(ctx, id)
	return args.Get(0).(*DocType), args.Error(1)
}

func (this *DbServiceMock[DocType]) UpdateDocument(ctx context.Context, id string, document *DocType) error {
	args := this.Called(ctx, id, document)
	return args.Error(0)
}

func (this *DbServiceMock[DocType]) DeleteDocument(ctx context.Context, id string) error {
	args := this.Called(ctx, id)
	return args.Error(0)
}

func (this *DbServiceMock[DocType]) Disconnect(ctx context.Context) error {
	args := this.Called(ctx)
	return args.Error(0)
}

func (this *DbServiceMock[DocType]) GetAllDocuments(ctx context.Context) ([]DocType, error) {
	args := this.Called(ctx)
	return args.Get(0).([]DocType), args.Error(0)
}

func (suite *MeetingsSuite) SetupTest() {
	suite.dbServiceMock = &DbServiceMock[MeetingsListEntry]{}

	var _ db_service.DbService[MeetingsListEntry] = suite.dbServiceMock

	suite.dbServiceMock.
		On("FindDocument", mock.Anything, mock.Anything).
		Return(
			&MeetingsListEntry{
				Id:          "test-entry",
				DoctorName:  "test-doctor",
				PatientName: "test-patient",
				Date:        "2038-02-02",
				StartTime:   "03:45",
				EndTime:     "04:45",
				Important:   false,
				Platform:    "ms_teams",
				Symptoms:    "deded",
				Diagnosis:   "deededed",
				Notes:       "ddedee",
			},
			nil,
		)
}

func (suite *MeetingsSuite) Test_UpdateMeetings_DbServiceUpdateCalled() {
	// ARRANGE
	suite.dbServiceMock.
		On("UpdateDocument", mock.Anything, mock.Anything, mock.Anything).
		Return(nil)
	json := `{
		"id": "test-entry",
		"doctorName": "Dr. Jozef Mrkva",
		"patientName": "Juraj Prvý",
		"date": "2038-02-02",
		"startTime": "03:45",
		"endTime": "04:45",
		"important": false,
		"platform": "ms_teams",
		"symptoms": "deded",
		"diagnosis": "deededed",
		"notes": "ddedee"
	  }`

	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Set("db_service", suite.dbServiceMock)
	ctx.Params = []gin.Param{
		{Key: "entryId", Value: "test-entry"},
	}
	ctx.Request = httptest.NewRequest("PUT", "/meetings/test-entry", strings.NewReader(json))
	sut := implMeetingsListAPI{}
	// ACT
	sut.UpdateMeeting(ctx)

	// ASSERT
	suite.dbServiceMock.AssertCalled(suite.T(), "UpdateDocument", mock.Anything, "test-entry", mock.Anything)
}
