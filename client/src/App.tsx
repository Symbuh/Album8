import React, { FC } from 'react';
import logo from './logo.svg';
import './App.css';
import Carousel from './Components/Carousel'
import UploadFile from './Components/UploadFile'

const App: FC = () => {

  return (
    <div id="App">
      <Carousel />
      <UploadFile />
    </div>
  );
}

export default App;
