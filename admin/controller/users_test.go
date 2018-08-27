package controller

import (
	"context"
	"fmt"
	"testing"

	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"

	"github.com/goadesign/goa"
	"github.com/stretchr/testify/assert"

	"github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/admin/app"
	"github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/admin/app/test"
)

func TestCreateUser(t *testing.T) {
	opt := &aetest.Options{StronglyConsistentDatastore: true}
	inst, err := aetest.NewInstance(opt)
	assert.NoError(t, err)
	defer inst.Close()

	req, err := inst.NewRequest("GET", "/", nil)
	assert.NoError(t, err)
	ctx := appengine.NewContext(req)
	ctx = context.WithValue(ctx, "NewRequestFunc", app.NewRequestFunc(inst.NewRequest))

	fmt.Printf("app/controller/users_test.go App Engine context: %v\n", ctx.Value("App Engine context"))

	service := goa.New("example-test")
	ctrl     := NewUsersController(service)

	payload := &app.UserPayload{
		ID: "dummy-id-111",
		Email: "test-user1@example.com",
	}

	_, user := test.CreateUsersCreated(t, ctx, service, ctrl, payload)
	if user.CreatedAt.IsZero() {
		t.Errorf("Invalid CreatedAt, expected NOT Zero, got Zero. user: %v\n", user)
	}
	if user.UpdatedAt.IsZero() {
		t.Errorf("Invalid UpdatedAt, expected NOT Zero, got Zero. user: %v\n", user)
	}
}
