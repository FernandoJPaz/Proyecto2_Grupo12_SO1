import React, { useState, useEffect } from "react";
import socketIOClient from "socket.io-client";
import Chart from "react-google-charts";
import axios from "axios";

const ENDPOINT = "http://localhost:5000";
//const ENDPOINT = "https://mongo-de6gmjxwaa-uc.a.run.app/"

function  MongoData() {

  const [all_data, setAll] = useState([]);
  useEffect(() => {
    setInterval(function(){ 
        axios.get(ENDPOINT)
            .then(function(response) {
             let datos = response.data.replaceAll("'","\"");
             let res = datos.split(";");
      
            let arreglo = [];
            for (const iterator of res) {
                arreglo.push(JSON.parse(iterator));
            }

            setAll(arreglo);
            })
            .catch(function(error) {
                console.log(error);
            })
            .then(function() {}); 
    }, 3000); 
  }, []);

  
    return (
    <div className="App">
      <div className="col-md-4">
      <h1 id='title'>Todos los Casos</h1>

      <div align="center" className="container-fluid">
      <table striped bordered hover variant="dark">
        <thead>
          <tr>
           <th>Name</th>
           <th>Location</th>
           <th>Age</th>
           <th>Vaccine Type</th>
           <th>N Dose</th>
          </tr>
        </thead>
        <tbody>
        {all_data.map(function(item, key){
            return(
              <tr key={key}>
                 <td>{item.name}</td>
                 <td>{item.location}</td>
                 <td>{item.age}</td>
                 <td>{item.vaccine_type}</td>
                 <td>{item.n_dose}</td>
              </tr>
            )
        })}
        </tbody>
      </table>
      </div>
     
      </div>
    
      
    </div>
    );

}

export default MongoData;
