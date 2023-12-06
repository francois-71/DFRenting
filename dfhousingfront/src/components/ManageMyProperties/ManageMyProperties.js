import React from 'react';
import './managemyproperties.css';

const ManageMyProperties = () => {

    const [myProperties, setMyProperties] = useState([])
    const [propertiesLoaded, isPropertiesLoaded] = useState(false)

    useEffect(() => {
        const getMyProperties = async () => {
            try{
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
                setMyProperties(data.data.data);
                isPropertiesLoaded(true);
            }
            catch (error) {
                console.error("Error:", error);
            }
        }
    }, []);

    return (
        <div>
            <h1>Manage My Properties</h1>
            
        </div>
    );
};

export default ManageMyProperties;