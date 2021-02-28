import React, { useState, useEffect } from 'react'
import { Route, Switch, withRouter } from 'react-router-dom'
import { adminLogin } from '../api/AdminAPI'
import LoginPage from './login/LoginPage'
import MainNavbar from '../components/MainNavbar'
import HomePage from './home/HomePage'
import UserPage from './user/UserPage'
import NotFoundPage from './notfound/NotFoundPage'

const routes = [
    { id: 1, path: '/home', component: HomePage },
    { id: 2, path: '/user', component: UserPage },
]

const Navigation = (props) => {
    const [auth, setAuth] = useState(false)

    const routeList = routes.map((route) => {
        return <Route key={route.id} path={route.path} component={route.component} />
    })

    useEffect(() => {
        if (sessionStorage.getItem("auth") !== null) {
            setAuth(true)
            props.history.push({
                pathname: props.location.pathname,
            });
        }
    }, []) // eslint-disable-line react-hooks/exhaustive-deps

    const onLogin = (username, password) => {
        console.log("nav = " + username);
        console.log("nav = " + password);
        adminLogin({
            username: username,
            password: password
        }).then((response) => {
            console.log(response.headers);
            setAuth(true)
            sessionStorage.setItem("auth", "loggedIn")
            props.history.push({
                pathname: "/home"
            })
        }).catch((err) => {
            console.log("err = " + err);
        })
    }

    const onLogout = () => {
        setAuth(false)
        sessionStorage.removeItem("auth")
        props.history.push({
            pathname: "/"
        })
    }

    return (
        <div>
            <MainNavbar onLogout={onLogout} auth={auth} />
            <Switch>
                <Route exact path="/" render={() => { return <LoginPage onLogin={onLogin} /> }} />
                {routeList}
                <Route path="*" component={NotFoundPage} />
            </Switch>
        </div>
    )

}

export default withRouter(Navigation);