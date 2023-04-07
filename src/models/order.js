import Auth from "./auth";

class Order {
  static async AddToOrder(xcsrf) {
    return await fetch(`/api/order/add?userId=${Auth.getPersonID}`, {
      method: "POST",
      headers: {
        "X-CSRF-TOKEN": xcsrf,
      },
    });
  }

  static async GetOrders() {
    return await fetch(`/api/order/all?userId=${Auth.getPersonID}`, {
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
