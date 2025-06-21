import { useState } from 'react';

export default function TextBox({ onTextChange }) {
  const [text, setText] = useState('');

  const handleChange = (e) => {
    const newText = e.target.value
    setText(newText)
    onTextChange(newText)
  }

  const handleKeyPress = (e) => {
    if (e.key === 'Enter') {
      console.log(text)
      onTextChange(text)
      setText('')
    }
  };

  return (
    <div>
      <input
        type="text"
        placeholder="Type here..."
        value={text}
        onChange={handleChange}
        onKeyPress={handleKeyPress}
      />
    </div>
  );
}