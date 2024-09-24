import React, { useState } from 'react';

function App() {
  const [url, setUrl] = useState('');
  const [shortUrl, setShortUrl] = useState('');
  const [error, setError] = useState('');

  // Handle URL input change
  const handleUrlChange = (e) => {
    setUrl(e.target.value);
  };

  // Simulate URL shortening process
  const handleShortenUrl = () => {
    if (isValidUrl(url)) {
      const shortened = `https://short.ly/${Math.random().toString(36).substring(7)}`;
      setShortUrl(shortened);
      setError('');
    } else {
      setError('Please enter a valid URL.');
    }
  };

  // URL validation function
  const isValidUrl = (string) => {
    try {
      new URL(string);
      return true;
    } catch (_) {
      return false;
    }
  };

  return (
    <div style={styles.container}>
      <header style={styles.header}>
        <h1>URL Shortener</h1>
        <p>Enter a long URL to shorten it.</p>
      </header>

      <div style={styles.formContainer}>
        <input
          type="text"
          value={url}
          onChange={handleUrlChange}
          placeholder="Enter your URL here"
          style={styles.input}
        />
        <button onClick={handleShortenUrl} style={styles.button}>
          Shorten URL
        </button>
      </div>

      {error && <p style={styles.error}>{error}</p>}

      {shortUrl && (
        <div style={styles.resultContainer}>
          <p>Original URL: {url}</p>
          <p>
            Shortened URL: <a href={shortUrl}>{shortUrl}</a>
          </p>
        </div>
      )}
    </div>
  );
}

// Inline CSS styles
const styles = {
  container: {
    textAlign: 'center',
    padding: '50px',
    fontFamily: 'Arial, sans-serif',
  },
  header: {
    marginBottom: '20px',
  },
  formContainer: {
    marginTop: '20px',
  },
  input: {
    padding: '10px',
    width: '300px',
    border: '1px solid #ccc',
    borderRadius: '4px',
    fontSize: '16px',
  },
  button: {
    padding: '10px 20px',
    marginLeft: '10px',
    backgroundColor: '#4CAF50',
    color: 'white',
    border: 'none',
    cursor: 'pointer',
    borderRadius: '4px',
    fontSize: '16px',
  },
  buttonHover: {
    backgroundColor: '#45a049',
  },
  resultContainer: {
    marginTop: '30px',
  },
  error: {
    color: 'red',
  },
};

export default App;
