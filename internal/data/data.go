package data

import (
	"github.com/beeker1121/goque"
	"time"
)

type FlattenData struct {
	Timestamp      time.Time     `json:"timestamp"`
	InputArray     []interface{} `json:"input-array"`
	FlattenedArray []interface{} `json:"flattened-array"`
}

func SaveData(inputArray []interface{}, flattenedArray []interface{}) error {
	data := FlattenData{
		Timestamp:      time.Now(),
		InputArray:     inputArray,
		FlattenedArray: flattenedArray,
	}

	return data.saveData()
}

func (fd *FlattenData) saveData() error {
	s, err := goque.OpenStack("data_dir")
	if err != nil {
		return err
	}
	defer s.Close()

	_, err = s.PushObjectAsJSON(&fd)
	if err != nil {
		return err
	}

	return nil
}

func GetListSuccessfulProcessedArray(load int) ([]FlattenData, error) {

	s, err := goque.OpenStack("data_dir")
	if err != nil {
		return nil, err
	}
	defer s.Close()

	output := make([]FlattenData, 0, 1)

	for i := uint64(0); i < uint64(load); i++ {
		item, err := s.PeekByOffset(i)
		if err != nil {
			break
		}

		var obj FlattenData
		err = item.ToObjectFromJSON(&obj)
		if err != nil {
			return nil, err
		}

		output = append(output, obj)
	}

	return output, nil
}
