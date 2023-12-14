package main

import "fmt"

type Colegio interface {
	Profesores() Contrato
	Alumnos()
}
type Contrato interface {
	ProfesoresIndefinido() string
	ProfesoresTemporal() string
}
type Planta struct {
}

func (p Planta) Alumnos() {
	fmt.Println(" Alumno con mas de 3 años de antiguedad ")
}
func (p Planta) Profesores() Contrato {
	return PlantaProfesores{}

}

type PlantaProfesores struct {
}

func (PlantaProfesores) ProfesoresIndefinido() string {
	return " Profesor de secundaria con contrato a termino indefinido "
}
func (PlantaProfesores) ProfesoresTemporal() string {
	return " Profesor de secundaria con contrato a un año "
}

type PrimariaColegio struct {
}

func (PrimariaColegio) Alumnos() {
	fmt.Println(" Alumno de basica primaria ")
}

func (PrimariaColegio) Profesores() Contrato {
	return PrimariaColegioContrato{}
}

type PrimariaColegioContrato struct {
}

func (PrimariaColegioContrato) ProfesoresIndefinido() string {
	return " Profesor de basica primaria con contrato a termino indefinido "
}

func (PrimariaColegioContrato) ProfesoresTemporal() string {
	return " Profesor de basica primaria con contrato a un año "
}

func ColegioFactory(Consulta string) (Colegio, error) {
	var ConsultaRealizar int
	fmt.Println(" Ingresa la consulta a realizar: (1) Primaria, (2) Planta ")
	fmt.Scan(&ConsultaRealizar)
	fmt.Println(" Su consulta es de :", ConsultaRealizar)

	if ConsultaRealizar == 1 {
		return &PrimariaColegio{}, nil
	}
	if ConsultaRealizar == 2 {
		return &Planta{}, nil
	}

	return nil, fmt.Errorf(" Consulta Invalida ")
}
func main() {

	consulta := "Planta"

	colegio, errorFactory := ColegioFactory(consulta)
	if errorFactory != nil {
		panic(errorFactory)
	}
	colegio.Alumnos()
	c := colegio.Profesores().ProfesoresIndefinido()
	d := colegio.Profesores().ProfesoresTemporal()

	contrato := colegio.Profesores()
	a := contrato.ProfesoresIndefinido()
	b := contrato.ProfesoresTemporal()

	fmt.Println(a, b)
	fmt.Println(c, d)

}
