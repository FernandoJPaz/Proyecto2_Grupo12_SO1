import React, { useState, useEffect } from "react";
import socketIOClient from "socket.io-client";
import Chart from "react-google-charts";
const ENDPOINT = "http://localhost:5000/";
//const ENDPOINT = "https://mongo-de6gmjxwaa-uc.a.run.app/"

function  MongoGraficaCircular() {

  const [pie_graph_data, setPieGraphData] = useState([]);
    useEffect(() => {
      const socket = socketIOClient(ENDPOINT);
      socket.on("FromAPI", data => {

        //Agrego la info para la grafica de pie 
        let arreglo=[['Departamento', 'Vacunados']];
        for (const iterator of data) {
          arreglo.push([iterator.location, iterator.cantidad]);
        }
        setPieGraphData(arreglo);

        
      });
    }, []);

  
    return (
    <div className="App">
     
      <div align="center" className="col-md-4">
          <Chart
            width={'1000px'}
            height={'1000px'}
            chartType="PieChart"
            loader={<div>Loading Chart</div>}
            data={pie_graph_data?pie_graph_data:[['Departamento','Vacunados'],['','100']]}
            options={{
              title: 'Porcentaje de infectados por departamento',
            }}
            //rootProps={{ 'data-testid': '1' }}
          />
      </div>
    </div>
    );

}

export default MongoGraficaCircular;
