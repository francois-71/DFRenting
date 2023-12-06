import React from "react";
import "./propertycard.css";
import { Link } from "react-router-dom";

const PropertyCard = ({ property }) => {
  return (
    <div className="property-card-container">
      <Link
        className="property-card-link"
        to={{ pathname: `/property/${property.id}` }}
        state={ property }
      >
        <div className="property-card">
          <div className="property-card-image">
            {property.image && (
              <img src={property.image.imageurl} alt="Property" />
            )}
          </div>
          <div className="property-card-info">
            <p>
              <strong>{property.city}</strong>,{" "}
              <strong>{property.country}</strong>
            </p>
            <p>{property.price_per_night} € / night</p>
          </div>
        </div>
      </Link>
    </div>
  );
};

export default PropertyCard;
