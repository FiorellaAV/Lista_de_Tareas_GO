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

//Funcion para pasara a completa una tarea.
func (t *tarea) pasarACompleta() {
	t.completado = true
}

//Funcion para actualizar la descripcion de una tarea
func (t *tarea) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

//Funcion para actualizar el nombre de una tarea
func (t *tarea) actualizarNombre(nombre string) {
	t.nombre = nombre
}

//Funcion para agregar una tarea a la lista
func (l *listaDeTareas) agregarALista(tarea *tarea) {
	l.lista = append(l.lista, tarea)
}

//Funcion para eliminar una tarea de la lista segun el numero.
func (l *listaDeTareas) eliminarDeLista(numeroDeTarea int) {
	l.lista = append(l.lista[:numeroDeTarea], l.lista[numeroDeTarea+1:]...)
}

//Funcion para imprimir toda la lista
func (l *listaDeTareas) imprimirLista() {
	for _, tarea := range l.lista {
		fmt.Println("Nombre", tarea.nombre)
		fmt.Println("Descripcion", tarea.descripcion)
	}
}

//Funcion que imprime solo las tareas completas
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
