package citadel

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func TestCreateClient(t *testing.T) {
	assert.NotPanics(t, func() {
		NewClient(&ClientConfig{})
	})
}

var mockedClient = client{
	baseUrl:      "https://example.com",
	client:       nil,
	apiKey:       "FOO",
	preSharedKey: "BAR",
}

func TestInviteUser(t *testing.T) {
	mockedClient.client = NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://example.com/users.invite", req.URL.String())

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`{"userId":"foo"}`)),
		}
	})

	invitedUser, err := mockedClient.InviteUser(&InviteUserRequest{
		Username:            "foobar",
		EmailAddress:        "foo@bar.com",
		AllowedAuthFlows:    []string{AuthFlowEmailCode, AuthFlowPassword},
		RedirectUri:         "https://example.com",
		ExpirationInSeconds: 3600,
		Language:            "en",
	})

	assert.Nil(t, err)
	assert.Equal(t, "foo", invitedUser.UserId)

	// TODO - Test Errors
}

func TestGetUser(t *testing.T) {
	mockedClient.client = NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://example.com/users.get", req.URL.String())

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`{"id":"foo"}`)),
		}
	})

	user, err := mockedClient.GetUser(&GetUserRequest{
		UserId: "foo",
	})

	assert.Nil(t, err)
	assert.Equal(t, "foo", user.UserId)

	// TODO - Test Errors
}

func TestDeleteUser(t *testing.T) {
	mockedClient.client = NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://example.com/users.delete", req.URL.String())

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`{"status":"deleted"}`)),
		}
	})

	resp, err := mockedClient.DeleteUser(&DeleteUserRequest{
		UserId: "foo",
	})

	assert.Nil(t, err)
	assert.Equal(t, "deleted", resp.Status)

	// TODO - Test Errors
}

func TestListUsers(t *testing.T) {
	t.Skip()
}

func TestUpdateUser(t *testing.T) {
	t.Skip()
}

func TestGetAllUserMetadata(t *testing.T) {
	t.Skip()
}

func TestSetUserMetadata(t *testing.T) {
	t.Skip()
}

func TestDeleteUserMetadata(t *testing.T) {
	t.Skip()
}

func TestAdminResetPassword(t *testing.T) {
	t.Skip()
}

func TestVerifyJwt(t *testing.T) {
	t.Skip()
}
