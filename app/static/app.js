document.addEventListener("DOMContentLoaded", function () {
  const userForm = document.getElementById("userForm");
  const usersList = document.getElementById("usersList");

  userForm.addEventListener("submit", function (e) {
    e.preventDefault();
    const userName = document.getElementById("userName").value;
    addUser(userName);
  });

  function addUser(name) {
    fetch("/api/users", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ name: name }),
    })
      .then((response) => response.json())
      .then((data) => {
        console.log("Success:", data);
        document.getElementById("userName").value = "";
        getUsers();
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  }

  function getUsers() {
    fetch("/api/users")
      .then((response) => response.json())
      .then((data) => {
        usersList.innerHTML = "";
        data.forEach((user) => {
          const li = document.createElement("li");
          li.textContent = `${user.id}: ${user.name}`;
          usersList.appendChild(li);
        });
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  }

  // Initial load of users
  getUsers();
});
