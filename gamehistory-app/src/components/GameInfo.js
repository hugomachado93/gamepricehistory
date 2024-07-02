import styled from 'styled-components'
import React, { useEffect, useState } from "react";
import { useParams, useLocation } from 'react-router-dom';
import { Line } from "react-chartjs-2";
import { Chart as ChartJS } from 'chart.js/auto'

const Screen = styled.div`
display:flex;
justify-content: center;
`

const LineChartBox = styled.div`
width: 80%;
`

const GameInfo = () => {

    useEffect(() => {
    }, []);

    const { appid, name } = useParams();
    const { state } = useLocation()

    const label = [];
    const itens = [];

    state.GameDataHistory.forEach(element => {
        label.push(element.CreatedAt);
        itens.push(element.Price);
    });


    const data2 = {
        labels: label,
        datasets: [
            {
                label: "First dataset",
                data: itens,
                fill: true,
                backgroundColor: "rgba(75,192,192,0.2)",
                borderColor: "rgba(75,192,192,1)"
            }
        ]
    };

    console.log(data2);

    return (
        <div>
            <Screen>
                <LineChartBox>
                    <Line data={data2} />
                </LineChartBox>
            </Screen>
            <p>{appid} {name}</p>
        </div>
    )
}

export default GameInfo;