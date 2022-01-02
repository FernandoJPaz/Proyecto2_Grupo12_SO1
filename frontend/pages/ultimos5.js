import React from 'react';
import Container from 'react-bootstrap/Container';
import RedisUltimos5Casos from '../Components/RedisUltimos5Casos.js'

const Ultimos5 = () => {
  return (
    <Container align="center" className="p-3">
    <RedisUltimos5Casos/>
  </Container>
  );
};

export default Ultimos5;
