// Este archivo contiene pruebas unitarias para la función GetFullTimeEmployeeById.
// Se utilizan funciones mock para simular el comportamiento de dependencias externas,
// permitiendo así probar la lógica de integración de manera aislada.
package main

import (
	"testing"
)

// TestGetFullTimeEmployeeById verifica que la función GetFullTimeEmployeeById
// retorne correctamente un empleado de tiempo completo (FullTimeEmployee) usando mocks.
//
// Para cada caso de prueba:
//   - Se define un id y un dni de entrada.
//   - Se define una función mockFunc que reemplaza temporalmente las funciones GetPersonByDni y GetEmployeeById
//     para devolver datos simulados (mock), evitando así dependencias externas o acceso a datos reales.
//   - Se define el resultado esperado (want) como un FullTimeEmployee.
//
// El test ejecuta cada caso:
//  1. Aplica los mocks definidos en mockFunc.
//  2. Llama a GetFullTimeEmployeeById con los parámetros del caso.
//  3. Verifica que no haya errores.
//  4. Compara los campos age y salary del resultado con los valores esperados.
//  5. Al finalizar, restaura las funciones originales para no afectar otros tests.
//
// Este enfoque permite probar la lógica de integración de GetFullTimeEmployeeById de forma aislada,
// asegurando que la función combine correctamente los datos de las dependencias simuladas.
func TestGetFullTimeEmployeeById(t *testing.T) {
	testCases := []struct {
		id       int
		dni      string
		mockFunc func()
		want     FullTimeEmployee
	}{
		{
			id:  1,
			dni: "12345678A",
			mockFunc: func() {
				GetPersonByDni = func(dni string) (Person, error) {
					return Person{dni: dni, name: "John Doe", age: 30}, nil
				}
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{id: id, salary: 50000}, nil
				}
			},
			want: FullTimeEmployee{
				Employee: Employee{id: 1, salary: 50000},
				Person:   Person{dni: "12345678A", name: "John Doe", age: 30},
			},
		},
		{
			id:  2,
			dni: "87654321B",
			mockFunc: func() {
				GetPersonByDni = func(dni string) (Person, error) {
					return Person{dni: dni, name: "Jane Smith", age: 25}, nil
				}
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{id: id, salary: 60000}, nil
				}
			},
			want: FullTimeEmployee{
				Employee: Employee{id: 2, salary: 60000},
				Person:   Person{dni: "87654321B", name: "Jane Smith", age: 25},
			},
		},
	}

	// Guardar las implementaciones originales para restaurarlas al final del test
	originalGetPersonByDni := GetPersonByDni
	originalGetEmployeeById := GetEmployeeById

	// Ejecutar cada caso de prueba
	for _, test := range testCases {
		// Aplicar los mocks definidos para este caso
		test.mockFunc()
		// Ejecutar la función bajo prueba
		ftEmployee, err := GetFullTimeEmployeeById(test.id, test.dni)
		if err != nil {
			t.Errorf("GetFullTimeEmployeeById(%d, %s) retornó error: %v", test.id, test.dni, err)
			return
		}

		// Verificar que los campos age y salary coincidan con lo esperado
		if ftEmployee.age != test.want.age {
			t.Errorf("Error en age: se obtuvo %d, se esperaba %d", ftEmployee.age, test.want.age)
		}
		if ftEmployee.salary != test.want.salary {
			t.Errorf("Error en salary: se obtuvo %.2f, se esperaba %.2f", ftEmployee.salary, test.want.salary)
		}
	}

	// Restaurar las funciones originales para no afectar otros tests
	GetPersonByDni = originalGetPersonByDni
	GetEmployeeById = originalGetEmployeeById
	t.Log("¡Todas las pruebas pasaron exitosamente!")
}
