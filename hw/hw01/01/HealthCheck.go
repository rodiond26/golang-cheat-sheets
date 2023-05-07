package main

const (
	PassStatus = "pass"
	FailStatus = "fail"
)

type HealthCheck struct {
	ServiceID int
	Status    string
}

func GenerateCheck() (res []HealthCheck) {
	for i := 0; i < 5; i++ {
		hc := HealthCheck{
			ServiceID: i,
			Status:    status(i),
		}
		res = append(res, hc)
	}
	return res
}

func status(i int) string {
	if i%2 == 0 {
		return PassStatus
	}
	return FailStatus
}
