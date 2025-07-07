document.addEventListener("DOMContentLoaded", function () {
    console.log("register.js loaded ✅");

    const urlParams = new URLSearchParams(window.location.search);
    const msg = urlParams.get("msg");

    if (msg === "register_failed") {
        Swal.fire({
            icon: 'error',
            title: 'Registration Failed',
            text: 'User already exists. Please try again.',
        });
    }
});
