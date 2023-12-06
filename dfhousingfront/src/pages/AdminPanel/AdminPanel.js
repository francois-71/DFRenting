import React, { useState, useEffect } from "react";
import "./adminpanel.css";

const AdminPanel = () => {
  const [propertiesRequiresApproval, setPropertiesRequiresApproval] = useState(
    []
  );
  const [isAdmin, setIsAdmin] = useState(false);

  useEffect(() => {
    const checkAdminRole = async () => {
      try {
        const response = await fetch("http://localhost:8080/api/user/isadmin", {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${localStorage.getItem("token")}`,
          },
        });
        const data = await response.json();
        if (!response.ok) {
          throw new Error(data.data.message);
        }
        console.log("data for admin", data.data.data);
        setIsAdmin(data.data.data);
      } catch (err) {
        console.log(err);
      }
    };
    checkAdminRole();
  }, []);

  useEffect(() => {
    const fetchProperties = async () => {
      try {
        const response = await fetch(
          "http://localhost:8080/api/properties/requireapproval",
          {
            method: "GET",
            headers: {
              "Content-Type": "application/json",
              Authorization: `Bearer ${localStorage.getItem("token")}`,
            },
          }
        );
        const data = await response.json();
        if (!response.ok) {
          throw new Error(data.data.message);
        }
        console.log("data for admin", data.data.data);
        setPropertiesRequiresApproval(data.data.data);
      } catch (err) {
        console.log(err);
      }
    };
    fetchProperties();
  }, []);

  const approveProperty = async (propertyId) => {
    try {
      const response = await fetch(
        `http://localhost:8080/api/properties/approve/${propertyId}`,
        {
          method: "PATCH",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${localStorage.getItem("token")}`,
          },
        }
      );
      const data = await response.json();
      if (!response.ok) {
        throw new Error(data.data.message);
      }
      // remove property from the property list
      setPropertiesRequiresApproval(
        propertiesRequiresApproval.filter(
          (property) => property.id !== propertyId
        )
      );
    } catch (err) {
      console.log(err);
    }
  };

  const rejectProperty = async (propertyId) => {
    try {
      const response = await fetch(
        `http://localhost:8080/api/properties/reject/${propertyId}`,
        {
          method: "PATCH",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${localStorage.getItem("token")}`,
          },
        }
      );
      const data = await response.json();
      if (!response.ok) {
        throw new Error(data.data.message);
      }
      // remove property from the property list

      setPropertiesRequiresApproval(
        propertiesRequiresApproval.filter(
          (property) => property.id !== propertyId
        )
      );
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <div>
      {isAdmin ? (
        <div className="adminpanel-container">
          <div className="adminpanel-title">
            <h1>Admin Panel</h1>
          </div>
          <div className="adminpanel-properties">
            <h2>Properties</h2>
            <div className="adminpanel-properties-list">
              {propertiesRequiresApproval.map((property) => {
                return (
                  <div
                    className="adminpanel-properties-list-item"
                    key={property.id}
                  >
                    <div className="adminpanel-properties-list-item-hostname">
                      <h3>
                        <strong>Host:</strong> {property.hostname}
                      </h3>
                    </div>
                    <div className="adminpanel-properties-list-item-propertyname">
                      <h3>
                        <strong>Property:</strong> {property.propertyname}
                      </h3>
                    </div>
                    <div className="adminpanel-properties-list-item-location">
                      <h3>
                        <strong>Location:</strong> {property.location}
                      </h3>
                    </div>
                    <div className="adminpanel-properties-list-item-price">
                      <h3>
                        <strong>Price:</strong> {property.price_per_night}
                      </h3>
                    </div>
                    <div className="adminpanel-properties-list-item-bedrooms">
                      <h3>
                        <strong>Number of bedrooms:</strong>{" "}
                        {property.number_of_bedrooms}
                      </h3>
                    </div>
                    <div className="adminpanel-properties-list-item-bathrooms">
                      <h3>
                        <strong>Number of bathrooms:</strong>{" "}
                        {property.number_of_bathrooms}
                      </h3>
                    </div>
                    <div className="adminpanel-properties-list-item-type">
                      <h3>
                        <strong>Property type:</strong> {property.type}
                      </h3>
                    </div>
                    <div className="adminpanel-properties-list-item-description">
                      <h3>
                        <strong>Property description:</strong>{" "}
                        {property.description}
                      </h3>
                    </div>
                    <div className="adminpanel-properties-list-item-rules">
                      <h3>
                        <strong>Property rules:</strong> {property.house_rules}
                      </h3>
                    </div>
                    <div className="adminpanel-properties-list-item-cancellation">
                      <h3>
                        <strong>Property cancellation policy:</strong>{" "}
                        {property.cancellation_policy}
                      </h3>
                    </div>
                    <div className="adminpanel-properties-list-item-date">
                      <h3>
                        <strong>Property date:</strong> {property.date}
                      </h3>
                    </div>
                    <div className="adminpanel-properties-list-item-city">
                      <h3>
                        <strong>Property city:</strong> {property.city}
                      </h3>
                    </div>
                    <div className="adminpanel-properties-list-item-country">
                      <h3>
                        <strong>Property country:</strong> {property.country}
                      </h3>
                    </div>
                    <div className="adminpanel-properties-list-item-state">
                      <h3>
                        <strong>Property state:</strong> {property.state}
                      </h3>
                    </div>
                    <div className="adminpanel-properties-list-item-zip">
                      <h3>
                        <strong>Property zip:</strong> {property.zip}
                      </h3>
                    </div>
                    <div className="adminpanel-properties-list-item-image">
                      <img src={property.image.imageurl} alt="property" />
                    </div>
                    <div className="adminpanel-properties-list-item-buttons">
                      <button
                        className="adminpanel-properties-list-item-buttons-approve"
                        onClick={() => approveProperty(property.id)}
                      >
                        Approve
                      </button>
                      <button
                        className="adminpanel-properties-list-item-buttons-reject"
                        onClick={() => rejectProperty(property.id)}
                      >
                        Reject
                      </button>
                    </div>
                  </div>
                );
              })}
            </div>
          </div>
        </div>
      ) : (
        <div className="admin-access-denied">
          <p>Access denied.</p>
        </div>
      )}
    </div>
  );
};

export default AdminPanel;
