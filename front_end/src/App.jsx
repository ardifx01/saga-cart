import { useState } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import ProductList from "./Pages/ProductList";
import OrderList from "./Pages/OrderList";
import Cart from "./Pages/Cart";
import Login from "./Pages/Login";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Navigate to="/products" replace />} />
        <Route path={"/products"} element={<ProductList />} />
        <Route path={"/orders"} element={<OrderList />} />
        <Route path={"/cart"} element={<Cart />} />
        <Route path={"/login"} element={<Login />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
