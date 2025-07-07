document.addEventListener("DOMContentLoaded", function () {
    console.log("login.js loaded ✅");

    const urlParams = new URLSearchParams(window.location.search);
    const msg = urlParams.get("msg");

    if (msg === "login_failed") {
        Swal.fire({
            icon: 'error',
            title: 'Login Failed',
            text: 'Invalid username or password.',
        });
    }
    
    if (msg === "register_success") {
        Swal.fire({
            icon: 'success',
            title: 'Registered Successfully',
            text: 'You can now login!',
        });
    }
    if (msg === "login_success") {
        Swal.fire({
            icon: 'success',
            title: 'Logged in Successfully',
            text: 'Welcome!',
        });
    }
    
    if (msg === "logout_success") {
        Swal.fire({
            icon: 'info',
            title: 'Logged Out',
            text: 'You have successfully logged out.',
        });
    }
    if (msg === "admin_not_approved") {
        Swal.fire({
            icon: 'error',
            title: 'Approved',
            text: 'Admin Not Approved by Super Admin.',
        });
    }
     if (msg === "register_failed") {
        Swal.fire({
            icon: 'error',
            title: 'Registration Failed',
            text: 'User already exists. Please try again.',
        });
    }
      if (msg) {
        const url = new URL(window.location.href);
        url.searchParams.delete("msg");
        window.history.replaceState({}, document.title, url.toString());
    }
});
