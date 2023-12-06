import React from 'react';
import FormBox from '../../components/FormBox/FormBox';
import './login.css';

const Login = () => {
    const [errors, setErrors] = React.useState({});


    const loginSubmit = (formData) => {
        fetch("http://localhost:8080/login", {
          method: "post",
          body: JSON.stringify(formData),
        })
          .then((response) => response.json())
          .then((data) => {
            if (data.message === "success") {
                console.log("token:", data.data.token);
                localStorage.setItem("token", data.data.token);
                window.location.href = "/profile";
            }
            else if (data.message === "error"){
              console.log("Error:", data);
              setErrors(data.data);
            }
            console.log("Login form data:", data);
          })
          .catch((error) => {
            console.log("Error:", error);
          });
      };

    const loginPlaceholders = {
        email: 'Enter your email',
        password: 'Enter your password',
    };

   

    return (
        <div className="form">
            <FormBox
                title="Login"
                buttonText="Login"
                placeholders={loginPlaceholders}
                onSubmit={loginSubmit}
                errors={errors}
                _type="password"
            />
        </div>
    );
};

export default Login;