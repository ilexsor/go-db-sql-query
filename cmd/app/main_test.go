package main

import (
	"database/sql"
	"testing"

	"github.com/Yandex-Practicum/final-project-encoding-go/internal"
	"github.com/Yandex-Practicum/final-project-encoding-go/internal/database"
	"github.com/stretchr/testify/require"
)

var dbPath = "../../internal/database/demo.db"

func TestInsertUpdateDelete(t *testing.T) {
	// prepare
	db, err := sql.Open("sqlite", dbPath)
	require.NoError(t, err)
	defer db.Close()

	newClient := internal.Client{
		FIO:      "TEST",
		Login:    "TEST",
		Birthday: "TEST",
		Email:    "TEST",
	}

	// insert
	id, err := database.InsertClient(db, &newClient)

	require.NoError(t, err)
	require.NotEmpty(t, id)
	newClient.ID = int(id)

	got, err := database.SelectClient(db, id)
	require.NoError(t, err)
	require.Equal(t, newClient, got)

	// update
	newLogin := "TEST_NEW"
	err = database.UpdateClientLogin(db, newLogin, id)
	require.NoError(t, err)

	got, err = database.SelectClient(db, id)
	require.NoError(t, err)
	require.Equal(t, newLogin, got.Login)

	// delete
	err = database.DeleteClient(db, id)
	require.NoError(t, err)

	got, err = database.SelectClient(db, id)
	require.Equal(t, sql.ErrNoRows, err)
	require.Empty(t, got)
}
