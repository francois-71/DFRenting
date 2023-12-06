// PropertyInfo.js
import React from "react";
import { useLocation } from "react-router-dom";
import "./propertyInfo.css";

const PropertyInfo = () => {
  const { state } = useLocation();
  const property = state;
  console.log("property", property);
  console.log("property type", property.type);

  return (
    <div>
      {property ? (
        <div className="property-details">
          <h2 className="property-details-name">Hostname: {property.name}</h2>
          {property.image && (
            <div>
              <div className="property-details-image">
                <img src={property.image.imageurl} alt="Property" />
              </div>
            </div>
          )}
          <p>Hostname: {property.hostname}</p>
          <p>Type: {property.type}</p>
          <p>Description: {property.description}</p>
          <p>Price per Night: {property.price_per_night}</p>
          <p>Number of Bedrooms: {property.number_of_bedrooms}</p>
          <p>Number of Bathrooms: {property.number_of_bathrooms}</p>
          <p>House Rules: {property.house_rules}</p>
          <p>Cancellation Policy: {property.cancellation_policy}</p>
          <p>Location: {property.location}</p>
          <p>City: {property.city}</p>
          <p>State: {property.state}</p>
          <p>Zip: {property.zip}</p>
          <p>Country: {property.country}</p>
          {/* Display other details as needed */}
          {property.reviews && (
            <div>
              <h3>Reviews:</h3>
              <ul>
                {property.reviews.map((review, index) => (
                  <li key={index}>
                    {review.reviewText} - {review.rating}
                  </li>
                ))}
              </ul>
            </div>
          )}
        </div>
      ) : (
        <p>No property details found.</p>
      )}
    </div>
  );
};

export default PropertyInfo;
