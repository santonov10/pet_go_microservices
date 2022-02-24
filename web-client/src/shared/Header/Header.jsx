import {useContext} from "react";
import {UserContext} from "../../providers/UserProvider";
import {Link} from "react-router-dom";
import './Header.css'

export default function Header() {
    const {user, logout} = useContext(UserContext)

    return (
        <div className="Header">
            {user
                ? (
                    <>
                        <div>{user?.login}</div>
                        <button onClick={logout}>
                            Выйти
                        </button>
                    </>
                )
                : <Link to="/auth">Auth</Link>}
        </div>
    );
}