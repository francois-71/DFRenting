import React from "react";
import "./register.css";
import FormBox from "../../components/FormBox/FormBox";

const Register = () => {
  const [errors, setErrors] = React.useState({});
  const [isRegistered, setIsRegistered] = React.useState(false);
  // Fetch method post to localhost:8080/register with body of formData
  const RegisterSubmit = (formData) => {
    console.log("Register form data:", formData);
    fetch("http://localhost:8080/register", {
      method: "post",
      body: JSON.stringify(formData),
    })
      .then((response) => response.json())
      .then((data) => {
        if (data.message === "success") {
          setIsRegistered(true);
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

  const RegisterPlaceholders = {
    first_name: "First Name",
    last_name: "Last Name",
    email: "Email",
    password: "Password",
    phone: "Phone",
    age: "Age",
    address: "Address",
    city: "City",
    state: "State",
    zip: "Zip",
    country: "Country",
  };

  return (
    <div>
      {isRegistered ? (
        <div className="success">
          <h2>Success!</h2>
          <p>You have successfully registered.</p>
        </div>
      ) : (
        <div className="form">
          <FormBox
            title="Register"
            buttonText="Register"
            placeholders={RegisterPlaceholders}
            onSubmit={RegisterSubmit}
            errors={errors}
          />
        </div>
      )}
    </div>
  );
};

export default Register;
