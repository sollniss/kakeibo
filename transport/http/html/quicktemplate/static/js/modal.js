const modalOpenButtons = document.getElementsByClassName("modal__open");

for(let i = 0; i < modalOpenButtons.length; i++) {
  // get modal id
  let modal = document.getElementById(modalOpenButtons[i].dataset.modal);

  // register click event for open button
  modalOpenButtons[i].addEventListener("click", openModal.bind(null, modal));

  // register click event for all close buttons inside the modal
  let closeButtons = modal.getElementsByClassName("modal__close");
  for(let j = 0; j < closeButtons.length; j++) {
    closeButtons[j].addEventListener("click", closeModal.bind(null, modal));
  }

  switch(modalOpenButtons[i].dataset.modal) {
    case "add":
      modalOpenButtons[i].addEventListener("click", initAddModal.bind(modal, modalOpenButtons[i].dataset));
      break;
    case "edit":
      modalOpenButtons[i].addEventListener("click", initEditModal.bind(modal, modalOpenButtons[i].dataset));
      break;
    case "delete":
      modalOpenButtons[i].addEventListener("click", initDeleteModal.bind(modal, modalOpenButtons[i].dataset));
      break;
  }
}

// register modal submit buttons
document.getElementById("add-submit").addEventListener("click", submitAdd);
document.getElementById("edit-submit").addEventListener("click", submitEdit);
document.getElementById("delete-submit").addEventListener("click", submitDelete);

// register modal close on backgroud click
const modals = document.getElementsByClassName("modal");
for(let i = 0; i < modals.length; i++) {
  modals[i].addEventListener("click", function(e) {
    if(e.target == this) {
      closeModal(this);
    }
  });
}

function openModal(modal) {
  modal.style.display = 'unset';
}

function closeModal(modal) {
  modal.style.display = 'none';
}

function initDeleteModal(data) {
  document.getElementById("delete-submit").dataset.target = data.target;
}

function submitDelete() {
  this.disabled = true;
  const url = this.dataset.target;
  console.log(url);

  fetch(url, {
    method: "DELETE"
  }).then(response => response.json()).then(result => {
    console.log(result);
    if(result.error_code === 0) {
      let baseUrl = location.href.split("#")[0];
      location.href = baseUrl;
      location.reload();
    }
    this.disabled = false;
  }).catch((error) => {
    console.error('Error:', error);
  });
}

function initEditModal(data) {
  document.getElementById("edit-date").value = data.date;
  document.getElementById("edit-amount").value = data.amount;
  document.getElementById("edit-type").value = data.type;
  document.getElementById("edit-person").value = data.person;
  document.getElementById("edit-comment").value = data.comment;

  document.getElementById("edit-submit").dataset.target = data.target;
}

function submitEdit() {
  this.disabled = true;

  const date = document.getElementById("edit-date");
  const amount = document.getElementById("edit-amount");
  const type_id = document.getElementById("edit-type");
  const person_id = document.getElementById("edit-person");
  const comment = document.getElementById("edit-comment");

  const updatedEntry = {
    date: date.value,
    amount: parseInt(amount.value),
    type_id: parseInt(type_id.value),
    person_id: parseInt(person_id.value),
    comment: comment.value
  }

  const url = this.dataset.target;

  fetch(url, {
    method: "PUT", 
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(updatedEntry)
  }).then(response => response.json()).then(result => {

    if(result.error_code === 0) {
      const baseUrl = location.href.split("#")[0];
      location.href = baseUrl + "#" + result.dom_id;
      location.reload();
    }

    this.disabled = false;
  }).catch((error) => {
    console.error('Error:', error);
  });
}

function initAddModal(data) {
  const d = document.getElementById("add-date");
  d.value = d.dataset.default;
  
  document.getElementById("add-amount").value = "";
  document.getElementById("add-type").selectedIndex = 8; // 遊興費
  document.getElementById("add-person").selectedIndex = 0;
  document.getElementById("add-comment").value = "";

  document.getElementById("add-submit").dataset.target = data.target;
}

function submitAdd() {
  this.disabled = true;

  const date = document.getElementById("add-date");
  const amount = document.getElementById("add-amount");
  const type_id = document.getElementById("add-type");
  const person_id = document.getElementById("add-person");
  const comment = document.getElementById("add-comment");

  let newEntry = {
    date: date.value,
    amount: parseInt(amount.value),
    type_id: parseInt(type_id.value),
    person_id: parseInt(person_id.value),
    comment: comment.value
  }

  const url = this.dataset.target;

  fetch(url, {
    method: "POST", 
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(newEntry)
  }).then(response => response.json()).then(result => {
    console.log(result);

    if(result.error_code === 0) {
      const baseUrl = location.href.split("#")[0];
      location.href = baseUrl + "#" + result.dom_id;
      location.reload();
    }

    this.disabled = false;
  }).catch((error) => {
    console.error('Error:', error);
  });
}