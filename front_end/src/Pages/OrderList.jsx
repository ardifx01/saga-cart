import React, { useEffect, useState } from "react";
import axios from "axios";
import Wrapper from "../components/Wrapper";

function OrderList() {
  const [orders, setOrders] = useState([]);

  useEffect(() => {
    (async () => {
      try {
        const { data } = await axios.get("/api/orders/");
        let res = await Promise.all(
          data.map(async (val) => {
            let {
              data: { name },
            } = await axios.get(`/api/products/${val.product_id}`);
            return {
              ...val,
              product_id: name,
            };
          })
        );
        setOrders(res);
      } catch (error) {
        console.error("Failed to fetch orders:", error);
      }
    })();
  }, []);

  return (
    <Wrapper>
      <div className="p-10">
        <h1 className="text-4xl font-extrabold text-black mb-10">Order List</h1>
        <div className="overflow-x-auto">
          <table className="table-auto w-full border-collapse border border-black text-left">
            <thead className="bg-gray-800 text-white">
              <tr>
                <th className="border border-black px-4 py-2">No</th>
                <th className="border border-black px-4 py-2">Customer Name</th>
                <th className="border border-black px-4 py-2">Product</th>
                <th className="border border-black px-4 py-2">Quantity</th>
                <th className="border border-black px-4 py-2">Order Date</th>
                <th className="border border-black px-4 py-2">Status</th>
              </tr>
            </thead>
            <tbody>
              {orders.length > 0 ? (
                orders.map((order, idx) => (
                  <tr key={order.id} className="hover:bg-gray-100">
                    <td className="border border-black px-4 py-2">{idx + 1}</td>
                    <td className="border border-black px-4 py-2">
                      {order.customer_name}
                    </td>
                    <td className="border border-black px-4 py-2">
                      {order.product_id}
                    </td>
                    <td className="border border-black px-4 py-2">
                      {order.qty}
                    </td>
                    <td className="border border-black px-4 py-2">
                      {new Date(order.order_date).toLocaleDateString("id-ID", {
                        year: "numeric",
                        month: "long",
                        day: "numeric",
                      })}
                    </td>
                    <td className="border border-black px-4 py-2">
                      {order.status}
                    </td>
                  </tr>
                ))
              ) : (
                <tr>
                  <td
                    className="border border-black px-4 py-2 text-center"
                    colSpan="6"
                  >
                    No orders found.
                  </td>
                </tr>
              )}
            </tbody>
          </table>
        </div>
      </div>
    </Wrapper>
  );
}

export default OrderList;
