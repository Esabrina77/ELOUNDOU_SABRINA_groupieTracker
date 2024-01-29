package main

import (
	"fmt"
	"groupieTrack/manager"
	"groupieTrack/roots"
	initTemplate "groupieTrack/templates"
)

func main() {
	manager.PrintColorResult("purple", "server is running...")

	fmt.Println("")
	manager.PrintColorResult("yellow", "CLICK HERE to OPEN  PAGE--->")
	manager.PrintColorResult("blue", " http://localhost:8080/connexion \n")
	fmt.Println("")
	manager.PrintColorResult("green", "TO STOP THE SERVER , PRESS  'ctrl+C' ")
	initTemplate.InitTemplate()
	roots.InitServe()
}
