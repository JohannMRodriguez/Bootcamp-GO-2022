package main

func main() {

	var apellido string = "Gomez" //bien, define apellio como string y lo decalara como string
	var edad int = "35"           //mal, define edad como entero y lo declara como string, correccion: var edad int = 35
	boolean := "false"            //mal, define boolean como booleano y lo declara como string, correccion: boolean := false
	var sueldo string = 45857.90  //mal, define sueldo como string y lo declara como float, correccion: var sueldo float32 = 45857.90
	var nombre string = "Julian"  //bien, define nombre como string y lo declara como string
}
