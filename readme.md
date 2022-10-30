# Simple Package For Return Response Validator

## How to Use

``  go get vandyahmad24/validator-formater ``

example implementation
```
type CakeInput struct {
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Rating      float32 `json:"float" validate:"required"`
	Image       string  `json:"image" validate:"required"`
}

func ValidateInputCake(input CakeInput) interface{} {
	var errors interface{}
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		errors = formater.FormatErrorValidation(err, "Request Wrong")
	}
	return errors
}

```


### Will Be Return response 
```
{
    "status": "Bad Request",
    "message": "",
    "data": [
        "Description is required",
        "Rating is required",
        "Image is required"
    ]
}

```