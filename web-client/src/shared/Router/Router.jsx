import { BrowserRouter, Route, Routes } from 'react-router-dom'
import Header from "../Header";
import './Router.css';

export default function Layout({routes}) {
    return (
        <div className="Router">
            <BrowserRouter>
                <Header/>
                <Routes>
                    {routes.map(route => (
                        <Route
                            key={route.id}
                            path={route.path}
                            exact={route.exact}
                            element={route.component()}
                        />
                    ))}
                </Routes>
            </BrowserRouter>
        </div>
    );
}