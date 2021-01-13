package main

import "github.com/atomgenie/prettier-init/configurators"

func main() {

	endChans := make(chan int)

	go (func() {
		configurators.InstallPrettier()
		endChans <- 0
	})()

	go (func() {
		configurators.ConfigureVSCode()
		endChans <- 0
	})()

	<-endChans
	<-endChans

}
