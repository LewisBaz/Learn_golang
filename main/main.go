package main

import (
	"fmt"
	"strconv"
	"strings"
)

var a int = 5

func main() {
	fmt.Print(a)
	fmt.Print("hello \n")
	fmt.Print(intToString())
	fmt.Print(a)
	bs := false
	fmt.Print(bs)
	fmt.Print("\n")
	fmt.Println(greetings("Ivan"))
	fmt.Println(DomainForLocale("sss.com", "ru"))
	fmt.Println(DomainForLocale("sss.com", ""))
	fmt.Println(Validate(UserCreateRequest{"John", 15}))
	fmt.Println(Validate(UserCreateRequest{"John f", 0}))
	Start()
	StringsStart()
	FuncsStart()
	StartErrors()
}

func intToString() string {
	s := strconv.Itoa(-43)
	return s
}

func greetings(name string) string {
	return "Hello, " + name
}

func DomainForLocale(domain, locale string) string {
	if locale == "" {
		return "en." + domain
	} else {
		return locale + "." + domain
	}
}

func ModifySpaces(s, mode string) string {
	switch mode {
	case "dash":
		return strings.ReplaceAll(s, " ", "-")
	case "underscore":
		return strings.ReplaceAll(s, " ", "_")
	default:
		return strings.ReplaceAll(s, " ", "*")
	}
}

type UserCreateRequest struct {
	FirstName string
	Age       int
}

func Validate(req UserCreateRequest) string {

	isNameValid := !strings.Contains(req.FirstName, " ") && req.FirstName != ""
	isAgeValid := req.Age != 0 && req.Age <= 150 && req.Age > 1

	if isNameValid && isAgeValid {
		return ""
	} else {
		return "invalid request"
	}
}

const (
	ok        = 0
	cancelled = 1
	unknown   = 2
)

func ErrorMessageToCode(msg string) int {
	switch msg {
	case "OK":
		return ok
	case "CANCELLED":
		return cancelled
	default:
		return unknown
	}
}

type IVoiceable interface {
	Voice() string
}

type Cat struct {
	//
}

type Cow struct {
	// …
}

type Dog struct {
	// …
}

func (cat Cat) Voice() string {
	return "Мяу"
}

func (cow Cow) Voice() string {
	return "Мууу"
}

func (dog Dog) Voice() string {
	return "Гав"
}
