import { render } from "react-dom";
import {
    BrowserRouter,
    Route,
    Routes
} from "react-router-dom";
import App from "./App";
import Main from "./pages/main";
import Auth from "./pages/auth";


const rootElement = document.getElementById("root");
render(
    <App/>,
    rootElement
);