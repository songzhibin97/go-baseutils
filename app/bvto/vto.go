package bvto

import "github.com/songzhibin97/gkit/tools/vto"

// VoToDoFromPoint src must point
func VoToDoFromPoint[DST any](src any) (*DST, error) {
	var zero DST
	err := vto.VoToDo(&zero, src)
	return &zero, err
}

// VoToDoFromNotPoint src must not point
func VoToDoFromNotPoint[DST any](src any) (*DST, error) {
	var zero DST
	err := vto.VoToDo(&zero, &src)
	return &zero, err
}

// VoToDoPlusFromPoint src must point
func VoToDoPlusFromPoint[DST any](src any, parameters vto.ModelParameters) (*DST, error) {
	var zero DST
	err := vto.VoToDoPlus(&zero, src, parameters)
	return &zero, err
}

// VoToDoPlusFromNotPoint src must not point
func VoToDoPlusFromNotPoint[DST any](src any, parameters vto.ModelParameters) (*DST, error) {
	var zero DST
	err := vto.VoToDoPlus(&zero, &src, parameters)
	return &zero, err
}

// VoToDoListFromPoint src must point
func VoToDoListFromPoint[DST any](src []any) ([]*DST, error) {
	zero := make([]*DST, 0, len(src))
	for _, v := range src {
		dv, err := VoToDoFromPoint[DST](v)
		if err != nil {
			return nil, err
		}
		zero = append(zero, dv)
	}
	return zero, nil
}

// VoToDoListFromNotPoint src must not point
func VoToDoListFromNotPoint[DST any](src []any) ([]*DST, error) {
	zero := make([]*DST, 0, len(src))
	for _, v := range src {
		dv, err := VoToDoFromNotPoint[DST](v)
		if err != nil {
			return nil, err
		}
		zero = append(zero, dv)
	}
	return zero, nil
}

// VoToDoListPlusFromPoint src must point
func VoToDoListPlusFromPoint[DST any](src []any, parameters vto.ModelParameters) ([]*DST, error) {
	zero := make([]*DST, 0, len(src))
	for _, v := range src {
		dv, err := VoToDoPlusFromPoint[DST](v, parameters)
		if err != nil {
			return nil, err
		}
		zero = append(zero, dv)
	}
	return zero, nil
}

// VoToDoListPlusFromNotPoint src must not point
func VoToDoListPlusFromNotPoint[DST any](src []any, parameters vto.ModelParameters) ([]*DST, error) {
	zero := make([]*DST, 0, len(src))
	for _, v := range src {
		dv, err := VoToDoPlusFromNotPoint[DST](v, parameters)
		if err != nil {
			return nil, err
		}
		zero = append(zero, dv)
	}
	return zero, nil
}
