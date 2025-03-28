package app

type Nodo struct {
	Producto  Producto
	Izquierdo *Nodo
	Derecho   *Nodo
}

type ArbolBinario struct {
	Raiz *Nodo
}

// Insertar un producto en el árbol
func (a *ArbolBinario) Insertar(producto Producto) {
	a.Raiz = insertarNodo(a.Raiz, producto)
}

func insertarNodo(nodo *Nodo, producto Producto) *Nodo {
	if nodo == nil {
		return &Nodo{Producto: producto}
	}
	if producto.Codigo < nodo.Producto.Codigo {
		nodo.Izquierdo = insertarNodo(nodo.Izquierdo, producto)
	} else {
		nodo.Derecho = insertarNodo(nodo.Derecho, producto)
	}
	return nodo
}

// Eliminar un producto del árbol por Código
func (a *ArbolBinario) Eliminar(codigo string) {
	a.Raiz = eliminarNodo(a.Raiz, codigo)
}

func eliminarNodo(nodo *Nodo, codigo string) *Nodo {
	if nodo == nil {
		return nil
	}
	if codigo < nodo.Producto.Codigo {
		nodo.Izquierdo = eliminarNodo(nodo.Izquierdo, codigo)
	} else if codigo > nodo.Producto.Codigo {
		nodo.Derecho = eliminarNodo(nodo.Derecho, codigo)
	} else {
		if nodo.Izquierdo == nil {
			return nodo.Derecho
		} else if nodo.Derecho == nil {
			return nodo.Izquierdo
		}
		// Encontrar el sucesor (mínimo en el subárbol derecho)
		sucesor := encontrarMin(nodo.Derecho)
		nodo.Producto = sucesor.Producto
		nodo.Derecho = eliminarNodo(nodo.Derecho, sucesor.Producto.Codigo)
	}
	return nodo
}

// Encontrar el nodo con el valor mínimo
func encontrarMin(nodo *Nodo) *Nodo {
	actual := nodo
	for actual.Izquierdo != nil {
		actual = actual.Izquierdo
	}
	return actual
}


// ObtenerTodos devuelve una lista con todos los productos en el árbol en orden
func (a *ArbolBinario) ObtenerTodos() []Producto {
	var productos []Producto
	recorrerInOrden(a.Raiz, &productos)
	return productos
}

// Recorrido inorden para obtener los productos ordenados
func recorrerInOrden(nodo *Nodo, productos *[]Producto) {
	if nodo != nil {
		recorrerInOrden(nodo.Izquierdo, productos)
		*productos = append(*productos, nodo.Producto)
		recorrerInOrden(nodo.Derecho, productos)
	}
}
