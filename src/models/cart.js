class Cart {
  static async getAllProducts() {
    return await fetch(`/api/cart/all`, {
      method: "GET",
    });
  }

  static async deleteFromCart(xcsrf, id) {
    return await fetch(`/api/cart/delete?id=${id}`, {
      method: "DELETE",
      headers: {
        "X-CSRF-TOKEN": xcsrf,
      },
    });
  }
}

export default Cart;
