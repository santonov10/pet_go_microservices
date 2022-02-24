import MainPage from "./MainPage";
import AuthPage from "./AuthPage";

export default [
    {
        id: 'MainPage',
        path: '/',
        exact: true,
        component: MainPage
    },
    {
        id: 'AuthPage',
        path: '/auth',
        exact: true,
        component: AuthPage
    },
    {
        id: 'TasksPage',
        path: '/tasks',
        exact: true,
        component: () => 'Задачи'
    }
]