import React, { useState, useEffect } from "react";
import axios from "axios";
import Chart from "react-google-charts";
//const URL_TEST = "https://us-central1-proyecto2-g8-s01.cloudfunctions.net/test";
const URL_F2 = "http://localhost:5000/";

function  RedisGraficaRangoEdades() {

  const [barras_graph_data, setBarrasGraphData] = useState([]);

  useEffect(() => {
    setInterval(function(){ 
        axios.get(URL_F2)
            .then(function(response) {
            //Agrego la info para la grafica de pie 
            let arreglo=[['Rango de edad', 'Vacunados']]; 
            arreglo.push(['0 - 20 años', response.data.rango1]);
            arreglo.push(['21 - 40 años', response.data.rango2]);
            arreglo.push(['41 - 60 años', response.data.rango3]);
            arreglo.push(['61 - 80 años', response.data.rango4]);
            arreglo.push(['81 - 100 años', response.data.rango5]);
            arreglo.push(['mas de 100 años', response.data.rango6]);
            
            setBarrasGraphData(arreglo);

            console.log(response.data);
            
            })
            .catch(function(error) {
                console.log("HAY ERROR "+error);
            })
            .then(function() {}); 
    }, 3000); 
  }, []);

  
    return (
    <div className="App">
      <div className="col-md-4">   
      </div>
      <div align="center" className="col-md-4">
          
          <Chart
            width={'1000px'}
            height={'1000px'}
            chartType="BarChart"
            loader={<div>Loading Chart</div>}
            data={barras_graph_data?barras_graph_data:[['Rango de edad', 'Infectados'],['','100']]}
            options={{
              title: 'Rango de edad de infectados en Redis',
              chartArea: { width: '60%' },
              hAxis: {
                title: 'Total Population',
                minValue: 0,
              },
              vAxis: {
                title: 'Rango de Edad',
              },
            }}
            // For tests
            //rootProps={{ 'data-testid': '1' }}
          />
      </div>
      
    </div>
    );

}

export default RedisGraficaRangoEdades;
