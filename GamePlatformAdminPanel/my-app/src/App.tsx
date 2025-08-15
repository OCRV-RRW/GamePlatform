import './App.css';
import AppRouter from './AppRouter';
import ControllerStatusCodePage from './StatusCodePage/ControllerStatusCodePage';


export default function App() {
  return (
    <div className="App">
      <ControllerStatusCodePage>
        <AppRouter />
      </ControllerStatusCodePage>
    </div>
  );
}