import "./App.css";
import Login from "./pages/Login/Login";
import Logout from "./pages/Logout/Logout";
import Register from "./pages/Register/Register";
import RegisterProperty from "./pages/RegisterProperty/RegisterProperty";
import PropertyInfo from "./pages/PropertyInfo/PropertyInfo";
import ProfilePage from "./pages/Profile/Profile";
import AdminPanel from "./pages/AdminPanel/AdminPanel";
import Home from "./pages/Home/Home";
import Header from "./components/Header/Header";
import Footer from "./components/Footer/Footer";
import { Routes, Route } from "react-router-dom";
import { useState, useEffect } from "react";


function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  useEffect(() => {
    const checkIsLoggedIn = async () => {
      try {
        const response = await fetch("http://localhost:8080/isloggedin", {
          method: "get",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${localStorage.getItem("token")}`,
          },
        });
        if (!response.ok) {
          console.log("Error fetching profile data:", response);
        }
        const data = await response.json();
        console.log("Data login:", data);
        if (data.message === "success") {
          setIsLoggedIn(true); 
        } else {
          setIsLoggedIn(false); 
        }
      } catch (error) {
        console.error("Error fetching profile data:", error);
      }
    };
    checkIsLoggedIn();
  }, []); // Empty dependency array, runs once on component mount

  return (
    <div className="App">
      <Header isLoggedIn={isLoggedIn} />
      <div className="pages-container">
        <Routes>
          {isLoggedIn ? (
            <>
            <Route path="/logout" element={<Logout />} />
            <Route path="/profile" element={<ProfilePage />} />
            
            <Route path="/admin" element={<AdminPanel />} />
            </>
          ) : (
            <>
            <Route path="/register" element={<Register />} />
            </>
          )}
          <Route path="/register-property" element={<RegisterProperty />} />
          <Route path="/property/:id" element={<PropertyInfo />} />
          <Route path="/login" element={<Login />} />
          <Route path="/home" element={<Home />} />
          <Route path="*" element={<Home />} />
        </Routes>
        <Footer />
      </div>
    </div>
  );
}

export default App;
