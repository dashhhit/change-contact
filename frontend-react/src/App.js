import RoutesR from "./routesR";
import {setAuthToken} from "./helpers/setAuthToken";



function App() {
    const token = localStorage.getItem("token")
    if (token) {
        setAuthToken(token);
    }
  return (
    <div>
      <RoutesR/>
    </div>
  );
}

export default App;
