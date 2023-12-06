import React from 'react';
import './formbox.css'; // Make sure to have your CSS file path correctly

const FormBox = ({ title, buttonText, placeholders, onSubmit, errors}) => {
    const [formData, setFormData] = React.useState({});

    const handleInputChange = (e, key) => {
        setFormData({ ...formData, [key]: e.target.value });
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        onSubmit(formData);
    };

    const getInputType = (key) => {
        if (key === "password") {
            return "password";
        }
        else if (key === "email") { 
            return "email";
        }
        else {
            return "text";
        }
    }

    return (
        <div className="form-container">
            <h2>{title}</h2>
            <div className="errors">
                {Object.keys(errors).map((key) => (
                    <p key={key}>{errors[key]}</p>
                ))}
            </div>
            <form onSubmit={handleSubmit}>
                {Object.keys(placeholders).map((key) => (
                    <input
                        key={key}
                        type={getInputType(key)}
                        placeholder={placeholders[key]}
                        onChange={(e) => handleInputChange(e, key)}
                    />
                ))}
                <button type="submit">{buttonText}</button>
            </form>
        </div>
    );
};

export default FormBox;
