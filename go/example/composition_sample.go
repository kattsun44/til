package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee            // 型のみ記述 (= 埋め込みフィールド) によって Name と ID が加わる
	Reports  []Employee // 部下 (報告対象者)
}

func (m Manager) FindNewEmployees() []Employee {
	newEmployees := []Employee{
		Employee{
			"石田三成",
			"13112",
		},
		Employee{
			"徳川家康",
			"13115",
		},
	}
	return newEmployees
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "豊臣秀吉",
			ID:   "12345",
		},
		Reports: []Employee{},
	}
	fmt.Println(m.ID)
	fmt.Println((m.Description()))
	m.Reports = []Employee{
		Employee{
			"石田三成",
			"13112",
		},
		Employee{
			"徳川家康",
			"13115",
		},
	}
	fmt.Println(m.Employee)
	fmt.Println(m.Reports)
}
