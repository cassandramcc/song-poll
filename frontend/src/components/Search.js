import React, { useState } from 'react';
import TextBox from './TextBox';
import ButtonAPICall from './ButtonAPICall';
import Artists from './model/Artists';

export default function Search() {
  const [currentText, setCurrentText] = useState('');

  return (
    <div>
      <TextBox onTextChange={setCurrentText} />
      <ButtonAPICall buttonText="Search" endpoint="http://localhost:8080/spotify/artists" query={currentText} ResponseComponent={Artists}/>
    </div>
  );
}