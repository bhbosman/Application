package main

type IWaitGroup interface {
	Add() error
	Done() error

}

type IMessageServiceItem interface {
	IWaitGroup
	Message() (interface{}, error)
}

