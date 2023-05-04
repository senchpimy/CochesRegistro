package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Venta de Autos")

	hello := widget.NewLabel("Venta de Autos")
	
	marca := widget.NewEntry()
	marca.SetPlaceHolder("Ingrese la Marca")
	modelo := widget.NewEntry()
	modelo.SetPlaceHolder("Ingrese el modelo")
	año := widget.NewEntry()
	año.SetPlaceHolder("Ingrese el año")
	color := widget.NewEntry()
	color.SetPlaceHolder("Ingrese la color")

	ingresoDatos:=container.New(layout.NewVBoxLayout(),
			marca,modelo,color,año)
	comprados:=container.New(layout.NewVBoxLayout())
	for i:=0;i<100;i++{
		zona:=container.New(layout.NewHBoxLayout())
		zona.AddObject(widget.NewLabel("Papu"))
		zona.AddObject(widget.NewButton(fmt.Sprintf("Hola %d", i), createButtonCallback(i)))
		comprados.AddObject(zona)
	}
	compradosScroll:=container.NewVScroll(comprados)
	compradosScroll.SetMinSize(fyne.NewSize(500, 500))
	zonaPrincipal:=container.NewVBox(compradosScroll)

	grid := container.New(layout.NewGridLayout(3), 
	widget.NewButton("Comprar", func() {zonaPrincipal.RemoveAll();zonaPrincipal.AddObject(ingresoDatos)}),
		widget.NewButton("Vender", func() {zonaPrincipal.RemoveAll()}),
		widget.NewButton("Disponibles", func() {zonaPrincipal.RemoveAll();zonaPrincipal.AddObject(compradosScroll)}) )

	content:=container.NewVBox(hello,grid,zonaPrincipal)
	w.SetContent(content)

	w.ShowAndRun()
}

func createButtonCallback(n int) func() {
    return func() {
        fmt.Println(n)
    }
}
