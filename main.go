// package main implements the main entrypoint

package main

func main() {
	a := App{}
	a.Initialize()

	a.Run(":8010")
}
