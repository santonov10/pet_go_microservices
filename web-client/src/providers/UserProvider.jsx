import {createContext, useCallback, useState} from "react";
import authApi from "../api/controllers/authApi";

export const UserContext = createContext({
    user: null,
    logout: () => {},
    login: () => {},
    signIn: () => {},
});

export default function UserProvider({children}) {
    const [user, setUser] = useState(null);

    const logout = useCallback(async () => {
        await authApi.signOut();
        setUser(null);
    }, []);

    const login = useCallback(async values => {
        const data =  await authApi.login(values)
        setUser(user);
    }, []);

    const signIn = useCallback(async values => {
        const user =  await authApi.signIn(values);
        setUser(user);
    }, []);

    return (
        <UserContext.Provider value={{
            user,
            logout,
            login,
            signIn
        }}>
            {children}
        </UserContext.Provider>
    )
}