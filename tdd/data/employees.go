package data

import "errors"

type EmployeeService struct {
	data map[string]Employee
}

type Employee struct {
	ID, Name, JobTitle string
}

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{
		data: map[string]Employee{},
	}
}

func (s *EmployeeService) Add(e Employee) (*string, error) {
	if e.Name == "" {
		return nil, errors.New("empty name")
	}
	s.data[e.ID] = e
	return &e.ID, nil
}

func (s *EmployeeService) Get(id string) (*Employee, error) {
	e, ok := s.data[id]
	if !ok {
		return nil, errors.New("id not found")
	}
	return &e, nil
}
