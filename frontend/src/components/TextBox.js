import { useState } from 'react';

export default function TextBox() {
  const [text, setText] = useState('');

  const handleKeyPress = (e) => {
    if (e.key === 'Enter') {
      console.log(text);
      setText('')
    }
  };

  return (
    <div>
      <input
        type="text"
        placeholder="Type here..."
        value={text}
        onChange={(e) => setText(e.target.value)}
        onKeyPress={handleKeyPress}
      />
    </div>
  );
}