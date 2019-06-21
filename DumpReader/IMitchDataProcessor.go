package main

type IMitchDataProcessor interface {
	IMessageService
	//IService
	DeclareInterestInMessages() ([]byte, error)
}

