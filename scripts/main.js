function logout() {
    $.ajax({
        type: "POST",
        url: "/users/logout",
        success: function(response) {
            window.location.reload()
        },
        error: function(xhr, status, error) {
            console.log(xhr, status, error)
            toast(status)
        }
    });
}

function toast(text) {
    Toastify({
        text: text,
        className: "info",
        style: {
          background: "linear-gradient(to right, #00b09b, #96c93d)",
        }
      }).showToast();
}