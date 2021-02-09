package flatteners

type Flattener interface {
	FlattenArray(interface{}) ([]interface{}, int)
}

type FirstAlgorithm []interface{}

func (f *FirstAlgorithm) FlattenArray() ([]interface{}, int){
	var outputArray []interface{}
	var depth = 0

	outputArray = *f

	depth = getDepth(outputArray)

	return getFlattenArray(outputArray), depth
}

func getDepth(input interface{}) int{
	var maxDepth = 0

	for _, val := range input.([]interface{}){
		if _, ok := val.([]interface{}); ok{
			depth := getDepth(val)
			if depth > maxDepth {
				maxDepth = depth
			}
		}
	}

	return maxDepth + 1
}

func getFlattenArray(input interface{}) []interface{}{
	output := make([]interface{},0, 0)

	for _, val := range input.([]interface{}){
		if _, ok := val.([]interface{}); ok{
			output = append(output, getFlattenArray(val)...)
		} else {
			output = append(output, val)
		}
	}

	return output
}