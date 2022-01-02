import React from 'react';
import './App.css';
import Navbar from './Components/Navbar';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Home from './pages';
import Top3 from './pages/top3';
import GraficaCircular from './pages/graficacircular';
import Ultimos5 from './pages/ultimos5';
import GraficaBarras from './pages/graficabarras';
import DashBoard from './pages/dashboard';

function App() {
  return (
    <Router>
      <Navbar />
      <Switch>
        <Route path='/' exact component={Home} />
        <Route path='/top3' component={Top3} />
        <Route path='/graficacircular' component={GraficaCircular} />
        <Route path='/ultimos5' component={Ultimos5} />
        <Route path='/graficabarras' component={GraficaBarras} />
        <Route path='/dashboard' component={DashBoard} />
      </Switch>
    </Router>
  );
}

export default App;
