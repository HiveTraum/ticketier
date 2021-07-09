import Keycloak, {KeycloakConfig, KeycloakInstance} from 'keycloak-js';
import * as React from "react";
import {createContext, FunctionComponent, useEffect, useState} from "react";

interface User {
    id: string,
    firstName: string,
    lastName: string,
    fullName: string
}

const getUserFromKeycloak = (keycloak: KeycloakInstance): User | undefined => {
    if (!keycloak) return;
    if (!keycloak.tokenParsed) return;

    // @ts-ignore
    const givenName = keycloak.tokenParsed['given_name']
    // @ts-ignore
    const familyName = keycloak.tokenParsed['family_name']

    const firstName = givenName && typeof givenName == "string" ? givenName : ""
    const lastName = familyName && typeof familyName == "string" ? familyName : ""
    const fullName = `${lastName} ${firstName}`
    const id = keycloak.tokenParsed.sub || "";

    return {id, firstName, lastName, fullName}
}

interface KeycloakState {
    authenticated: boolean,
    user?: User,
    keycloak?: KeycloakInstance
    initialized: boolean,
}

const useKeycloak = (config: KeycloakConfig): KeycloakState => {
    const [state, setState] = useState<KeycloakState>({authenticated: false, initialized: false})
    useEffect(() => {
        const keycloak: KeycloakInstance = Keycloak(config);
        keycloak
            .init({
                onLoad: "check-sso",
                checkLoginIframe: false,
                silentCheckSsoFallback: true,
            })
            .then(authenticated => setState({
                ...state, authenticated, keycloak,
                initialized: true,
                user: getUserFromKeycloak(keycloak)
            })).catch(console.log);
    }, []);

    return state
}

interface UserContextState {
    user?: User,
    signOut: () => void,
    signIn: () => void,
    initialized: boolean,
    authenticated: boolean
}

export const UserContext = createContext<UserContextState>({
    signOut: () => undefined,
    signIn: () => undefined,
    initialized: false,
    authenticated: false
});

interface Props {
    config: KeycloakConfig
}

const UserProvider: FunctionComponent<Props> = ({children, config}) => {
    const {authenticated, user, initialized, keycloak} = useKeycloak(config);

    const signOut = () => keycloak?.logout()
    const signIn = () => keycloak?.login()

    return <UserContext.Provider value={{user, signOut, signIn, initialized, authenticated}}>
        {children}
    </UserContext.Provider>
}

export default UserProvider;
