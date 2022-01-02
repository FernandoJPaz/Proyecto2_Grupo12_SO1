import React from 'react';
import Container from 'react-bootstrap/Container';
import RedisGraficaRangoEdades from '../Components/RedisGraficaRangoEdades.js'

const GraficaBarras = () => {
  return (
    <Container align="center" className="p-3">
    <RedisGraficaRangoEdades/>
  </Container>
  );
};

export default GraficaBarras;
