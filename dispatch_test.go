package main

import (
	"context"
	"testing"
)

func Test_dispatch(t *testing.T) {
	ctx := context.Background()
	projectID := fakeProjectID(t)

	client := testFirestoreClient(t, ctx, projectID)

	pl := &PairingLogic{
		rdb:     &FirestoreRecurserDB{client},
		version: "test string",
	}

	t.Run("version", func(t *testing.T) {
		resp, err := dispatch(ctx, pl, "version", nil, 0, "fake@recurse.example.net", "Your Name")
		if err != nil {
			t.Fatal(err)
		}

		expected := "test string"
		if resp != expected {
			t.Errorf("expected %q, got %q", expected, resp)
		}
	})
}
