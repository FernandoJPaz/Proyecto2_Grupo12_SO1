import React, { useState, useEffect } from "react";
import socketIOClient from "socket.io-client";
import '../App.css';
const ENDPOINT = "http://localhost:5000/";

function  MongoTop3() {
  
  const [top3_data, setTop3] = useState([]);
    useEffect(() => {
      const socket = socketIOClient(ENDPOINT);
      socket.on("FromAPI", data => {

        //Agrego los top3 al arreglo
        if(data.length>3){
          let arr=[];
          for (let index = 0; index < 3; index++) {
              const element = data[index];
              arr.push(element);
          }
          setTop3(arr);
        }else{
          setTop3(data);
        }
      });
    }, []);

    return (
      <div className="App">
        <div className="col-md-4">
        <h1 id='title'>Top 3 - departamentos con mas casos</h1>
  
        <div align="center" className="container-fluid">
        <table striped bordered hover variant="dark">
          <thead>
            <tr>
             <th>Location</th>
             <th>Cantidad</th>
            </tr>
          </thead>
          <tbody>
          {top3_data.map(function(item, key){
              return(
                <tr key={key}>
                   <td>{item.location}</td>
                   <td>{item.cantidad}</td>
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

export default MongoTop3;
