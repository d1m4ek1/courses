class Auth {
  static async Login(data) {
    return await fetch("/api/auth/login", {
      method: "POST",
      body: JSON.stringify(data),
      headers: {
        "Content-Type": "text/json",
      },
    });
  }

  static async Register(data) {
    return await fetch("/api/auth/register", {
      method: "POST",
      body: JSON.stringify(data),
      headers: {
        "Content-Type": "text/json",
      },
    });
  }

  static async Logout() {
    return await fetch("/api/auth/logout", {
      method: "POST",
    });
  }

  static async CheckUserAuth() {
    return await fetch("/api/auth/check", {
      method: "GET",
    });
  }
}

export default Auth;
