package csrfRepository

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	redisServer *miniredis.Miniredis
	session     csrfRepository
	tokenValue  string
}

func (s *Suite) SetupSuite() {
	var err error
	s.redisServer, err = miniredis.Run()
	require.NoError(s.T(), err)

	addr := s.redisServer.Addr()
	redisConn := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	s.tokenValue = "1"

	s.session = NewCsrfRepository(redisConn)
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

func (s *Suite) TestAdd() {
	token := "testToken"

	err := s.session.Add(token)
	require.NoError(s.T(), err)

	value, err := s.redisServer.Get(token)
	require.NoError(s.T(), err)
	require.Equal(s.T(), value, s.tokenValue)

	s.redisServer.FastForward(time.Second * 4000)

	_, err = s.redisServer.Get(token)
	require.Equal(s.T(), err, nil)

	s.redisServer.Close()

	err = s.session.Add(token)
	require.Error(s.T(), err)
}

func (s *Suite) TestGetLoginBySessionID() {
	token := "testToken"
	require.NoError(s.T(), s.redisServer.Set(token, s.tokenValue))

	err := s.session.Check(token)
	require.Error(s.T(), err)

	newToken := "wewxcvqsd"

	err = s.session.Check(newToken)
	require.NoError(s.T(), err)

	s.redisServer.Close()

	err = s.session.Check(token)
	require.Error(s.T(), err)
}
