import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import MainHeader from './components/MainHeader'
import MainFooter from './components/MainFooter'
import { BrowserRouter } from 'react-router-dom'
import Navigation from './pages/Navigation'

function App() {
  return (
    <>
      <BrowserRouter>
        <div className="App">
          <MainHeader />
          <Navigation />
          <MainFooter />
        </div>
      </BrowserRouter>
    </>
  );
}

export default App;
