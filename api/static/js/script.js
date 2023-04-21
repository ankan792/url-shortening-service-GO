const form = document.getElementById('shorten-form');
const responseDiv = document.getElementById('response');
const shortUrl = document.getElementById('short-url');
const expiry = document.getElementById('expiry');
const rateLimit = document.getElementById('rate-limit');
const resetTime = document.getElementById('reset-time');
const errorMessage = document.getElementById('error-message');




form.addEventListener('submit', event => {
    event.preventDefault();

    // Get form data
    const url = form.url.value;
    const short_url = form.alias.value;

    // Send request to server
    fetch('/api/v1', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ url, short_url })
    })
    .then(response => response.json())
    .then(data => {
        if (data.error) {
            // Handle error
            errorMessage.textContent = data.error;
            resetTime.textContent = data.limit_reset;
            responseDiv.classList.add("hidden");
        } else {
            // Update response div with data
            shortUrl.href = data.short_url;
            shortUrl.textContent = data.short_url;
            expiry.textContent = data.expiry;
            rateLimit.textContent = data.rate_remaining;
            resetTime.textContent = data.rate_reset;

            // Clear error message
            errorMessage.textContent = '';
        }

        // Show response div
        responseDiv.style.display = 'block';
    });
});
