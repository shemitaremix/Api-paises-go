package main

import (
	"api-go/server"
	"errors"
	"fmt"
	"io"
	"sync"
)

type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Valor demasiado pequeño")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

var ErrPoolClosed = errors.New("Pool cerrado.")

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		fmt.Println("Adquirir: Recurso existente")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		fmt.Println("Adquirir: Creando nuevo recurso")
		return p.factory()
	}
}

func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		r.Close()
		return
	}
	select {
	case p.resources <- r:
		fmt.Println("Lanzamiento: En cola")
	default:
		fmt.Println("Lanzamiento: Finalizado")
		r.Close()
	}
}

func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		return
	}
	p.closed = true
	close(p.resources)
	for r := range p.resources {
		r.Close()
	}
}

func statusServe(status string, sta chan string) {
	fmt.Println("status: ", status)
	sta <- "el estatus del servidor es: "
}

func main() {

	canal := make(chan string)
	go statusServe("anda cachondisimo ", canal)
	go statusServe("ahi va el server ", canal)
	go statusServe("anda mas prendido que un horno", canal)
	statusMensaje := <-canal
	fmt.Println(statusMensaje)

	srv := server.New(":8080")

	err := srv.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
