class Auth {
  static async Login(xcsrf, data) {
    return await fetch("/api/auth/login", {
      method: "POST",
      body: JSON.stringify(data),
      headers: {
        "Content-Type": "application/json",
        "X-CSRF-TOKEN": xcsrf,
      },
    });
  }

  static async Register(xcsrf, data) {
    return await fetch("/api/auth/register", {
      method: "POST",
      body: JSON.stringify(data),
      headers: {
        "Content-Type": "application/json",
        "X-CSRF-TOKEN": xcsrf,
      },
    });
  }

  static async Logout(xcsrf) {
    return await fetch("/api/auth/logout", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "X-CSRF-TOKEN": xcsrf,
      },
    });
  }

  static async CheckUserAuth() {
    return await fetch("/api/auth/check", {
      method: "GET",
    });
  }
}

export default Auth;
