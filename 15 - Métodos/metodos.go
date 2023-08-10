package main

import "fmt"

type user struct {
	name string
	age  uint
}

//métodos estão vinculados à estruturas
func (u user) showAge() {
	fmt.Println("A idade é", u.age)
}

func (u user) isUnderAge() bool {
	return u.age >= 18
}

//transformando o parametro em ponteiro, iremos alterar a referencia do struct
func (u *user) incrementAge() {
	u.age++
}

func main() {
	user1 := user{"Gabriel Gambôa", 22}
	user1.showAge()

	if isUnderAge := user1.isUnderAge(); isUnderAge {
		fmt.Println("Maior de idade :)")
	}

	user1.incrementAge()
	fmt.Println(user1)
}
