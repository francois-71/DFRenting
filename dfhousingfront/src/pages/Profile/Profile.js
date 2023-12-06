import React from "react";
import "./profile.css";
import { useEffect } from "react";

const Profile = () => {
  const [profileData, setProfileData] = React.useState(null);

  useEffect(() => {
    const fetchProfileData = async () => {
      try {
        const response = await fetch("http://localhost:8080/api/user", {
          method: "get",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${localStorage.getItem("token")}`,
          },
        });

        if (!response.ok) {
          throw new Error("Error fetching profile data:", response);
        }

        const data = await response.json();
        console.log("Data:", data);
        setProfileData(data.data.data);
      } catch (error) {
        console.error("Error fetching profile data:", error);
      }
    };

    fetchProfileData();
  }, []);
  return (
    <div>
      <main className="main-content">
        <section className="profile-details">
          <h2>Details</h2>
          {profileData && (
            <div className="profile-details-container">
              <div className="profile">
                <div className="profile-button">
                  <div className="profile-add-property">
                    <a href="/register-property">Add a Property</a>
                  </div>
                  <div className="profile-modify-profile">
                    <a href="/modify-profile">Modify Profile</a>
                  </div>
                </div>
                <div className="profile-item">
                  <h3>First Name</h3>
                  <p>{profileData.first_name}</p>
                </div>
                <div className="profile-item">
                  <h3>Last Name</h3>
                  <p>{profileData.last_name}</p>
                </div>
                <div className="profile-item">
                  <h3>Email</h3>
                  <p>{profileData.email}</p>
                </div>
                <div className="profile-item">
                  <h3>Phone</h3>
                  <p>{profileData.phone}</p>
                </div>
                <div className="profile-item">
                  <h3>Age</h3>
                  <p>{profileData.age}</p>
                </div>
                <div className="profile-item">
                  <h3>Address</h3>
                  <p>{profileData.address}</p>
                </div>
                <div className="profile-item">
                  <h3>City</h3>
                  <p>{profileData.city}</p>
                </div>
                <div className="profile-item">
                  <h3>State</h3>
                  <p>{profileData.state}</p>
                </div>
                <div className="profile-item">
                  <h3>Zip</h3>
                  <p>{profileData.zip}</p>
                </div>
                <div className="profile-item">
                  <h3>Country</h3>
                  <p>{profileData.country}</p>
                </div>
                <div className="profile-item">
                  <h3>Role</h3>
                  <p>{profileData.role}</p>
                </div>
              </div>
            </div>
          )}
        </section>
      </main>
    </div>
  );
};

export default Profile;
