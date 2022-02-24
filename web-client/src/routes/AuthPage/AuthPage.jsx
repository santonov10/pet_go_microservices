import { Formik } from 'formik';
import {useCallback, useContext, useState} from "react";
import './AuthPage.css';
import {UserContext} from "../../providers/UserProvider";

export default function AuthPage() {
    const [isLogin, setLogin] = useState(true);
    const {user, login, signIn} = useContext(UserContext);

    const onSubmit = useCallback(async values => {
        isLogin ? login(values) : signIn(values);
    }, [isLogin])

    const toggleLogin = useCallback(() => setLogin(!isLogin), [isLogin])

    return (
        <div className="AuthPage">
            <h3>
                {isLogin ? 'Авторизация' : 'Регистрация'}
            </h3>
            <Formik
                onSubmit={onSubmit}
                initialValues={user || {
                    login: '',
                    password: ''
                }}
            >
                {({
                      values,
                      handleChange,
                      handleSubmit,
                      isSubmitting,
                  }) => (
                    <form
                        onSubmit={handleSubmit}
                        className={'AuthPage__form'}
                    >
                        <label htmlFor="login">
                            Логин
                        </label>
                        <input
                            id="login"
                            type="text"
                            name="login"
                            onChange={handleChange}
                            value={values.login}
                        />
                        <label htmlFor="password">
                            Пароль
                        </label>
                        <input
                            id="password"
                            type="password"
                            name="password"
                            onChange={handleChange}
                            value={values.password}
                        />
                        <button type="submit" disabled={isSubmitting}>
                            Отправить
                        </button>
                    </form>
                )}
            </Formik>
            <button
                onClick={toggleLogin}
                className="AuthPage__switcher"
            >
                {isLogin ? 'Регистрация' : 'Авторизация'}
            </button>
        </div>
    );
}