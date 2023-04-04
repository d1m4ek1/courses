export const DOMToCurrency = () => {
  function toCurrency(price) {
    return new Intl.NumberFormat("ru-RU", {
      currency: "rub",
      style: "currency",
    }).format(price);
  }

  document.querySelectorAll(".price").forEach((node) => {
    node.textContent = toCurrency(node.textContent);
  });
};

export const DOMToDate = () => {
  function toDate(date) {
    return new Intl.DateTimeFormat("ru-RU", {
      day: "2-digit",
      month: "long",
      year: "numeric",
      hour: "2-digit",
      minute: "2-digit",
      second: "2-digit",
    }).format(new Date(date));
  }

  document.querySelectorAll(".date").forEach((node) => {
    node.textContent = toDate(node.textContent);
  });
};

// const $card = document.getElementById("card");

// if ($card) {
//   $card.addEventListener("click", (event) => {
//     if (event.target.classList.contains("js-remove")) {
//       const { id } = event.target.dataset;

//       fetch("/card/remove/" + id, {
//         method: "DELETE",
//       })
//         .then((res) => res.json())
//         .then((card) => {
//           if (card.courses.length) {
//             const html = card.courses
//               .map((c) => {
//                 return `
//               <tr>
//                 <td><a href="/courses/${c.id}" target="_blank">${c.title}</a></td>
//                 <td class="price">${toCurrency(c.price)}</td>
//                 <td>${c.count}</td>
//                 <td>
//                   <button class="btn btn-small js-remove" data-id="${c.id}">Удалить</button>
//                 </td>
//               </tr>`;
//               })
//               .join("");

//             $card.querySelector("tbody").innerHTML = html;
//             $card.querySelector("#total-price").textContent = toCurrency(card.totalPrice);
//           } else {
//             $card.innerHTML = `<p>Корзина пуста</p>`;
//           }
//         });
//     }
//   });
// }
