package validators
import ("github.com/go-playground/validator/v10"
		"errors"
		"net/http"
		"strings")


var Validate = validator.New()


func init() {
	Validate.RegisterValidation("email", CustomEmailValidation)
	Validate.RegisterValidation("password", CustomPasswordValidation)
	Validate.RegisterValidation("propertyname", CustomPropertyNameValidation)
	Validate.RegisterValidation("type", CustomTypeValidation)
	Validate.RegisterValidation("description", CustomDescriptionValidation)
	Validate.RegisterValidation("price_per_night", CustomPricePerNightValidation)
	Validate.RegisterValidation("number_of_bedrooms", CustomNumberOfBedroomsValidation)
	Validate.RegisterValidation("number_of_bathrooms", CustomNumberOfBathroomsValidation)
	Validate.RegisterValidation("house_rules", CustomHouseRulesValidation)
	Validate.RegisterValidation("cancellation_policy", CustomCancellationPolicyValidation)
	Validate.RegisterValidation("location", CustomLocationValidation)
	Validate.RegisterValidation("city", CustomCityValidation)
	Validate.RegisterValidation("state", CustomStateValidation)
	Validate.RegisterValidation("zip", CustomZipValidation)
	Validate.RegisterValidation("country", CustomCountryValidation)
	// Validate.RegisterValidation("date", CustomDateValidation)
	// Validate.RegisterValidation("reviews", CustomReviewsValidation)
	// Validate.RegisterValidation("image", CustomImageValidation)
	// Validate.RegisterValidation("is_active", CustomIsActiveValidation)
}

func IsValidImageFormat(header http.Header) error {
	contentType := header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		return errors.New("not an image file")
	}

	switch contentType {
	case "image/jpeg", "image/png", "image/svg+xml":
		return nil
	default:
		return errors.New("unsupported image format")
	}
}