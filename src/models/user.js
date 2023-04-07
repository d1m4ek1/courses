class User {
  static async addCourse(data) {
    data.addDate = Date();

    return await fetch("/api/course/add", {
      method: "POST",
      body: JSON.stringify(data),
      headers: {
        "Content-Type": "text/json",
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

  static async editCourse(id, data) {
    return await fetch(`/api/course/edit?id=${id}`, {
      method: "PUT",
      body: JSON.stringify(data),
    });
  }

  static async deleteCourse(id) {
    return await fetch(`/api/course/delete?id=${id}`, {
      method: "DELETE",
    });
  }

  static async AddCourseToCart(data) {
    return await fetch(`/api/course/buy`, {
      method: "POST",
      body: JSON.stringify(data),
    });
  }
}

export default User;
