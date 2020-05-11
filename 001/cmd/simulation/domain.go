package main

import (
	"fmt"
	"strconv"
)

type Document struct {
	Type   string
	Number int
}

type Registration struct {
	Name string
	Age  *int
	Docs []Document
}

func (r *Registration) String() string {
	age := "?"
	if r.Age != nil {
		age = strconv.Itoa(*r.Age)
	}
	return fmt.Sprintf("Registration{Name:%v, Age:%s, Docs:%v}", r.Name, age, r.Docs)
}

type RequestType string

const (
	CustomerInvited   RequestType = "CUSTOMER-INVITED"
	AgeIsMissing      RequestType = "AGE-IS-MISSING"
	TooYoung          RequestType = "TOO-YOUNG"
	CustomerResponded RequestType = "CUSTOMER-RESPONDED"
	DocsRequired      RequestType = "DOCS-REQUIRED"
	ReviewDocs        RequestType = "REVIEW-DOCS"
	InvalidDoc        RequestType = "INVALID-DOC"
	ValidDoc          RequestType = "VALID-DOC"
)

type Request struct {
	Type         RequestType
	Registration *Registration
}

func (r *Request) WithType(t RequestType) *Request {
	return &Request{
		Type:         t,
		Registration: r.Registration,
	}
}

func (r *Request) String() string {
	return fmt.Sprintf("Request{Type:%v, Data:%v}", r.Type, r.Registration)
}
