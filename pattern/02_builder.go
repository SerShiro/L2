package pattern

import "fmt"

type Product struct {
	PartA string
	PartB string
	PartC string
}

type Builder interface {
	BuildPartA()
	BuildPartB()
	BuildPartC()
	GetProduct() *Product
}

type MyBuilder struct {
	product *Product
}

func (m MyBuilder) BuildPartA() {
	m.product.PartA = "BuilderA_PartA"
}

func (m MyBuilder) BuildPartB() {
	m.product.PartB = "BuilderB_PartB"
}

func (m MyBuilder) BuildPartC() {
	m.product.PartC = "BuilderC_PartC"
}

func (m MyBuilder) GetProduct() *Product {
	return m.product
}

func newBuilder() *MyBuilder {
	return &MyBuilder{product: &Product{}}
}

type Director struct {
	builder Builder
}

func (d *Director) Construct() {
	d.builder.BuildPartA()
	d.builder.BuildPartB()
	d.builder.BuildPartC()
}

func newDirector(builder *MyBuilder) *Director {
	return &Director{builder: builder}
}

func builderStart() {
	builder1 := newBuilder()
	director := newDirector(builder1)
	director.Construct()
	builder1.GetProduct()
	fmt.Printf("Parts:\n%s %s %s", builder1.product.PartA, builder1.product.PartB, builder1.product.PartC)
}
