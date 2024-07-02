// import React, { Fragment, useEffect, useState } from "react";
import "./index.css"
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Game from "./components/games";
import GameInfo from "./components/GameInfo";


export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Game />} />
        <Route path="/info/game/:appid/:name" element={<GameInfo/>}/>
      </Routes>
    </BrowserRouter>
  );
}