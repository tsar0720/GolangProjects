package main

import "fmt"

const PassStatus = "pass"
const FailStatus = "fail"

type HealthCheck struct {
	ServiceId int
	Status    string
}

func GenerateCheck() []HealthCheck {
	healthCheckMass := make([]HealthCheck, 5)

	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			healthCheckMass[i].ServiceId = i
			healthCheckMass[i].Status = PassStatus
			fmt.Println(healthCheckMass[i])

		} else {
			healthCheckMass[i].ServiceId = i
			healthCheckMass[i].Status = FailStatus
		}

	}
	return healthCheckMass
}

func main() {

	check := GenerateCheck()
	for _, k := range check {
		fmt.Println(k.ServiceId, k.Status)
	}
	fmt.Println("Тут будет выведен идентификатор")

}

func return1(a bool) int {
	if a == true {
		return 1
	}
	return 0
}
