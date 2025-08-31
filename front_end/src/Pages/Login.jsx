import axios from "axios";
import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

export default function Login() {
    const navigate = useNavigate();
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")

    const HandleLogin = async (e) => {
        e.preventDefault();

        try{
            const {
                data: {message, token},
            } = await axios.post("/api/auth/login", {
                username,
                password
            });
            localStorage.setItem("token", token);
            if(token) {
                navigate("/products") 
            } 
        } catch(error) {
            console.error(error)
        }
    } 


  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="w-full max-w-md bg-white shadow-lg rounded-2xl p-8">
        {/* Title */}
        <h2 className="text-2xl font-bold text-center text-slate-800 mb-6">
          Saga Cart Login
        </h2>

        {/* Form */}
        <form onSubmit={HandleLogin} className="space-y-4">
          {/* Username */}
          <div>
            <label className="block text-slate-700 font-medium mb-1">
              Username
            </label>
            <input
              type="text"
              onChange={(e) => setUsername(e.target.value)}
              placeholder="Enter username"
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-slate-600"
            />
          </div>

          {/* Password */}
          <div>
            <label className="block text-slate-700 font-medium mb-1">
              Password
            </label>
            <input
                onChange={(e) => setPassword(e.target.value)}
              type="password"
              placeholder="Enter password"
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-slate-600"
            />
          </div>

          {/* Login Button */}
          <button
            type="submit"
            className="w-full bg-amber-500 text-white py-2 rounded-lg font-semibold hover:bg-amber-600 transition"
          >
            Login
          </button>
        </form>
      </div>
    </div>
  );
}
