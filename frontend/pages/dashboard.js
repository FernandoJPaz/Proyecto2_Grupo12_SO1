import React from 'react';
import MongoData from '../Components/MongoData.js'
import RedisUltimos5Casos from '../Components/RedisUltimos5Casos.js'
import RedisGraficaRangoEdades from '../Components/RedisGraficaRangoEdades.js'


function DashBoard() {

    return (
      <div className="DashBoard">
        <MongoData />

        
      </div>
    );
    
  }

export default DashBoard;

