package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/addetz/gophertale/tdd/data"
	"github.com/addetz/gophertale/tdd/handlers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListEmployeeByID(t *testing.T) {
	if os.Getenv("LONG") == "" {
		t.Skip("Skipping TestListEmployeeByID in short mode.")
	}
	// Arrange
	e := data.Employee{
		ID:       "A-1",
		Name:     "Nikita",
		JobTitle: "Boss of everything",
	}
	es := data.NewEmployeeService()
	id, err := es.Add(e)
	require.Nil(t, err)
	ha := handlers.NewHandler(es)
	svr := httptest.NewServer(http.HandlerFunc(ha.ListEmployeeByID))
	defer svr.Close()

	// Act
	r, err := http.Get(fmt.Sprintf("%s/employees?id=%s", svr.URL, *id))

	// Assert
	require.Nil(t, err)
	assert.Equal(t, http.StatusOK, r.StatusCode)

	body, err := io.ReadAll(r.Body)
	r.Body.Close()
	require.Nil(t, err)

	var resp data.Employee
	err = json.Unmarshal(body, &resp)
	require.Nil(t, err)

	assert.Equal(t, e, resp)
}

func TestEmployeeUpsertIntegration(t *testing.T) {
	if os.Getenv("LONG") == "" {
		t.Skip("Skipping TestEmployeeUpsertIntegration in short mode.")
	}
	// Arrange
	e := data.Employee{
		ID:       "A-1",
		Name:     "Nikita",
		JobTitle: "Boss of everything",
	}
	employeePayload, err := json.Marshal(e)
	require.Nil(t, err)
	es := data.NewEmployeeService()
	ha := handlers.NewHandler(es)
	svr := httptest.NewServer(http.HandlerFunc(ha.UpsertEmployee))
	defer svr.Close()

	// Act
	r, err := http.Post(fmt.Sprintf("%s/employees", svr.URL), "application/json", bytes.NewBuffer(employeePayload))

	// Assert
	require.Nil(t, err)
	assert.Equal(t, http.StatusOK, r.StatusCode)
	body, err := io.ReadAll(r.Body)
	r.Body.Close()
	require.Nil(t, err)

	var resp data.Employee
	err = json.Unmarshal(body, &resp)
	require.Nil(t, err)
	assert.Equal(t, e, resp)
}
