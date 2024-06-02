document.addEventListener("DOMContentLoaded", () => {
    const loginForm = document.getElementById('login-form');

    loginForm.addEventListener("submit", async (event) => {
        event.preventDefault(); // Prevent form from submitting the default way
        
        const email = document.getElementById('email').value;
        const password = document.getElementById('password').value;
        
        try {
            const apiRes = await fetch('http://localhost:8080/auth/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Accept': 'application/json'
                },
                body: JSON.stringify({ email, password })
            });
            
            if (!apiRes.ok) {
                throw new Error(`HTTP error! status: ${apiRes.status}`);
            }
            
            const content = await apiRes.json();
            console.log(content);
        } catch (error) {
            console.error('Error:', error);
        }
    });
});
