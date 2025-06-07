import React, { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import Wrapper from "../components/Wrapper";
import axios from "axios";

function Cart() {
  const location = useLocation();
  const productId = location.state?.productId;

  const [product, setProduct] = useState({ name: "" });
  const [customerName, setCustomerName] = useState("");
  const [quantity, setQuantity] = useState("");
  const [money, setMoney] = useState("");
  const [alert, setAlert] = useState("");

  useEffect(() => {
    (async () => {
      const { data } = await axios.get(`/api/products/${productId}`);
      setProduct(data);
    })();
  }, [productId]);

  const formatWithCommas = (numberString) => {
    const cleaned = numberString.replace(/[^\d]/g, "");
    return cleaned.replace(/\B(?=(\d{3})+(?!\d))/g, ",");
  };

  const handleMoneyChange = (e) => {
    const formatted = formatWithCommas(e.target.value);
    setMoney(formatted);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    const availableMoney = parseInt(money.replace(/,/g, ""), 10) || 0;

    try {
      await axios.post("/api/orders/", {
        customer_name: customerName,
        product_id: productId,
        qty: +quantity,
        amount: availableMoney, // send raw money to backend if needed
      });
      setAlert("Product successfully added to the cart!");
      setCustomerName("");
      setQuantity("");
      setMoney("");
    } catch (error) {
      setAlert("Failed to add product to the cart. Please try again.");
    }

    setTimeout(() => setAlert(""), 3000);
  };

  return (
    <Wrapper>
      <div className="p-8">
        <h1 className="text-3xl font-extrabold text-gray-800 mb-6">Cart</h1>
        <p className="text-1xl mb-4">
          {product.description}
        </p>

        {alert && (
          <div
            className={`mb-4 p-4 rounded-lg text-white font-bold ${
              alert.startsWith("Product successfully")
                ? "bg-green-500"
                : "bg-red-500"
            }`}
          >
            {alert}
          </div>
        )}

        <form
          onSubmit={handleSubmit}
          action="POST"
          className="bg-gray-100 p-6 rounded-lg border-4 border-black shadow-lg"
        >
          <div className="mb-4">
            <label
              htmlFor="customerName"
              className="block text-lg font-bold text-gray-800"
            >
              Customer Name
            </label>
            <input
              id="customerName"
              type="text"
              placeholder="Enter your name"
              value={customerName}
              onChange={(e) => setCustomerName(e.target.value)}
              className="w-full p-2 border-2 border-black rounded-lg shadow-md text-gray-800 focus:ring-2 focus:ring-black"
              required
            />
          </div>

          <div className="mb-4">
            <label
              htmlFor="productName"
              className="block text-lg font-bold text-gray-800"
            >
              Product Name
            </label>
            <input
              id="productName"
              type="text"
              value={product.name}
              disabled
              className="w-full p-2 bg-gray-200 border-2 border-black rounded-lg shadow-md text-gray-800 cursor-not-allowed"
            />
          </div>

          <div className="mb-4">
            <label
              htmlFor="quantity"
              className="block text-lg font-bold text-gray-800"
            >
              Quantity
            </label>
            <input
              id="quantity"
              type="number"
              placeholder="Enter quantity"
              value={quantity}
              onChange={(e) => setQuantity(e.target.value)}
              className="w-full p-2 border-2 border-black rounded-lg shadow-md text-gray-800 focus:ring-2 focus:ring-black"
              min="1"
              required
            />
          </div>

          <div className="mb-4">
            <label
              htmlFor="money"
              className="block text-lg font-bold text-gray-800"
            >
              Enter Your Money (IDR)
            </label>
            <input
              id="money"
              type="text"
              placeholder="e.g., 10,000"
              value={money}
              onChange={handleMoneyChange}
              inputMode="numeric"
              className="w-full p-2 border-2 border-black rounded-lg shadow-md text-gray-800 focus:ring-2 focus:ring-black"
              required
            />
          </div>

          <button
            type="submit"
            className="w-full py-2 px-4 bg-yellow-500 text-black font-bold border-4 border-black rounded-lg shadow-lg hover:bg-yellow-600 transition"
          >
            Add to Cart
          </button>
        </form>
      </div>
    </Wrapper>
  );
}

export default Cart;
