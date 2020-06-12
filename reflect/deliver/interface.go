package main

type Environment interface {
	getVm()
	getLog()
	getDB()
	getEnv(string)
}

type Execute interface {
	Create(namespace string, data []byte)
	Call(address string, input []byte) []byte
}
