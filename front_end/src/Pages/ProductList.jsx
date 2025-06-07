import React, { useEffect, useState } from "react";
import Wrapper from "../components/Wrapper";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import { debounce } from "lodash";

function ProductList() {
  const navigate = useNavigate();
  const [products, setProducts] = useState([]);
  const [searchQuery, setSearchQuery] = useState("");
  const [isSearching, setIsSearching] = useState(false);

  // Fetch products (regular or search)
  const fetchProducts = async (query = "") => {
    try {
      setIsSearching(true);
      const endpoint = query 
        ? `http://localhost:8081/api/products/search?q=${encodeURIComponent(query)}`
        : `/api/products/`;
      
      const { data: { data } } = await axios.get(endpoint);
      setProducts(data);
    } catch (error) {
      console.error("Error fetching products:", error);
    } finally {
      setIsSearching(false);
    }
  };

  // Debounced search
  const debouncedSearch = debounce((query) => {
    fetchProducts(query);
  }, 500);

  useEffect(() => {
    fetchProducts();
  }, []);

  useEffect(() => {
    if (searchQuery) {
      debouncedSearch(searchQuery);
    } else {
      fetchProducts();
    }

    return () => debouncedSearch.cancel();
  }, [searchQuery]);

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
        <div className="flex justify-between items-center mb-10">
          <h1 className="text-4xl font-extrabold text-black">Product List</h1>
          <div className="relative w-1/3">
            <input
              type="text"
              placeholder="Search products..."
              className="w-full p-3 pl-10 rounded-lg border-2 border-gray-300 focus:border-yellow-500 focus:outline-none"
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
            />
            <svg
              className="absolute left-3 top-3.5 h-5 w-5 text-gray-400"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
              />
            </svg>
          </div>
        </div>

        {isSearching ? (
          <div className="flex justify-center items-center h-64">
            <p className="text-xl">Searching products...</p>
          </div>
        ) : products.length === 0 ? (
          <div className="flex justify-center items-center h-64">
            <p className="text-xl">No products found</p>
          </div>
        ) : (
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
        )}
      </div>
    </Wrapper>
  );
}

export default ProductList;