package validator

func ComposeValidators[T interface{}](validators ...func(T) error) func(T) error {
	return func(value T) error {
		for _, validator := range validators {
			if err := validator(value); err != nil {
				return err
			}
		}
		return nil
	}
}
