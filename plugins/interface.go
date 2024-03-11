package main

type PluginInterface interface {
	Init() error
	Execute() error
}
