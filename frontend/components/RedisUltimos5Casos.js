import '../App.css';
import React, { useState, useEffect } from "react";
import axios from "axios";
//const URL_TEST = "https://us-central1-proyecto2-g8-s01.cloudfunctions.net/test";
const URL_F1 = "http://localhost:5000/";

function  RedisUltimos5Casos() {


  const [lastfive_data, setLastFive] = useState([]);
  useEffect(() => {
    setInterval(function(){ 
        axios.get(URL_F1)
            .then(function(response) {
             let datos = response.data.replaceAll("'","\"");
             let res = datos.split(";");
      
            let arreglo = [];

            for (let index = 0; index < 3; index++) {
                arreglo.push(JSON.parse(res[index]));
            }

            setLastFive(arreglo);
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
      <h1 id='title'>Ultimos 5 casos</h1>

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
        {lastfive_data.map(function(item, key){
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

export default RedisUltimos5Casos;
