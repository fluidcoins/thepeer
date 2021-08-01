// +build integration

package thepeer

import (
	"os"
	"testing"

	"github.com/Pallinder/go-randomdata"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func getAPIKey(t *testing.T) string {
	t.Helper()

	k := os.Getenv("THEPEER_SECRET_KEY")

	if IsStringEmpty(k) {
		t.Fatal("Please provide a valid API key in the env variable ( THEPEER_SECRET_KEY )")
	}

	return k
}

func getClient(t *testing.T) *Client {
	t.Helper()

	c, err := New(WithSecretKey(getAPIKey(t)))
	require.NoError(t, err)

	return c
}

// consists of 3 tests really, kind of bad but it's our only option
// we index, update and deindex a user.
// cannot afford leaving silly data with ThePeer
// Also there are no mocks here, this SDK is pretty small, so I am not sure if
// it is worth it running test servers here and there to run

type errorMsg struct {
	Message string `json:"message"`
	Errors  struct {
		Identifier []string `json:"identifier"`
	} `json:"errors"`
}

func TestClient_User(t *testing.T) {
	c := getClient(t)

	opts := &IndexUserOptions{
		Name:       randomdata.FullName(randomdata.RandomGender),
		Email:      randomdata.Email(),
		Identifier: randomdata.StringNumber(10, ""),
	}

	idxUser, err := c.IndexUser(opts)
	require.NoError(t, err)

	require.Equal(t, opts.Name, idxUser.Name)
	require.Equal(t, opts.Email, idxUser.Email)
	require.Equal(t, opts.Identifier, idxUser.Identifier)

	uuid.MustParse(idxUser.Reference)

	newIdxUser, err := c.UpdateUser(&UpdateUserOptions{
		Identifier: idxUser.Identifier,
		Reference:  idxUser.Reference,
	})
	require.Error(t, err) // same identifier

	newIdxUser, err = c.UpdateUser(&UpdateUserOptions{
		Identifier: randomdata.StringNumber(10, ""),
		Reference:  idxUser.Reference,
	})
	require.NoError(t, err)
	require.NotEqual(t, idxUser.Identifier, newIdxUser.Identifier)

	require.NoError(t, c.DeleteUser(&DeIndexUserOptions{
		UserReference: idxUser.Reference,
	}))
}
