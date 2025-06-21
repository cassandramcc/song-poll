import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import reportWebVitals from './reportWebVitals';
import TextBox from './components/TextBox';
import TitleBar from './components/TitleBar';
import ButtonAPICall from './components/ButtonAPICall';
import Search from './components/Search';
import Artists from './components/model/Artists';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <TitleBar />
    <ButtonAPICall 
      buttonText='Login'
      endpoint='http://localhost:8080/login'
    />
    <ButtonAPICall 
      buttonText='Get Artists' 
      endpoint='http://localhost:8080/artists'
      ResponseComponent={Artists}
    />
    <Search />
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
