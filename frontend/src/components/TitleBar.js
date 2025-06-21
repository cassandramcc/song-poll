import { useState } from 'react';
import LoginButton from './LoginButton';
import './TitleBar.css'

export default function TitleBar () {
    const [isLoggedIn, setIsLoggedIn] = useState(false)
    const [username, setUsername] = useState('')
    return (
       <div className="title-bar">
            <h1>Song Poll!</h1>
            <div className="login-section">
                {isLoggedIn && (
                    <p>Logged in as: {username}</p>
                )}
                <LoginButton setIsLoggedIn={setIsLoggedIn} setUsername={setUsername} />
            </div>
        </div>
    );
}