import React from "react";
import { Link } from "react-router-dom";

const Wrapper = ({ children }) => {
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
          {/* <li>
            <Link to="/cart" className="text-white font-bold hover:text-gray-400">Cart</Link>
          </li> */}
          <li>
            <Link
              to="/orders"
              className="text-white font-bold hover:text-gray-400"
            >
              Orders
            </Link>
          </li>
        </ul>
      </div>

      {/* Main Content */}
      <div className="flex-1 p-10">{children}</div>
    </div>
  );
};

export default Wrapper;
