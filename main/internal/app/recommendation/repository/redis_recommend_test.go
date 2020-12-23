package recommendRepository

import (
	"testing"

	"github.com/go-park-mail-ru/2020_2_JMickhs/main/configs"
	"github.com/lib/pq"
	"github.com/spf13/viper"

	recommModels "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/recommendation/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	redisServer *miniredis.Miniredis
	recommend   PostgreRecommendationRepository
	mock        sqlmock.Sqlmock
}

func (s *Suite) SetupSuite() {
	var err error
	s.redisServer, err = miniredis.Run()
	require.NoError(s.T(), err)

	addr := s.redisServer.Addr()
	redisConn := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	sqlxDb := sqlx.NewDb(db, "sqlmock")
	s.recommend = NewPostgreRecommendationRepository(sqlxDb, redisConn)
	s.mock = mock
}

func (s *Suite) AfterTest(_, _ string) {
	s.SetupSuite()
}

func (s *Suite) TearDownSuite() {
	s.redisServer.Close()
}

func TestSessions(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestAddMessageInChat() {
	err := s.recommend.AddInSearchHistory(2, "moscow")
	require.NoError(s.T(), err)
	value, err := s.redisServer.List("2")

	require.NoError(s.T(), err)
	require.Equal(s.T(), value, []string{"moscow"})

	s.redisServer.Close()
}

func (s *Suite) TestGetChatHistory() {
	hotelsTest := []recommModels.HotelRecommend{
		{HotelID: 3, Name: "moscow", Image: "sfd", Location: "piter", Rating: "3"},
	}
	err := s.recommend.AddInSearchHistory(2, "moscow")
	require.NoError(s.T(), err)
	value, err := s.redisServer.List("2")

	require.NoError(s.T(), err)
	require.Equal(s.T(), value, []string{"moscow"})

	rowChat := sqlmock.NewRows([]string{"hotel_id", "name", "concat", "location", "curr_rating"}).AddRow(
		3, "moscow", "sfd", "piter", "3")

	s.mock.ExpectQuery(GetRecommendationFromSearchHistory).
		WithArgs(viper.GetString(configs.ConfigFields.S3Url),
			viper.GetInt(configs.ConfigFields.RecommendationCount), "moscow", pq.Array([]int64{4, 5})).
		WillReturnRows(rowChat)

	hotels, err := s.recommend.GetHotelsFromHistory(2, []int64{4, 5})

	require.Equal(s.T(), hotelsTest, hotels)
	s.redisServer.Close()
}
