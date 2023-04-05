import Auth from "./auth";

class Cart {
  static async getAllProducts() {
    return await fetch(`/api/cart/all?userId=${Auth.getPersonID}`, {
      method: "GET",
    });
  }

  static async deleteFromCart(id) {
    return await fetch(`/api/cart/delete?id=${id}`, {
      method: "DELETE",
    });
  }
}

export default Cart;