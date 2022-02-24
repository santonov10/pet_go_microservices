import './App.css';
import Router from "./shared/Router";
import routes from "./routes";
import UserProvider from "./providers/UserProvider";

function App() {
  return (
    <div className="App">
        <UserProvider>
            <Router routes={routes}/>
        </UserProvider>
    </div>
  );
}

export default App;
