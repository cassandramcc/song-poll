import { useState, useEffect } from 'react';

export default function LoginButton(
    { 
        setIsLoggedIn,
        setUsername
    }
) {
  const [response, setResponse] = useState(null);
  const [error, setError] = useState(null);

  const handleClick = async () => {

    try {
      const res = await fetch("http://localhost:8080/login", {
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
    }
  };

  useEffect(() => {
    if (response && !error) {
      console.log(response)
      setIsLoggedIn(true);
      setUsername(response.username);
    }
  }, [response, error, setIsLoggedIn, setUsername]);

  return (
    <div>
      <button onClick={handleClick}>
        Login
      </button>
    </div>
  );
}