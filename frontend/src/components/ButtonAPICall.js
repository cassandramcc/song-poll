import React, { useState } from 'react';

export default function ButtonAPICall(
    { 
        buttonText = "Default Button Text", 
        endpoint = "no endpoint",
        query = '',
        ResponseComponent = null
    }
) {
  const [isClicked, setIsClicked] = useState(false);
  const [response, setResponse] = useState(null);
  const [error, setError] = useState(null);

  const handleClick = async () => {

    if (endpoint == "no endpoint") {
      throw new Error('Please add an endpoint for this button')
    }
    console.log("Button "+ buttonText + " clicked")

    try {
      let url = endpoint
      if (query != '') {
        url = url + "?query="+query
      }
      const res = await fetch(url, {
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

  return (
    <div>
      <button onClick={handleClick}>
        {buttonText}
      </button>

      {response && !error && ResponseComponent && (
        <ResponseComponent data={response} />
      )}
    </div>
  );
}