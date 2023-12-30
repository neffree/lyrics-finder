async function searchLyrics() {
    const searchInput = document.getElementById('searchInput').value;
    const artistInput = document.getElementById('artistInput').value;
    const apiUrl = 'https://www.stands4.com/services/v2/lyrics.php';
    const uid = '[insert UID]'; // Replace with  UID
    const tokenid = '[insert token ID]'; // Replace with Token ID
    const format = 'json'; // Format can be xml or json, here we use json

    try {
        const response = await fetch(`http://localhost:8080/searchLyrics?term=${encodeURIComponent(searchInput)}&artist=${encodeURIComponent(artistInput)}`);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();
        displayResults(data);
        console.log(response);
    } catch (error) {
        console.error('Error fetching data: ', error);
    }
}

function displayResults(data) {
    const resultsDiv = document.getElementById('results');
    resultsDiv.innerHTML = '';

    // Check if the 'result' array exists and has elements
    if (data.result && data.result.length > 0) {
        // Loop through each result and append its content to the resultsDiv
        data.result.forEach(result => {
            resultsDiv.innerHTML += `
                <div class="result-item">
                    <p><strong>Song:</strong> ${result.song} (<a href="${result['song-link']}" target="_blank">Lyrics</a>)</p>
                    <p><strong>Artist:</strong> ${result.artist} (<a href="${result['artist-link']}" target="_blank">Artist Info</a>)</p>
                    <p><strong>Album:</strong> ${result.album} (<a href="${result['album-link']}" target="_blank">Album Info</a>)</p>
                </div>
            `;
        });
    } else {
        resultsDiv.innerHTML = '<p>No results found.</p>';
    }
}

