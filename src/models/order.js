class Order {
  static async AddToOrder(xcsrf) {
    const data = {
      addDateOrder: new Date(),
    };

    return await fetch(`/api/order/add`, {
      method: "POST",
      body: JSON.stringify(data),
      headers: {
        "X-CSRF-TOKEN": xcsrf,
        "Content-Type": "application/json",
      },
    });
  }

  static async GetOrders() {
    return await fetch(`/api/order/all`, {
      method: "GET",
    });
  }

  static async DeleteOrder(xcsrf, orderId) {
    return await fetch(`/api/order/delete?orderId=${orderId}`, {
      method: "DELETE",
      headers: {
        "X-CSRF-TOKEN": xcsrf,
      },
    });
  }
}

export default Order;
