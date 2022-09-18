import React from "react";
import {Navigate, Routes, Route, Router, Switch} from "react-router-dom";

//history
import { history } from './helpers/history';

import RouteGuard from "./components/RouteGuard";
import Main from "./pages/main";
import Auth from "./pages/auth";
import Login from "./pages/auth";

//pages


function RoutesR() {
    return (
        <div>
            <Router history={history} >
                <Switch>
                    <RouteGuard
                        exact
                        path="/"
                        component={Main}
                    ></RouteGuard>
                    <Route
                        exact
                        path="/auth"
                        component={Login}
                    ></Route>
                    {/*<Navigate to="/" replace={true}/>*/}
                </Switch>
            </Router>
        </div>

    );
}

export default RoutesR