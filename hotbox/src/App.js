import './App.css';
import Temperature from "./components/Temperature"

function App() {
  return (
    <div>
      <Temperature
        desiredScale='f'
        givenScale='c'
		/>
    </div>
  );
}

export default App;
