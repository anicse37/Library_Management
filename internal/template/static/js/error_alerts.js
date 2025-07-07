document.addEventListener("DOMContentLoaded", function () {
    console.log("login.js loaded ✅");

    const urlParams = new URLSearchParams(window.location.search);
    const msg = urlParams.get("msg");

    if (msg === "unauthorized_access") {
        Swal.fire({
            icon: 'error',
            title: 'Unauthorized Access',
            text: 'Can not access this page ',
        });
    }
    if (msg === "return_error") {
        Swal.fire({
            icon: 'error',
            title: 'Can not return',
            text: 'Error Returning Book',
        });
    }

    if (msg === "error_in_borrowed_books") {
        Swal.fire({
            icon: 'error',
            title: 'Can get books',
            text: 'Error Getting Books',
        });
    }
    if (msg) {
        const url = new URL(window.location.href);
        url.searchParams.delete("msg");
        window.history.replaceState({}, document.title, url.toString());
    }
});
