import React, { useEffect, useState } from "react";
import "./home.css";
import PropertyCard from "../../components/PropertyCard/PropertyCard";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";



const Home = () => {
  const [propertiesLoaded, setPropertiesLoaded] = useState(false);
  const [properties, setProperties] = useState([]);

  const [startDate, setStartDate] = useState(new Date());
  const [endDate, setEndDate] = useState(new Date().setDate(new Date().getDate() + 1));

  const handleDateChange = (dates) => {
    const [start, end] = dates;
    setStartDate(start);
    setEndDate(end);
  };

  useEffect(() => {
    const getAllProperties = async () => {
      try {
        const response = await fetch("http://localhost:8080/api/properties", {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${localStorage.getItem("token")}`,
          },
        });

        if (!response.ok) {
          throw new Error("Failed to fetch properties");
        }

        const data = await response.json();
        console.log(data);
        setProperties(data.data.data);
        setPropertiesLoaded(true);
      } catch (error) {
        console.error("Error:", error);
      }
    };

    getAllProperties();
  }, []);

  return (
    <div className="home-container">
      <div className="home-container-title">
        <h1>Welcome to DFRenting</h1>
        <h5>Where should I bring you?</h5>
        <form action="">
          <input type="text" placeholder="Enter your destination" />
          <DatePicker
            selected={startDate}
            onChange={handleDateChange}
            startDate={startDate}
            endDate={endDate}
            selectsRange
            placeholderText="Select dates range"
          />
          <input type="number" placeholder="Number of guests" />
          <button type="submit">Search</button>
        </form>
      </div>
      <div className="home-properties">
        {propertiesLoaded ? (
          properties.map((property) => (
            <div className="home-properties-individual"key={property.id}>
              <PropertyCard property={property} />
            </div>
          ))
        ) : (
          <p>Loading...</p>
        )}
      </div>
    </div>
  );
};

export default Home;
