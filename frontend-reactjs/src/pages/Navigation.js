import React, { useState, useEffect } from 'react'
import { Route, Switch, withRouter } from 'react-router-dom'
import Swal from 'sweetalert2';
import { adminLogin } from '../api/AdminAPI'
import LoginPage from './login/LoginPage'
import MainNavbar from '../components/MainNavbar'
import HomePage from './home/HomePage'
import UserPage from './user/UserPage'
import NotFoundPage from './notfound/NotFoundPage'

const routes = [
    { id: 1, path: '/home', component: HomePage },
    { id: 2, path: '/user', component: UserPage },
    { id: 3, path: '*', component: NotFoundPage }
]

const Navigation = (props) => {
    const [auth, setAuth] = useState(false)
    const [errorMessage, setErrorMessage] = useState("")
    const routeList = routes.map((route) => {
        return <Route key={route.id} path={route.path} render={() => { return <route.component onLogout={onLogout}  /> }} />
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
        adminLogin({
            username: username,
            password: password
        }).then((response) => {
            if (response.data.meta.code === 202) {
                Swal.fire(
                    'Login Success',
                    '',
                    'success'
                ).then(() => {
                    setAuth(true)
                    sessionStorage.setItem("auth", response.data.data.token)
                    props.history.push({
                        pathname: "/home"
                    })
                    setErrorMessage("")
                })
            } else {
                setErrorMessage("invalid username or password")
            }
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
                <Route exact path="/" render={() => { return <LoginPage onLogin={onLogin} errorMessage={errorMessage} /> }} />
                {routeList}
            </Switch>
        </div>
    )

}

export default withRouter(Navigation);