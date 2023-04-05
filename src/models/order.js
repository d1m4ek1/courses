import Auth from "./auth";

class Order {
  static async AddToOrder() {
    return await fetch(`/api/order/add?userId=${Auth.getPersonID}`, {
      method: "POST",
    });
  }

  static async GetOrders() {
    return await fetch(`/api/order/all?userId=${Auth.getPersonID}`, {
      method: "GET",
    });
  }

  static async DeleteOrder(orderId) {
    return await fetch(`/api/order/delete?orderId=${orderId}`, {
      method: "DELETE",
    });
  }
}

export default Order;
