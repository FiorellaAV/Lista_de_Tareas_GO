package main

import "fmt"

type tarea struct {
	nombre      string
	descripcion string
	completado  bool
}
type listaDeTareas struct {
	lista []*tarea
}

func (t *tarea) pasarACompleta() {
	t.completado = true
}

func (t *tarea) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *tarea) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func (l *listaDeTareas) agregarALista(tarea *tarea) {
	l.lista = append(l.lista, tarea)
}
func (l *listaDeTareas) eliminarDeLista(numeroDeTarea int) {
	l.lista = append(l.lista[:numeroDeTarea], l.lista[numeroDeTarea+1:]...)
}
func (l *listaDeTareas) imprimirLista() {
	for _, tarea := range l.lista {
		fmt.Println("Nombre", tarea.nombre)
		fmt.Println("Descripcion", tarea.descripcion)
	}
}
func (l *listaDeTareas) imprimirListaCompletados() {
	for _, tarea := range l.lista {
		if tarea.completado == true {
			fmt.Println("Nombre", tarea.nombre)
			fmt.Println("Descripcion", tarea.descripcion)
		}
	}
}

func main() {
	t := tarea{
		nombre:      "Limpiar teclado",
		descripcion: "Sacar las teclas del teclado y limpiar con cepillo",
		completado:  false,
	}
	t2 := tarea{
		nombre:      "Limpiar notebook",
		descripcion: "limpiar con cepillo y blem",
		completado:  false,
	}
	//%+v\n es para que nos imprima por pantalla como si fuera un clave-valor
	t.pasarACompleta()
	fmt.Printf("%+v\n", t)

	ls := listaDeTareas{
		lista: []*tarea{
			&t,
		},
	}
	ls.agregarALista(&t2)
	ls.imprimirLista()

}
