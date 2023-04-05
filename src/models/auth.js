function generatePersonID() {
  var pass = "";
  var str = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_-=";

  for (let i = 1; i <= 128; i++) {
    var char = Math.floor(Math.random() * str.length + 1);

    pass += str.charAt(char);
  }

  return pass;
}

class Auth {
  static setNewPersonID() {
    if (!localStorage.getItem("person_id")) {
      const personId = generatePersonID();
      localStorage.setItem("person_id", personId);
      return personId;
    } else {
      return localStorage.getItem("person_id");
    }
  }

  static get getAnswerAllow() {
    return localStorage.getItem("person_id") ? true : false;
  }

  static get getPersonID() {
    return localStorage.getItem("person_id");
  }
}

export default Auth;
