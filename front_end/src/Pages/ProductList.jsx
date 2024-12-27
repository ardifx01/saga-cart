import React, { useEffect, useState } from "react";
import Wrapper from "../components/Wrapper";
import axios from "axios";
import { useNavigate } from "react-router-dom";

function ProductList() {
  const navigate = useNavigate();
  const [products, setProducts] = useState([]);
  useEffect(() => {
    (async () => {
      try {
        const {
          data: { data },
        } = await axios.get("/api/products/");
        setProducts(data);
      } catch (error) {
        console.log(error);
      }
    })();
  }, []);

  const formatPriceToIDR = (price) => {
    return new Intl.NumberFormat("id-ID", {
      style: "currency",
      currency: "IDR",
    }).format(price);
  };

  const HandleAddToCard = (productId) => {
    navigate("/cart", {
      state: { productId },
    });
  };

  return (
    <Wrapper>
      <div className="flex-1 p-10">
        <h1 className="text-4xl font-extrabold text-black mb-10">
          Product List
        </h1>
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-10">
          {products.map((product) => (
            <div
              key={product.id}
              className="bg-gray-200 p-6 rounded-xl shadow-lg border-4 border-black flex flex-col items-center space-y-4 hover:bg-gray-300 transition-all"
            >
              <h2 className="text-2xl font-bold text-black">{product.name}</h2>
              <p className="text-lg text-gray-700">
                Quantity: {product.quantity}
              </p>
              <p className="text-xl font-bold text-black">
                {formatPriceToIDR(product.price)}
              </p>
              <button
                onClick={() => HandleAddToCard(product.id)}
                className="bg-yellow-500 text-white py-2 px-6 rounded-lg text-xl font-semibold border-4 border-black mt-4 hover:bg-yellow-400 transition-all"
              >
                Add to Cart
              </button>
            </div>
          ))}
        </div>
      </div>
    </Wrapper>
  );
}

export default ProductList;
