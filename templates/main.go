package main

import "fmt"

func main() {
	name := "Eonsu Bae"

	tpl := `
		<!DOCTYPE HTML>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Hello World!</title>
		</head>
		<body>
		<h1>` + name + ` </h1>
		</bdoy>
		</html>
	`

	fmt.Println(tpl)
}