const loginForm = document.getElementById('login-form')

loginForm.addEventListener("submit", async (event) => {
    event.preventDefault(); // Prevent form from submitting the default way
    console.log('test');
    
    try {
        const apiRes = await fetch('http://192.168.1.191:8080/auth/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Accept': '*/*'
            },
            body: JSON.stringify({
                email: document.getElementById('email').value,
                password: document.getElementById('password').value
            })
        });
        
        if (!apiRes.ok) {
            throw new Error(`HTTP error! status: ${apiRes.status}`);
        }
        
        const content = await apiRes.json();
        console.log(content);
    } catch (error) {
        console.error('Error:', error);
    }
})