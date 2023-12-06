import React from "react";
import "./registerproperty.css";

const RegisterProperty = () => {
  const [errors, setErrors] = React.useState({});
  const [isPropertyRegistered, setIsPropertyRegistered] = React.useState(false);

  const RegisterPropertySubmit = (formData) => {
    fetch("http://localhost:8080/api/createproperty", {
      method: "post",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
      body: formData,
    })
      .then((response) => response.json())
      .then((data) => {
        if (data.message === "success") {
          setIsPropertyRegistered(true);
          setErrors({});
        } else if (data.message === "error") {
          console.log("Error:", data);
          setErrors(data.data);
        }
      })
      .catch((error) => {
        console.log("Error:", error);
      });
  };

  const handleFormSubmit = (e) => {
    e.preventDefault();
    const formData = new FormData(e.target);

    RegisterPropertySubmit(formData);
  };

  return (
      <div className="register-property">
        {isPropertyRegistered ? (
          <div className="success">
            <h2 className="success-title">Your property has been registered</h2>
            <p className="success-text">
              Your property has been registered and is pending approval from the
              admin.
            </p>
          </div>
        ) : (
          <div className="register-property-form">
            <h2 className="register-property-title">Register a property</h2>
            {Object.keys(errors).length > 0 && (
              <div className="register-property-error">
                <h3 className="register-property-error-title">Error:</h3>
                <p>{JSON.stringify(errors)}</p>
              </div>
            )}
            <form onSubmit={handleFormSubmit}>  
              <div>
                <label htmlFor="propertyName" className="register-property-label">Property Name:</label>
                <input type="text" id="propertyname" name="propertyname" placeholder="Name of the property" className="register-property-input" />
              </div>
              <div>
                <label htmlFor="type" className="register-property-label">Type:</label>
                <input type="text" id="type" name="type" placeholder="Type" className="register-property-input" />
              </div>
              <div>
                <label htmlFor="description" className="register-property-label">Description:</label>
                <input type="text" id="description" name="description" placeholder="Description" className="register-property-input" />
              </div>
              <div>
                <label htmlFor="price_per_night" className="register-property-label">Price Per Night:</label>
                <input type="text" id="price_per_night" name="price_per_night" placeholder="Price Per Night" className="register-property-input" />
              </div>
              <div>
                <label htmlFor="number_of_bedrooms" className="register-property-label">Number Of Bedrooms:</label>
                <input type="text" id="number_of_bedrooms" name="number_of_bedrooms" placeholder="Number Of Bedrooms" className="register-property-input" />
              </div>
              <div>
                <label htmlFor="number_of_bathrooms" className="register-property-label">Number Of Bathrooms:</label>
                <input type="text" id="number_of_bathrooms" name="number_of_bathrooms" placeholder="Number Of Bathrooms" className="register-property-input" />
              </div>
              <div>
                <label htmlFor="house_rules" className="register-property-label">House Rules:</label>
                <input type="text" id="house_rules" name="house_rules" placeholder="House Rules" className="register-property-input" />
              </div>
              <div>
                <label htmlFor="cancellation_policy" className="register-property-label">Cancellation Policy:</label>
                <input type="text" id="cancellation_policy" name="cancellation_policy" placeholder="Cancellation Policy" className="register-property-input" />
              </div>
              <div>
                <label htmlFor="location" className="register-property-label">Location:</label>
                <input type="text" id="location" name="location" placeholder="Location" className="register-property-input" />
              </div>
              <div>
                <label htmlFor="city" className="register-property-label">City:</label>
                <input type="text" id="city" name="city" placeholder="City" className="register-property-input" />
              </div>
              <div>
                <label htmlFor="state" className="register-property-label">State:</label>
                <input type="text" id="state" name="state" placeholder="State" className="register-property-input" />
              </div>
              <div>
                <label htmlFor="zip" className="register-property-label">Zip:</label>
                <input type="text" id="zip" name="zip" placeholder="Zip" className="register-property-input" />
              </div>
              <div>
                <label htmlFor="country" className="register-property-label">Country:</label>
                <input type="text" id="country" name="country" placeholder="Country" className="register-property-input" />
              </div>
              <div>
                <label htmlFor="image" className="register-property-label">Image:</label>
                <input type="file" id="image" name="image" placeholder="Image" className="register-property-input" />
              </div>
              <button type="submit" className="register-property-btn">Submit for approval</button>
            </form>
        </div>
      )}
    </div>
  );
};

export default RegisterProperty;
