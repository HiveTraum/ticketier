import * as React from 'react';
import {FunctionComponent} from 'react';
import * as ReactDOM from 'react-dom';

import {BrowserRouter as Router, Route, Switch} from "react-router-dom";
import Navigation from "./components/Navigation";
import {KeycloakConfig} from 'keycloak-js'
import UserProvider from "./storages/UserStorage";

const config: KeycloakConfig = {
    clientId: "frontend",
    realm: "clients",
    url: "http://localhost:8080/auth",
}


const App: FunctionComponent = () => <div>

    <Router>
        <UserProvider config={config}>
            <Navigation/>
        </UserProvider>
        <Switch>
            <Route path="/admin">
                <h1>Admin</h1>
            </Route>
            <Route path="/">
                <h1>Kiosk</h1>
            </Route>
        </Switch>
    </Router>

</div>;


ReactDOM.render(<App/>, document.getElementById('root'));