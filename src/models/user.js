class User {
  static async addCourse(data) {
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
    return await fetch(`/api/course/byid?id=${id}`, {
      method: "GET",
    });
  }
}

export default User;
