import React, { useState } from 'react';

export default function BasicButton({ buttonText = "Click Me!" }) {
  const [isClicked, setIsClicked] = useState(false);
  const [response, setResponse] = useState(null);
  const [error, setError] = useState(null);

  const handleClick = async () => {
    console.log("button clicked")

    try {
      // Call your Go backend endpoint
      const res = await fetch('http://localhost:8090/artists', {
        method: 'GET',
      });

      if (!res.ok) {
        throw new Error(`HTTP error! status: ${res.status}`);
      }

      const data = await res.json();
      setResponse(data);
      
    } catch (err) {
      setError(err.message);
      console.error('Error calling backend:', err);
    } finally {
      // Reset the clicked state after animation
      setTimeout(() => setIsClicked(false), 150);
    }
  };

  return (
    <div>
      <div>
        <button
          onClick={handleClick}
        >
          {buttonText}
        </button>

        {response && !error && (
            <div>
              <p>
                Artists:
              </p>
              <pre>
                {JSON.stringify(response, null, 2)}
              </pre>
            </div>
          )}
      </div>
    </div>
  );
}