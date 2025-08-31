import axios from "axios";
import React, { useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";

const Wrapper = ({ children }) => {
  const navigate = useNavigate();
  const [currentLoginUser, setCurrentLoginUser] = useState(null);

  const HandleLogout = () => {
    const token = localStorage.getItem("token");
    if (token) localStorage.removeItem("token");
    navigate("/login");
  };

  useEffect(() => {
    (async () => {
      try {
        const token = localStorage.getItem("token");
        let {data} = await axios.get("/api/auth/me",{
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
        setCurrentLoginUser(data);
      } catch (error) {}
    })();
  }, []);
 
  return (
    <div className="flex h-screen">
      {/* Sidebar */}
      <div className="w-64 bg-gray-900 p-5">
        <h2 className="text-white text-xl font-bold mb-10">Saga Cart</h2>
        <ul className="space-y-4">
          <li>
            <Link
              to="/products"
              className="text-white font-bold hover:text-gray-400"
            >
              Products
            </Link>
          </li>
          <li>
            <Link
              to="/orders"
              className="text-white font-bold hover:text-gray-400"
            >
              Orders
            </Link>
          </li>
          {/* {currentLoginUser.name} */}
        </ul>
      </div>

      {/* Main Content */}
      <div className="flex-1 flex flex-col">
        {/* Header */}
        <div className="flex justify-between items-center px-10 py-4 border-b bg-gray-800">
          <h1 className="text-white text-xl font-bold">Saga Cart</h1>
          {currentLoginUser ? (
            <div className="flex">
              <span className="text-white mr-5">{currentLoginUser.username}</span>
               <button onClick={HandleLogout} className="bg-gray-700 hover:bg-gray-600 text-white font-bold py-2 px-4 rounded">
            Logout
          </button>
            </div>
          ): (
            <Link
            to="/login"
            className="bg-gray-700 hover:bg-gray-600 text-white font-bold py-2 px-4 rounded"
          >
            Login
          </Link>
          )}
        </div>

        {/* Content Area */}
        <div className="flex-1 p-10 bg-gray-100 overflow-auto">{children}</div>
      </div>
    </div>
  );
};

export default Wrapper;
