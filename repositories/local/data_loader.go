package local

import (
	"encoding/json"
	"os"

	"car-services-api.totote05.ar/domain/entities"
)

var (
	ld *localData = nil
)

type (
	vehicleData struct {
		entities.Vehicle
		RegisteredKm     entities.KmList            `json:"registered_km"`
		ServiceRegisters []entities.ServiceRegister `json:"registered_services"`
	}
	localData struct {
		Vehicle  []vehicleData      `json:"vehicle"`
		Services []entities.Service `json:"services"`
	}
)

func getLocalData() localData {
	if ld != nil {
		return *ld
	}

	ld = &localData{}
	if data, err := loadLocalData(); err == nil {
		ld = data
	}

	return *ld
}

func loadLocalData() (*localData, error) {
	file, err := os.Open(os.Getenv("DATA_DIR") + "/data.json")
	if err != nil {
		return nil, err
	}

	defer file.Close()
	ld := &localData{}
	if err = json.NewDecoder(file).Decode(ld); err != nil {
		return nil, err
	}

	return ld, nil
}
