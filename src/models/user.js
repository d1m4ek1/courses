class User {
  static async addCourse(xcsrf, data) {
    data.addDate = Date();

    return await fetch("/api/course/add", {
      method: "POST",
      body: JSON.stringify(data),
      headers: {
        "Content-Type": "application/json",
        "X-CSRF-TOKEN": xcsrf,
      },
    });
  }

  static async getCourses() {
    return await fetch("/api/course/all", {
      method: "GET",
    });
  }

  static async getCourseById(id) {
    return await fetch(`/api/course/get?id=${id}`, {
      method: "GET",
    });
  }

  static async editCourse(xcsrf, id, data) {
    return await fetch(`/api/course/edit?id=${id}`, {
      method: "PUT",
      body: JSON.stringify(data),
      headers: {
        "Content-Type": "application/json",
        "X-CSRF-TOKEN": xcsrf,
      },
    });
  }

  static async deleteCourse(xcsrf, id) {
    return await fetch(`/api/course/delete?id=${id}`, {
      method: "DELETE",
      headers: {
        "X-CSRF-TOKEN": xcsrf,
      },
    });
  }

  static async AddCourseToCart(xcsrf, data) {
    return await fetch(`/api/course/buy`, {
      method: "POST",
      body: JSON.stringify(data),
      headers: {
        "X-CSRF-TOKEN": xcsrf,
      },
    });
  }
}

export default User;
