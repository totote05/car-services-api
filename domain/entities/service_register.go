package entities

type (
	ServiceRegister struct {
		ID        ServiceRegisterID
		VehicleID VehicleID
		ServiceID ServiceID
		Km        Km
	}
	ServiceRegisterID string
)
