package data_test

import (
	"errors"
	"testing"

	"github.com/addetz/gophertale/tdd/data"
	"github.com/stretchr/testify/assert"
)

func TestEmployeeService(t *testing.T) {
	t.Run("add new employee", func(t *testing.T) {
		s := data.NewEmployeeService()
		want := data.Employee{
			ID:       "A-1",
			Name:     "Nikita",
			JobTitle: "Boss of everything",
		}
		id, err := s.Add(want)
		assert.Nil(t, err)
		assert.Equal(t, want.ID, *id)
	})
	t.Run("add invalid employee", func(t *testing.T) {
		s := data.NewEmployeeService()
		want := data.Employee{
			ID:       "A-1",
			JobTitle: "Boss of everything",
		}
		id, err := s.Add(want)
		assert.Nil(t, id)
		assert.Equal(t, errors.New("empty name"), err)
	})
	t.Run("get existing employee", func(t *testing.T) {
		s := data.NewEmployeeService()
		want := data.Employee{
			ID:       "A-1",
			Name:     "Nikita",
			JobTitle: "Boss of everything",
		}
		id, err := s.Add(want)
		assert.Nil(t, err)
		got, err := s.Get(*id)
		assert.Nil(t, err)
		assert.Equal(t, want, *got)
	})
	t.Run("get invalid employee", func(t *testing.T) {
		s := data.NewEmployeeService()
		got, err := s.Get("A-3")
		assert.Nil(t, got)
		assert.Equal(t, errors.New("id not found"), err)
	})
}
