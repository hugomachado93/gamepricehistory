import styled from 'styled-components'
import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from 'axios';
import SteamLogoImg from '../images/Steam_Logo.png'
import { click } from '@testing-library/user-event/dist/click';

const Screen = styled.div`
width:100%;
height:100%;
overflow:hidden;
background-color: black;
`

const Image = styled.img`
object-fit: cover;
background-position: center center;
background-repeat: no-repeat;
background-size:contain;
margin: 10px 0 10px 0;
max-width: 100%;
height: auto;
`

const Box = styled.div`
display:grid;
grid-template-columns: repeat(auto-fit, minmax(min(450px, 100%), 1fr));
justify-content: space-evenly;
margin: 20px;
grid-gap: 20px;
`

const GameBox = styled.div`
color: white;
display:flex;
flex-direction:column;
cursor: pointer;
`
const GameTitle = styled.div`

`

const PriceContainer = styled.div`
display:flex;
align-items: center;
font-size:30px;
@media (max-width: 500px) {
  font-size: 20px;
}
`

const Price = styled.div`
width:200px;
`

const SteamLogo = styled.img`
background-position: center center;
background-repeat: no-repeat;
background-size:contain;
width:5%;
height:auto;
margin-right: 10px;
`

const Game = () => {
  const navigate = useNavigate();

  useEffect(() => {
    fetchData()
  }, [])

  const [gameInfo, setGameInfo] = useState([]);


  const fetchData = async () => {
    const result = await axios.get("http://localhost:8080/api/game/paginated?page=0&size=10")
    console.log(result)
    setGameInfo(result.data)
  }

  const clickHere = (gameInfo) => {
    navigate(`/info/game/${gameInfo.AppId}/${gameInfo.Name}`, { state: gameInfo })
  }
  return (
    <div>
      <Screen>
        <Box>
          {gameInfo.map(x => <GameBox onClick={() => clickHere(x)}><Image src={x.CoverUrl}/><a>{x.Name}</a><PriceContainer><SteamLogo src={SteamLogoImg}/><Price>{x.Price == 0? 'FREE': 'R$ ' + x.Price}</Price></PriceContainer></GameBox>)}
        </Box>
      </Screen>
    </div>
  );
};

export default Game;