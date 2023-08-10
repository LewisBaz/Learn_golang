package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

func StartErrors() {
	fmt.Println(DecodeAndValidateRequest([]byte("{\"email\":\"test\",\"password\":\"test\",\"password_confirmation\":\"test2\"}")))
	fmt.Println(GetErrorMsg(errors.New("random error")))
	fmt.Println(ExecuteMergeDictsJob(&MergeDictsJob{Dicts: []map[string]string{
		{"a": "b", "b": "c"},
		{"d": "e", "f": "g"},
		{"a": "z", "f": "g"},
	}}))
	fmt.Println(ExecuteMergeDictsJob(&MergeDictsJob{Dicts: []map[string]string{
		{"a": "b"},
		nil,
	}}))
}

type CreateUserRequest struct {
    Email                string `json:"email"`
    Password             string `json:"password"`
    PasswordConfirmation string `json:"password_confirmation"`
}

var (
	errEmailRequired                = errors.New("email is required")
	errPasswordRequired             = errors.New("password is required")
	errPasswordConfirmationRequired = errors.New("password confirmation is required")
	errPasswordDoesNotMatch         = errors.New("password does not match with the confirmation")
)

func DecodeAndValidateRequest(requestBody []byte) (CreateUserRequest, error) {
	usr := CreateUserRequest{}
	err := json.Unmarshal(requestBody, &usr)
	if len(usr.Email) == 0 {
		return CreateUserRequest{}, errEmailRequired
	}
	if len(usr.Password) == 0 {
		return CreateUserRequest{}, errPasswordRequired
	}
	if len(usr.PasswordConfirmation) == 0 {
		return CreateUserRequest{}, errPasswordConfirmationRequired
	}
	if usr.Password != usr.PasswordConfirmation {
		return CreateUserRequest{}, errPasswordDoesNotMatch
	}

	return usr, err
}

type nonCriticalError struct{}

func (e nonCriticalError) Error() string {
    return "validation error"
}

var (
    errBadConnection = errors.New("bad connection")
    errBadRequest    = errors.New("bad request")
)

const unknownErrorMsg = "unknown error"

func GetErrorMsg(err error) string {
	if errors.As(err, &nonCriticalError{}) {
		return ""
	} 
	if errors.Is(err, errBadConnection) {
		return errBadConnection.Error()
	}
	if errors.Is(err, errBadRequest) {
		return errBadRequest.Error()
	} else {
		return unknownErrorMsg
	}
}

type MergeDictsJob struct {
	Dicts      []map[string]string
	Merged     map[string]string
	IsFinished bool
}

var (
	errNotEnoughDicts = errors.New("at least 2 dictionaries are required")
	errNilDict        = errors.New("nil dictionary")
)

func (j *MergeDictsJob) setFinished() {
	j.IsFinished = true
}

func ExecuteMergeDictsJob(job *MergeDictsJob) (*MergeDictsJob, error) {
	job.Merged = make(map[string]string)
	containsNil := false

	defer job.setFinished()

	for _, dic := range job.Dicts {
		if dic == nil {
			containsNil = true
		} else {
			for key, element := range dic {
				job.Merged[key] = element
			}
		}
	}

	if len(job.Dicts) < 2 {
		return job, errNotEnoughDicts
	}

	if containsNil {
		return job, errNilDict
	}

	return job, nil
}