type Facility struct {
	ID   int    `json:"facility_id"`
	Name string `json:"facility_name"`
	// other fields
}

func getFacilityByID(data []byte, id int) (*Facility, error) {
	var facilities []Facility
	err := json.Unmarshal(data, &facilities)
	if err != nil {
		return nil, err
	}

	for _, facility := range facilities {
		if facility.ID == id {
			return &facility, nil
		}
	}

	return nil, fmt.Errorf("facility with ID %d not found", id)
}

